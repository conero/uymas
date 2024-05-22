package pinyin

import (
	"gitee.com/conero/uymas/str"
	"gitee.com/conero/uymas/util/rock"
	"strings"
)

const (
	SepTitle = `__LIB_TITLE__`
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

func (e Element) FirstPinyin() string {
	list := e.PinyinList()
	if list == nil {
		return ""
	}
	return list[0]
}

func (e Element) Polyphony() bool {
	return len(e.PinyinList()) > 1
}

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

func (e List) Text() []string {
	var text []string
	for _, v := range e {
		text = append(text, v.Text)
	}
	return text
}

// PyinFormat set format date
func PyinFormat(pinyin, vFmt string) string {
	switch vFmt {
	case SepTitle:
		pinyin = str.Ucfirst(pinyin)
	}
	return pinyin
}
