// Package pinyin the chinese pinyin.
package pinyin

import (
	"fmt"
	"gitee.com/conero/uymas/v2/rock"
	"gitee.com/conero/uymas/v2/util/fs"
	"regexp"
	"strings"
)

var (
	CcReg          = regexp.MustCompile("^[\u4E00-\u9FA5]$")
	ChineseToneMap = map[string]map[string]int{
		"ue": {
			"üē": 1,
			"üé": 2,
			"üě": 3,
			"üè": 4,
		},
		"a": {
			"ā": 1,
			"á": 2,
			"ǎ": 3,
			"à": 4,
		},
		"e": {
			"ē": 1,
			"é": 2,
			"ě": 3,
			"è": 4,
		},
		"i": {
			"ī": 1,
			"í": 2,
			"ǐ": 3,
			"ì": 4,
		},
		"o": {
			"ō": 1,
			"ó": 2,
			"ǒ": 3,
			"ò": 4,
		},
		"u": {
			"ū": 1,
			"ú": 2,
			"ǔ": 3,
			"ù": 4,
		},
	}
)

var (
	hanRegString = `\p{Han}`
	hanReg       *regexp.Regexp
)

const (
	SearchAlphaLimit = 1000
)

// Pinyin the pinyin dick creator
type Pinyin struct {
	filename string
	dicks    map[string]Element
}

func NewPinyin(filename string) *Pinyin {
	pyt := &Pinyin{
		filename: filename,
	}
	pyt.loadData()
	return pyt
}

// 支持：https://github.com/mozillazg/pinyin-data/blob/master/pinyin.txt 文本格式
// 数据加载
func (pyt *Pinyin) loadData() {
	lines := GetLinesFromFile(pyt.filename)
	pyt.LineToDick(lines)
}

// LineToDick turn lines to dick data.
func (pyt *Pinyin) LineToDick(lines []string) *Pinyin {
	innerDick := map[string]Element{}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		firstComment := "#"
		if line[0:1] == firstComment {
			continue
		}
		queue := strings.Split(line, ":")
		unicode := strings.TrimSpace(queue[0])
		queue = strings.Split(strings.TrimSpace(queue[1]), firstComment)
		pinyin := strings.TrimSpace(queue[0])
		chinese := strings.TrimSpace(queue[1])

		innerDick[chinese] = Element{
			Unicode: unicode,
			pinyin:  pinyin,
			Text:    chinese,
		}
	}
	pyt.dicks = innerDick
	return pyt
}

// GetPyTone get pinyin with tone
func (pyt *Pinyin) GetPyTone(chinese string) string {
	chinese = pyt.GetPyToneFunc(chinese, nil)
	return chinese
}

// GetPyToneNumber get pinyin with tone that replace by number (1-4)
func (pyt *Pinyin) GetPyToneNumber(chinese string) string {
	chinese = pyt.GetPyToneFunc(chinese, func(s string) string {
		return PyinNumber(s)
	})
	return chinese
}

// GetPyToneAlpha get pinyin without tone
func (pyt *Pinyin) GetPyToneAlpha(chinese string) string {
	chinese = pyt.GetPyToneFunc(chinese, func(s string) string {
		return PyinAlpha(s)
	})
	return chinese
}

// GetPyToneFunc get pinyin with tone by callback Func
func (pyt *Pinyin) GetPyToneFunc(chinese string, call func(string) string) string {
	queue := strings.Split(chinese, "")
	var words []string

	dicks := pyt.dicks
	for _, c := range queue {
		c = strings.TrimSpace(c)
		if c == "" {
			continue
		}
		if dd, has := dicks[c]; has {
			var word string
			if call == nil {
				word = pyt.checkOneMultipleWords(dd.pinyin)
			} else {
				word = call(pyt.checkOneMultipleWords(dd.pinyin))
			}
			words = append(words, word)
		} else {
			var word string
			if call == nil {
				word = c
			} else {
				word = call(c)
			}
			words = append(words, word)
		}
	}

	chinese = strings.Join(words, " ")
	return chinese
}

// 多音字中获取其一
func (pyt *Pinyin) checkOneMultipleWords(word string) string {
	if word != "" {
		queue := strings.Split(word, ",")
		if len(queue) > 0 {
			word = queue[0]
		}
	}

	return word
}

func (pyt *Pinyin) FileLoadSuccess() bool {
	if pyt.filename != "" {
		if !fs.ExistPath(pyt.filename) {
			return false
		}

		return len(pyt.dicks) > 0
	}
	return false
}

func (pyt *Pinyin) Dicks() map[string]Element {
	return pyt.dicks
}

// Len get the length of dicks
func (pyt *Pinyin) Len() int {
	return len(pyt.dicks)
}

// SearchByGroupFunc @todo implement grouped text queries and call callback functions
func (pyt *Pinyin) SearchByGroupFunc(s string, call func(el Element)) {
	stc := ZhSentences(s)
	// smail dynamic cache dick
	var cacheDick = map[string]Element{}
	dicks := pyt.dicks
	for _, w := range stc.Words() {
		if !IsHanWord(w) {
			// empty
			call(Element{
				Text: w,
			})
		}

		// smail
		el, exist := cacheDick[w]
		if exist {
			call(el)
			continue
		}

		// full
		el, exist = dicks[w]
		if exist {
			cacheDick[w] = el
			call(el)
			continue
		}

	}
}

// SearchByGroup @todo implement grouped text queries
func (pyt *Pinyin) SearchByGroup(words string) List {
	var elList List
	pyt.SearchByGroupFunc(words, func(el Element) {
		elList = append(elList, el)
	})
	return elList
}

// IsHanWord detect if it is chinese word.
func IsHanWord(w string) bool {
	if hanReg == nil {
		hanReg = regexp.MustCompile(hanRegString)
	}

	return hanReg.MatchString(w)
}

// PyinNumber turn pinyin with number
//
// Since alphanumeric pinyin is loaded at the end of the word rather than replacing it in situ,
// new segmentation symbols are added
func PyinNumber(word string, seqs ...string) string {
	seq := rock.Param(",", seqs...)
	var queue []string

	for _, vs := range strings.Split(word, seq) {
		for k, m := range ChineseToneMap {
			isBreak := false
			for s, n := range m {
				if strings.Contains(vs, s) {
					vs = strings.ReplaceAll(vs, s, k)
					vs = fmt.Sprintf("%v%v", vs, n)
					queue = append(queue, vs)
					isBreak = true
					break
				}
			}
			if isBreak {
				break
			}
		}
	}

	return strings.Join(queue, seq)
}

// PyinNumberList convert the pinyin number to C and output the list
func PyinNumberList(word []string) []string {
	for i, vs := range word {
		for k, m := range ChineseToneMap {
			isBreak := false
			for s, n := range m {
				if strings.Contains(vs, s) {
					vs = strings.ReplaceAll(vs, s, k)
					vs = fmt.Sprintf("%v%v", vs, n)
					isBreak = true
					break
				}
			}
			if isBreak {
				break
			}
		}
		word[i] = vs
	}

	return word
}

// PyinAlpha turn pinyin with alpha
func PyinAlpha(word string, isAllArgs ...bool) string {
	isAll := rock.Param(false, isAllArgs...)
	for k, m := range ChineseToneMap {
		isBreak := false
		for s := range m {
			if strings.Contains(word, s) {
				word = strings.ReplaceAll(word, s, k)
				isBreak = !isAll
				if isBreak {
					break
				} else {
					continue
				}
			}
		}
		if isBreak {
			break
		}
	}
	return word
}

// SearchAlpha search word by single alpha
func (pyt *Pinyin) SearchAlpha(alpha string, limits ...int) List {
	limit := rock.Param(SearchAlphaLimit, limits...)
	list := List{}

	for _, v := range pyt.dicks {
		matchAlpha := PyinAlpha(v.pinyin)
		if limit > 0 && len(list) >= limit {
			break
		}
		// Search from polyphonic word
		pyList := v.PinyinList()
		if len(pyList) > 0 {
			isContinue := false
			for _, py := range pyList {
				if strings.Index(py, alpha) == 0 {
					list = append(list, v)
					isContinue = true
					break
				}
			}
			if isContinue {
				continue
			}
		}
		if strings.Index(matchAlpha, alpha) == 0 {
			list = append(list, v)
		}
	}

	return list
}

func (pyt *Pinyin) IsEmpty() bool {
	return pyt.Len() == 0
}
