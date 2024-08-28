// Package util implements other tool more, like type cover, type value check.
//
// Usually based on the reflection implementation
package util

import (
	"reflect"
)

// @Date：   2018/10/30 0030 13:26
// @Author:  Joshua Conero
// @Name:    工具栏

// NullDefault null value handler to default.
func NullDefault(value, def any) any {
	if ValueNull(value) {
		return def
	}
	return value
}

// ValueNull to find if is null
func ValueNull(value any) bool {
	if nil == value {
		return true
	}
	v := reflect.ValueOf(value)
	return v.IsZero()
}
