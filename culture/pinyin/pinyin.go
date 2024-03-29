// Package pinyin the chinese pinyin.
package pinyin

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	CcReg          = regexp.MustCompile("^[\u4E00-\u9FA5]$")
	ChineseToneMap = map[string]map[string]int{
		"ue": map[string]int{
			"üē": 1,
			"üé": 2,
			"üě": 3,
			"üè": 4,
		},
		"a": map[string]int{
			"ā": 1,
			"á": 2,
			"ǎ": 3,
			"à": 4,
		},
		"e": map[string]int{
			"ē": 1,
			"é": 2,
			"ě": 3,
			"è": 4,
		},
		"i": map[string]int{
			"ī": 1,
			"í": 2,
			"ǐ": 3,
			"ì": 4,
		},
		"o": map[string]int{
			"ō": 1,
			"ó": 2,
			"ǒ": 3,
			"ò": 4,
		},
		"u": map[string]int{
			"ū": 1,
			"ú": 2,
			"ǔ": 3,
			"ù": 4,
		},
	}
)

// Pinyin the pinyin dick creator
type Pinyin struct {
	filename string
	Dicks    map[string]map[string]string
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
	innerDick := map[string]map[string]string{}
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

		innerDick[chinese] = map[string]string{
			"unicode": unicode,
			"pinyin":  pinyin,
			"chinese": chinese,
		}
	}
	pyt.Dicks = innerDick
	return pyt
}

// GetPyTone get pinyin with tone
func (pyt *Pinyin) GetPyTone(chinese string) string {
	chinese = pyt.GetPyToneFunc(chinese, nil)
	return chinese
}

// GetPyToneNumber get pinyin with tone that replace by number (1-4)
func (pyt *Pinyin) GetPyToneNumber(chinese string) string {
	chinese = pyt.GetPyToneFunc(chinese, func(word string) string {
		for k, m := range ChineseToneMap {
			isBreak := false
			for s, n := range m {
				if strings.Index(word, s) > -1 {
					word = strings.ReplaceAll(word, s, k)
					word = fmt.Sprintf("%v%v", word, n)
					isBreak = true
					break
				}
			}
			if isBreak {
				break
			}
		}
		return word
	})
	return chinese
}

// GetPyToneAlpha get pinyin without tone
func (pyt *Pinyin) GetPyToneAlpha(chinese string) string {
	chinese = pyt.GetPyToneFunc(chinese, func(word string) string {
		for k, m := range ChineseToneMap {
			isBreak := false
			for s, _ := range m {
				if strings.Index(word, s) > -1 {
					word = strings.ReplaceAll(word, s, k)
					isBreak = true
					break
				}
			}
			if isBreak {
				break
			}
		}
		return word
	})
	return chinese
}

// GetPyToneFunc get pinyin with tone by callback Func
func (pyt *Pinyin) GetPyToneFunc(chinese string, call func(string) string) string {
	queue := strings.Split(chinese, "")
	var words []string

	dicks := pyt.Dicks
	for _, c := range queue {
		c = strings.TrimSpace(c)
		if c == "" {
			continue
		}
		if dd, has := dicks[c]; has {
			var word string
			if call == nil {
				word = pyt.checkOneMultipleWords(dd["pinyin"])
			} else {
				word = call(pyt.checkOneMultipleWords(dd["pinyin"]))
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
