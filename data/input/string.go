package input

import (
	"strconv"
	"strings"
)

// Stringer string input
//
// As a string type fast converter, and no exceptions are thrown
type Stringer string

func (s Stringer) Int64() int64 {
	if s == "" {
		return 0
	}
	v, _ := strconv.ParseInt(floatStringTrim(string(s)), 10, 64)
	return v
}

func (s Stringer) Float() float64 {
	if s == "" {
		return 0
	}
	v, _ := strconv.ParseFloat(string(s), 64)
	return v
}

func (s Stringer) Uint64() uint64 {
	if s == "" {
		return 0
	}
	v, _ := strconv.ParseUint(floatStringTrim(string(s)), 10, 64)
	return v
}

func (s Stringer) Uint32() uint32 {
	if s == "" {
		return 0
	}
	v, _ := strconv.ParseUint(floatStringTrim(string(s)), 10, 32)
	return uint32(v)
}

func (s Stringer) Int() int {
	iVal, _ := strconv.Atoi(floatStringTrim(string(s)))
	return iVal
}

func (s Stringer) Bool() bool {
	bVal, _ := strconv.ParseBool(string(s))
	return bVal
}

// make float string to int
func floatStringTrim(s string) string {
	if strings.Contains(s, "_") {
		s = strings.ReplaceAll(s, "_", "")
	}
	idx := strings.Index(s, ".")
	if idx > 0 {
		return s[:idx]
	}
	return s
}
