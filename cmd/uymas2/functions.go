package main

import (
	"gitee.com/conero/uymas/v2/culture/pinyin"
	"gitee.com/conero/uymas/v2/culture/pinyin/material"
)

var (
	gPinyinCache *pinyin.Pinyin
)

func getPinyin() *pinyin.Pinyin {
	if gPinyinCache == nil {
		py := material.NewPinyin()
		gPinyinCache = &py
	}
	return gPinyinCache
}
