package input

import (
	"strconv"
)

type SimpleStr string

func (s SimpleStr) Int64() int64 {
	if s == "" {
		return 0
	}
	v, _ := strconv.ParseInt(floatStringTrim(string(s)), 10, 64)
	return v
}

func (s SimpleStr) Float() float64 {
	if s == "" {
		return 0
	}
	v, _ := strconv.ParseFloat(string(s), 64)
	return v
}

func (s SimpleStr) Uint64() uint64 {
	if s == "" {
		return 0
	}
	v, _ := strconv.ParseUint(floatStringTrim(string(s)), 10, 64)
	return v
}

func (s SimpleStr) Uint32() uint32 {
	if s == "" {
		return 0
	}
	v, _ := strconv.ParseUint(floatStringTrim(string(s)), 10, 32)
	return uint32(v)
}

func (s SimpleStr) Int() int {
	iVal, _ := strconv.Atoi(floatStringTrim(string(s)))
	return iVal
}

func (s SimpleStr) Bool() bool {
	if s == "" {
		return false
	}
	bVal, _ := strconv.ParseBool(string(s))
	return bVal
}
