//the chinese pinyin.
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

type Pinyin struct {
	Filename        string                       // 文件
	Dicks           map[string]map[string]string // 地址
	pinyinWordsDick map[string]string            // 拼音字典； [pinyin => 文字字典]
}

// 文件读取
func (py *Pinyin) ReadIni(filename string) *Pinyin {
	py.Filename = filename
	py.Dicks = ReadDickFromIni(filename)
	py.simpleWordsDick()
	return py
}

func (py *Pinyin) ReadIniLines(lines []string) *Pinyin {
	py.Dicks = ReadIniLines(lines)
	py.simpleWordsDick()
	return py
}

// 内部文件读取
func (py *Pinyin) readInnerIni() *Pinyin {
	py.Dicks = ReadDickFromIni(py.Filename)
	return py
}

// 字典简单化
func (py *Pinyin) simpleWordsDick() {
	if py.Dicks != nil {
		py.pinyinWordsDick = map[string]string{}
		for _, vv := range py.Dicks {
			for k, v := range vv {
				py.pinyinWordsDick[k] = v
			}
		}
	}
}

// 文件读取
func NewPinyin() *Pinyin {
	py := new(Pinyin)
	return py
}

// 获取汉字的拼音
func (py *Pinyin) GetWord(word string) string {
	var pinyin string
	if py.pinyinWordsDick != nil {
		for py, words := range py.pinyinWordsDick {
			if strings.Index(words, word) > -1 {
				pinyin = py
				break
			}
		}
	}
	return pinyin
}

// 字符分割
func (py *Pinyin) GetWordsSplit(words string) []string {
	var pinyin = []string{}
	wordsSplit := ([]rune)(words)
	for i := 0; i < len(wordsSplit); i++ {
		wd := string(wordsSplit[i : i+1])
		if IsChineseCharacters(wd) {
			pinyin = append(pinyin, py.GetWord(wd))
		} else {
			pinyin = append(pinyin, wd)
		}

	}
	return pinyin
}

//获取句子拼音
func (py *Pinyin) GetWords(words string) string {
	var pinyin string
	wordsSplit := ([]rune)(words)
	for i := 0; i < len(wordsSplit); i++ {
		wd := string(wordsSplit[i : i+1])
		if IsChineseCharacters(wd) {
			pinyin += py.GetWord(wd)
		} else {
			pinyin += wd
		}

	}
	return pinyin
}

//含声调的拼音
type PinyinTone struct {
	filename string
	Dicks    map[string]map[string]string
}

//初始化
func NewPinyinTone(filename string) *PinyinTone {
	pyt := &PinyinTone{
		filename: filename,
	}
	pyt.loadData()
	return pyt
}

//支持：https://github.com/mozillazg/pinyin-data/blob/master/pinyin.txt 文本格式
//数据加载
func (pyt *PinyinTone) loadData() {
	lines := com.ReadIniLines(pyt.filename)
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
}

//获取拼音声调
func (pyt *PinyinTone) GetPyTone(chinese string) string {
	chinese = pyt.GetPyToneFunc(chinese, nil)
	return chinese
}

//拼音数字标注法
func (pyt *PinyinTone) GetPyToneNumber(chinese string) string {
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

//拼音数字标注法
func (pyt *PinyinTone) GetPyToneAlpha(chinese string) string {
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

//获取pinyin tone 字符，待回调
// call 为 `nil` 是为默认
func (pyt *PinyinTone) GetPyToneFunc(chinese string, call func(word string) string) string {
	queue := strings.Split(chinese, "")
	words := []string{}

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

//多音字中获取其一
func (pyt *PinyinTone) checkOneMultipleWords(word string) string {
	if word != "" {
		queue := strings.Split(word, ",")
		if len(queue) > 0 {
			word = queue[0]
		}
	}

	return word
}
