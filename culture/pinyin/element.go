package pinyin

import (
	"fmt"
	"gitee.com/conero/uymas/str"
	"gitee.com/conero/uymas/util/rock"
	"strings"
)

const (
	SepTitle = `__LIB_TITLE__`
)

const (
	_ = iota
	// PinyinTone pinyin as tone for raw
	PinyinTone
	// PinyinNumber pinyin as tone for number
	PinyinNumber
	// PinyinAlpha pinyin as tone for alpha
	PinyinAlpha
)

// Element the data dictionary enter
type Element struct {
	Unicode string
	// possible existence of polyphonic characters
	pinyin string
	// it can be chinese or other char
	Text string
}

// IsEmpty test if is empty that support all unicode
func (e Element) IsEmpty() bool {
	return e.pinyin == ""
}

// FirstPinyin if pinyin exists get the first pinyin, compatible with polyphonics
func (e Element) FirstPinyin() string {
	list := e.PinyinList()
	if list == nil {
		return ""
	}
	return list[0]
}

// Polyphony test if the pinyin is polyphonic
func (e Element) Polyphony() bool {
	return len(e.PinyinList()) > 1
}

// PinyinList get the polyphonic pinyin as list
func (e Element) PinyinList() []string {
	if e.pinyin == "" {
		return nil
	}

	return strings.Split(e.pinyin, ",")
}

// List the list of elements
type List []Element

func (e List) String() string {
	var queue []string
	for _, v := range e {
		queue = append(queue, v.Text)
	}
	return strings.Join(queue, "")
}

// Tone Tone(seps, fmt string)
func (e List) Tone(seps ...string) string {
	sep := rock.ExtractParam("", seps...)
	vFmt := rock.ExtractParamIndex("", 2, seps...)
	var queue []string
	for _, v := range e {
		if v.IsEmpty() {
			queue = append(queue, v.Text)
		} else {
			queue = append(queue, PyinFormat(v.FirstPinyin(), vFmt))
		}
	}
	return strings.Join(queue, sep)
}

// Number Number(seps, fmt string)
func (e List) Number(seps ...string) string {
	sep := rock.ExtractParam("", seps...)
	vFmt := rock.ExtractParamIndex("", 2, seps...)
	var queue []string
	for _, v := range e {
		if v.IsEmpty() {
			queue = append(queue, v.Text)
		} else {
			queue = append(queue, PyinFormat(PyinNumber(v.FirstPinyin()), vFmt))
		}
	}
	return strings.Join(queue, sep)
}

// Alpha Alpha(seps, fmt string)
func (e List) Alpha(seps ...string) string {
	sep := rock.ExtractParam("", seps...)
	vFmt := rock.ExtractParamIndex("", 2, seps...)
	var queue []string
	for _, v := range e {
		if v.IsEmpty() {
			queue = append(queue, v.Text)
		} else {
			queue = append(queue, PyinFormat(PyinAlpha(v.FirstPinyin()), vFmt))
		}
	}
	return strings.Join(queue, sep)
}

// Text get pinyin text as the list of element
func (e List) Text() []string {
	var text []string
	for _, v := range e {
		text = append(text, v.Text)
	}
	return text
}

// Polyphony gets all columns composed of polyphonics
//
// @todo to be realized
func (e List) Polyphony(args ...string) []string {
	var polys []string
	queue := polyphonyTraverse(e, 0, nil)
	//fmt.Printf("queue: %#v\n", queue)
	fmt.Printf("queue: %v\n", len(queue))
	return polys
}

// PyinFormat set format date
//
// support `title` like 'latest name' to 'Latest Name'
func PyinFormat(pinyin, vFmt string) string {
	switch vFmt {
	case SepTitle:
		pinyin = str.Ucfirst(pinyin)
	}
	return pinyin
}

// recursive polyphonics
func polyphonyTraverse(ls List, next int, queue []string) []string {
	vLen := len(ls)
	for j := next; j < vLen; j++ {
		elNext := ls[j]
		if elNext.IsEmpty() {
			queue = append(queue, elNext.Text)
			continue
		}
		l := elNext.PinyinList()
		if len(l) == 1 {
			queue = append(queue, elNext.pinyin)
			continue
		}

		for _, cv := range l {
			fmt.Printf("i: %d, cv: %s\n", j, cv)
			queue = append(queue, cv)
			child := polyphonyTraverse(ls, j+1, queue)
			queue = append(queue, child...)
		}
	}
	return queue
}
