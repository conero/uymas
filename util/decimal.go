package util

import (
	"github.com/conero/uymas/str"
	"strings"
)

// @Date：   2018/12/12 0012 13:45
// @Author:  Joshua Conero
// @Name:    10 进制处理
const (
	// 9+26+26+3
	NumberStr = "0123456789abcdefghijklmnopkrstuvwxyzABCDEFGHIJKLMNOPKRSTUVWXYZ+-="
	//进制类型
	N2  = 2
	N8  = 8
	N16 = 16
	N32 = 32
	N36 = 36
	N62 = 62
)

type Decimal struct {
	dec int
}

// 转化为N进制
func (d *Decimal) ToN(base int) string {
	num := d.dec
	bits := []int{}
	for {
		if num < base {
			bits = append(bits, num)
			break
		}
		mod := num % base
		bits = append(bits, mod)
		num = (num - mod) / base
	}
	maxLen := len(NumberStr)
	var value string
	if base <= maxLen {
		nRefBits := strings.Split(NumberStr[:base], "")
		nBits := []string{}
		for _, n := range bits {
			nBits = append(nBits, nRefBits[n])
		}
		value = str.Reverse(strings.Join(nBits, ""))
	}
	return value
}

// 2 进制
func (d *Decimal) T2() string {
	return d.ToN(N2)
}

// 8 进制
func (d *Decimal) T8() string {
	return d.ToN(N8)
}

// 16 进制
func (d *Decimal) T16() string {
	return d.ToN(N16)
}

// 32 进制
func (d *Decimal) T32() string {
	return d.ToN(N32)
}

// 36 进制
func (d *Decimal) T36() string {
	return d.ToN(N36)
}

func (d *Decimal) T62() string {
	return d.ToN(N62)
}

// 十进制
func NewDec(dec int) *Decimal {
	return &Decimal{dec}
}
