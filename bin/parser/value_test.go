package parser

import (
	"testing"
)

func TestConvBool(t *testing.T) {
	var raw = "True"

	// case
	if !ConvBool(raw) {
		t.Errorf("%v: 字符串解析为bool失败", raw)
	}

	// case
	raw = "TRUE"
	if !ConvBool(raw) {
		t.Errorf("%v: 字符串解析为bool失败", raw)
	}

	// case
	raw = "vaild"
	if ConvBool(raw) {
		t.Errorf("%v: 字符串解析为bool失败", raw)
	}
}

func TestConvI64(t *testing.T) {
	raw := "1949"
	if 1949 != ConvI64(raw) {
		t.Errorf("%v: 字符串解析为 int64 失败", raw)
	}

	// case
	raw = "+1949"
	if 1949 != ConvI64(raw) {
		t.Errorf("%v: 字符串解析为 int64 失败", raw)
	}

	// case
	raw = "-1949"
	if -1949 != ConvI64(raw) {
		t.Errorf("%v: 字符串解析为 int64 失败", raw)
	}

	// case
	raw = "-1949.01"
	rf := ConvI64(raw)
	if -1949 != rf {
		t.Errorf("%v: 字符串解析为 int64 失败 => %v", raw, rf)
	}

	// case
	raw = "yang"
	if 0 != ConvI64(raw) {
		t.Errorf("%v: 字符串解析为 int64 失败", raw)
	}
}

func TestConvInt(t *testing.T) {
	raw := "1949"
	if 1949 != ConvInt(raw) {
		t.Errorf("%v: 字符串解析为 int64 失败", raw)
	}

	// case
	raw = "+1949"
	if 1949 != ConvInt(raw) {
		t.Errorf("%v: 字符串解析为 int64 失败", raw)
	}

	// case
	raw = "-1949"
	if -1949 != ConvInt(raw) {
		t.Errorf("%v: 字符串解析为 int64 失败", raw)
	}

	// case
	raw = "-1949.01"
	rf := ConvInt(raw)
	if -1949 != rf {
		t.Errorf("%v: 字符串解析为 int64 失败 => %v", raw, rf)
	}

	// case
	raw = "yang"
	if 0 != ConvInt(raw) {
		t.Errorf("%v: 字符串解析为 int64 失败", raw)
	}
}

func TestConvF64(t *testing.T) {
	raw := "1949"
	if 1949 != ConvF64(raw) {
		t.Errorf("%v: 字符串解析为 float64 失败", raw)
	}

	// case
	raw = "+1949"
	if 1949 != ConvF64(raw) {
		t.Errorf("%v: 字符串解析为 float64 失败", raw)
	}

	// case
	raw = "-1949"
	if -1949 != ConvF64(raw) {
		t.Errorf("%v: 字符串解析为 float64 失败", raw)
	}

	// case
	raw = "yang"
	if 0 != ConvF64(raw) {
		t.Errorf("%v: 字符串解析为 float64 失败", raw)
	}

	// case
	raw = "3.14159265359"
	if 3.14159265359 != ConvF64(raw) {
		t.Errorf("%v: 字符串解析为 float64 失败", raw)
	}

	// case
	raw = "-3.14159265359"
	if -3.14159265359 != ConvF64(raw) {
		t.Errorf("%v: 字符串解析为 float64 失败", raw)
	}

	// case
	raw = "-1949.01"
	rf := ConvF64(raw)
	if -1949.01 != rf {
		t.Errorf("%v: 字符串解析为 float64 失败 => %v", raw, rf)
	}
}
