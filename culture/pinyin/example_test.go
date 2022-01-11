package pinyin

import (
	"fmt"
	"testing"
)

func TestNewPinyin(t *testing.T) {
	ExampleNewPinyin()
}

func ExampleNewPinyin() {
	var filename, pinyin string
	// filename is pinyin dick file or use `pinyin/material` for builtin.
	py := NewPinyin(filename)
	pinyin = py.GetPyToneAlpha(`古丞秋`)
	fmt.Println(pinyin)
	pinyin = py.GetPyTone(`中华人民共和国`)
	fmt.Println(pinyin)

	// Output:
	// gu cheng qiu
	// zhōng huá rén mín gòng hé guó
}
