package util

import (
	"gitee.com/conero/uymas/v2/str"
	"strings"
)

// @Date：   2018/12/12 0012 13:45
// @Author:  Joshua Conero
// @Name:    10 进制处理
const (
	// NumberStr 9+26+26+3
	NumberStr = "0123456789abcdefghijklmnopkrstuvwxyzABCDEFGHIJKLMNOPKRSTUVWXYZ+-="
	// N2 进制类型
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

// ToN convert to n-ary
func (d *Decimal) ToN(base int) string {
	num := d.dec
	var bits []int
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
		var nBits []string
		for _, n := range bits {
			nBits = append(nBits, nRefBits[n])
		}
		value = str.Reverse(strings.Join(nBits, ""))
	}
	return value
}

// T2 2 Base system
func (d *Decimal) T2() string {
	return d.ToN(N2)
}

// T8 8 Base system
func (d *Decimal) T8() string {
	return d.ToN(N8)
}

// T16 16 Base system
func (d *Decimal) T16() string {
	return d.ToN(N16)
}

// T32 32 Base system
func (d *Decimal) T32() string {
	return d.ToN(N32)
}

// T36 36 Base system
func (d *Decimal) T36() string {
	return d.ToN(N36)
}

func (d *Decimal) T62() string {
	return d.ToN(N62)
}

// NewDec decimal system
func NewDec(dec int) *Decimal {
	return &Decimal{dec}
}
