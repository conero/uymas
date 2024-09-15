// Package convert common data converter, base "reflect".
package convert

import (
	"gitee.com/conero/uymas/v2/data/input"
	"gitee.com/conero/uymas/v2/str"
	"reflect"
	"strings"
)

// SetByStr set the literal string to any specified type
//
// The types are as follows:
//
// 1. string, convert the original input string.
//
// 2. bool, parse string true, True, false, False etc.
//
// 3. int64, parse a numeric string to int64.
//
// 4. float, parse numeric strings into floats.
//
// 5. uint64, parse a numeric string to uint64.
func SetByStr(value reflect.Value, s string) bool {
	if !value.CanSet() {
		return false
	}

	kind := value.Kind()
	// string
	if kind == reflect.String {
		value.Set(reflect.ValueOf(s))
		return true
	}

	if kind == reflect.Bool {
		value.Set(reflect.ValueOf(input.Stringer(s).Bool()))
		return true
	}

	vSlice, isMatch := ToSlice(s)
	if isMatch {
		return SetByStrSlice(value, vSlice)
	}

	if value.CanInt() {
		value.SetInt(input.Stringer(s).Int64())
		return true
	}

	if value.CanFloat() {
		value.SetFloat(input.Stringer(s).Float())
		return true
	}

	if value.CanUint() {
		value.SetUint(input.Stringer(s).Uint64())
		return true
	}

	return false
}

// SetByStrSlice Assigns a string array to another type of slice
func SetByStrSlice(value reflect.Value, vSlice []string) bool {
	vKind := value.Kind()
	var sliceType reflect.Type

	sliceValue := value
	if vKind == reflect.Ptr {
		sliceValue = value.Elem()
		sliceType = sliceValue.Type()
	} else {
		sliceType = value.Type()
	}

	if sliceValue.Kind() != reflect.Slice {
		return false
	}

	vLen := len(vSlice)
	newSlice := reflect.MakeSlice(sliceType, vLen, vLen)
	for i, s := range vSlice {
		if !SetByStr(newSlice.Index(i), s) {
			return false
		}
	}

	sliceValue.Set(newSlice)
	return true
}

// ToSlice convert string to slice, format `[e0, e1, ..., en]`
func ToSlice(s string) (vSlice []string, isMatch bool) {
	if s == "" {
		return
	}

	vLen := len(s)
	if s[:1] == "[" && s[vLen-1:] == "]" {
		ss := s[1 : vLen-1]
		vSlice = strings.Split(str.Str(ss).ClearSpace(), ",")
		isMatch = true
		return
	}

	return
}

// IsSlice determine whether the string conforms to the slice
func IsSlice(s string) bool {
	vLen := len(s)
	if vLen < 2 {
		return false
	}
	return s[:1] == "[" && s[vLen-1:] == "]"
}
