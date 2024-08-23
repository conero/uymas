package input

import "strconv"

// Stringer string input
//
// As a string type fast converter, and no exceptions are thrown
type Stringer string

func (s Stringer) Int64() int64 {
	if s == "" {
		return 0
	}
	v, _ := strconv.ParseInt(string(s), 10, 64)
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
	v, _ := strconv.ParseUint(string(s), 10, 64)
	return v
}

func (s Stringer) Int() int {
	iVal, _ := strconv.Atoi(string(s))
	return iVal
}

func (s Stringer) Bool() bool {
	bVal, _ := strconv.ParseBool(string(s))
	return bVal
}
