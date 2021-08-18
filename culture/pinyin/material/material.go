// Package material the Material of pinyin dick
package material

import (
	_ "embed"
	"fmt"
	"gitee.com/conero/uymas/culture/pinyin"
	"gitee.com/conero/uymas/number"
)

// reference the resource from the link: https://github.com/mozillazg/pinyin-data
//go:embed mt_pinyin.txt
var pinyinDick []byte

// the common list word, resource link: https://www.zdic.net/zd/zb/ty/
//go:embed mt_common_list.ini
var commonList []byte

// GetDickRaw get the raw pinyin dick
func GetDickRaw() []byte {
	return pinyinDick
}

// GetCommonRaw get the raw of common list
func GetCommonRaw() []byte {
	return commonList
}

// NewPinyin create Pinyin instance.
func NewPinyin() pinyin.Pinyin {
	py := pinyin.Pinyin{}
	py.LineToDick(pinyin.GetLinesFromByte(GetDickRaw()))
	return py
}

type CommonList struct {
	dicks map[string]string
}

// WordList get word list by strokes(笔画)
func (c *CommonList) WordList(strokes int) string {
	if c.dicks != nil {
		value, _ := c.dicks[fmt.Sprintf("%v", strokes)]
		return value
	}
	return ""
}

// StrokesList get strokes(笔画) list
func (c *CommonList) StrokesList() []int {
	var sl []int
	for idx, _ := range c.dicks {
		sl = append(sl, int(number.AnyInt64(idx)))
	}
	return sl
}

func NewCommonList() *CommonList {
	return &CommonList{
		dicks: pinyin.ReadDickFromByteKv(GetCommonRaw()),
	}
}
