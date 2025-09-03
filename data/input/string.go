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
	value := string(s)
	fVal, base, isMatch := supportUnit(value)
	if isMatch {
		return int(fVal * float64(base))
	}
	iVal, _ := strconv.Atoi(floatStringTrim(value))
	return iVal
}

func (s Stringer) Bool() bool {
	bVal, _ := strconv.ParseBool(string(s))
	return bVal
}

// make float string to int
func floatStringTrim(s string) string {
	// 100_000.00
	if strings.Contains(s, "_") {
		s = strings.ReplaceAll(s, "_", "")
	}
	// 100,000.00
	if strings.Contains(s, ",") {
		s = strings.ReplaceAll(s, ",", "")
	}
	idx := strings.Index(s, ".")
	if idx > 0 {
		return s[:idx]
	}
	return s
}

// split string to float and unit
func supportUnit(s string) (float64, int, bool) {
	if s == "" {
		return 0, 0, false
	}

	// k-1000
	rpl := strings.ToLower(s)
	index := strings.IndexAny(rpl, "k")
	if index > 0 {
		f, _ := strconv.ParseFloat(strings.TrimSpace(rpl[:index]), 64)
		return f, 1000, true
	}

	// w-10000
	index = strings.IndexAny(rpl, "w")
	if index > 0 {
		f, _ := strconv.ParseFloat(strings.TrimSpace(rpl[:index]), 64)
		return f, 10_000, true
	}

	// m-1,000,000
	index = strings.IndexAny(rpl, "m")
	if index > 0 {
		f, _ := strconv.ParseFloat(strings.TrimSpace(rpl[:index]), 64)
		return f, 1_000_000, true
	}

	// b-1,000,000,000
	index = strings.IndexAny(rpl, "b")
	if index > 0 {
		f, _ := strconv.ParseFloat(strings.TrimSpace(rpl[:index]), 64)
		return f, 1_000_000_000, true
	}
	return 0, 0, false
}
