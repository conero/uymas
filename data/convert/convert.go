// Package convert common data converter
package convert

import (
	"gitee.com/conero/uymas/v2/data/input"
	"reflect"
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
// @todo bug
func SetByStrSlice(value reflect.Value, vSlice []string) bool {
	vKind := value.Kind()
	var sliceType reflect.Type
	if vKind == reflect.Ptr {
		sliceType = value.Elem().Type()
	} else {
		sliceType = value.Type()
	}

	vLen := len(vSlice)

	newSlice := reflect.MakeSlice(sliceType, vLen, vLen)
	for i, s := range vSlice {
		if !SetByStr(newSlice.Index(i), s) {
			return false
		}
	}

	value.Set(newSlice)
	return true
}
