package rock

import (
	"fmt"
)

func ExampleFormatList() {
	s := FormatList([]string{"1", "this\\'s test", "box", "a", "must", "base", "念奴娇", "√√√√√", "", "saw", "no-list", "box",
		"box", "box", "box", "box", "box", "box", "box", "box", "box", "box", "box", "box", "box"})
	fmt.Println("列表输出如下：\n" + s)
}

func ExampleFormatKv() {
	s := FormatKv(map[string]any{
		"name":    "Joshua Conero",
		"meeting": 20.1,
	})
	fmt.Println("列表输出如下：\n" + s)

	// Output:
	// name       Joshua Conero
	// meeting    20.1
}

// 用于执行 Example 用例
//func TestExample(t *testing.T) {
//	ExampleFormatList()
//}
