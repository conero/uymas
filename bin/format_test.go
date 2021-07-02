package bin

import (
	"testing"
)

func TestFormatKv(t *testing.T) {
	tdd := map[string]interface{}{
		"author":                   "Joshua Conero",
		"email":                    "conero@163",
		"a":                        "TestFormatKv for beautify string.",
		"canBeALongStringTestAlso": 2,
	}
	t.Log("\n" + FormatKv(tdd))
	t.Log("\n" + FormatKv(tdd, ". "))
	t.Log("\n" + FormatKv(tdd, ". ", "*"))

	// support more type.
	tdd2 := map[interface{}]interface{}{
		"author":                   "Joshua Conero",
		210609:                     "conero@163",
		true:                       "TestFormatKv for beautify string.",
		"canBeALongStringTestAlso": 2,
	}
	t.Log("\n" + FormatKv(tdd2))
	t.Log("\n" + FormatKv(tdd2, ". "))
	t.Log("\n" + FormatKv(tdd2, ". ", "*"))
}

func TestFormatQue(t *testing.T) {
	// 用于输出格式
	// 实际测试时将会忽略信息
	//t.Skip()

	var data []interface{}
	//CASE1
	data = []interface{}{
		"中文", "letter", "letter", "letter", "letter", "letter", "letter",
		"中文", "letter", "letter", "letter", "letter", "letter", "letter",
	}
	t.Logf("字符串：\n%v", FormatQue(data))
	t.Logf("字符串：\n%v", FormatQue(data, " .", "-"))

	//CASE2
	data = []interface{}{
		"My name is JC", 1992, nil, "SomeMore", 3.14,
	}
	t.Logf("interface：\n%v", FormatQue(data))
	t.Logf("interface：\n%v", FormatQue(data, " .", "-"))

	//[]int
	t.Logf("[]int：\n%v", FormatQue([]int{1992, 5, 20210618, 0, 62}))
}
