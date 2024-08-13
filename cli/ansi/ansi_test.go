package ansi

import (
	"testing"
)

func TestStyle(t *testing.T) {
	input := "Yzs."
	styleText := Style(input, Green)
	t.Logf("%v, raw: %#v", styleText, styleText)
	clearStr := Clear(styleText)
	toTestFn := func() {
		if clearStr != input {
			t.Errorf("颜色码设置与清洗不匹配！%#v -> %#v", styleText, clearStr)
		}
	}

	toTestFn()

	// case
	input = "Wip: 2024, uymas v2 coding."
	styleText = Style(input, Black, BkgCyan)
	t.Logf("%v, raw: %#v", styleText, styleText)
	clearStr = Clear(styleText)
	toTestFn()
}

func TestClear(t *testing.T) {
	styleStr := "\033[31mThis is red text.\033[0m"
	ref := "This is red text."

	toTestFn := func() {
		t.Log(styleStr)
		clearStr := Clear(styleStr)
		if ref != clearStr {
			t.Errorf("clear 错误：%#v -> %#v", clearStr, ref)
		}
	}

	// case
	toTestFn()
}
