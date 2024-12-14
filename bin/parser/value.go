package parser

import (
	"strconv"
	"strings"
)

// Value Analysis of original string type conversion
type Value struct {
	original string
}

func (c *Value) Bool() bool {
	return ConvBool(c.original)
}

func NewValue(value string) *Value {
	return &Value{
		original: value,
	}
}

// ConvBool convert string to bool
func ConvBool(raw string) (value bool) {
	if raw == "" {
		return
	}

	v, _ := strconv.ParseBool(strings.ToLower(raw))
	return v
}

// ConvI64 convert string to int64, compatible float types
func ConvI64(raw string) (value int64) {
	if raw == "" {
		return
	}

	var er error
	value, er = strconv.ParseInt(raw, 10, 64)
	if er == nil {
		return
	}

	flt, err := strconv.ParseFloat(raw, 64)
	if err == nil {
		value = int64(flt)
	}
	return
}

// ConvInt convert string to int
func ConvInt(raw string) (value int) {
	if raw == "" {
		return
	}

	var err error
	value, err = strconv.Atoi(raw)
	if err != nil {
		value = int(ConvI64(raw))
	}
	return
}

// ConvF64 convert string to float64
func ConvF64(raw string) (value float64) {
	if raw == "" {
		return
	}
	value, _ = strconv.ParseFloat(raw, 64)
	return
}
