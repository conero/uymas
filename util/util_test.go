package util

import (
	"testing"
)

// @Date：   2018/12/12 0012 14:25
// @Author:  Joshua Conero
// @Name:    名称描述

func TestValueNull(t *testing.T) {
	//null values test
	values := []any{
		"", 0, int16(0), int32(0), int64(0), float32(0), float64(0), false,
	}
	for _, v := range values {
		if !ValueNull(v) {
			t.Fail()
		}
	}

	//not null values test
	values = []any{
		" ", -1, int16(3), int32(2), int64(1), float32(-0.00000001), 0.01, true,
	}
	for _, v := range values {
		if ValueNull(v) {
			t.Fail()
		}
	}
}
