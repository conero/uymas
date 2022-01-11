package material

import (
	"fmt"
	"testing"
)

func TestNewPinyin_example(t *testing.T) {
	ExampleNewPinyin()
}

func ExampleNewPinyin() {
	var pinyin string
	py := NewPinyin()
	pinyin = py.GetPyToneAlpha(`古丞秋`)
	fmt.Println(pinyin)
	pinyin = py.GetPyTone(`中华人民共和国`)
	fmt.Println(pinyin)

	// Output:
	// gu cheng qiu
	// zhōng huá rén mín gòng hé guó
}
