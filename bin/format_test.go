package bin

import (
	"math"
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

	// struct
	type Ty struct {
		Name     string
		Age      int
		Sex      byte
		birthday uint
	}
	ty := &Ty{
		Name:     "Joshua Conero",
		Age:      1994,
		Sex:      'M',
		birthday: 0522,
	}
	t.Logf("Struct/Ty -> %v\n", FormatKv(*ty))
	t.Logf("Struct/Ty -> %v\n", FormatKv(ty))
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

func TestFormatTable(t *testing.T) {
	// 用于输出格式
	// 实际测试时将会忽略信息
	//t.Skip()

	data := [][]interface{}{
		[]interface{}{"1", 2, "eree", "dsdsdsd", 8},
		// TODO [BUG-20181220]中文无效
		//[]interface{}{"中国", "贵州", "贵阳", "经开", ".."},
		[]interface{}{"xx", "yyy", "a", "abc", "success"},
		[]interface{}{"a", "bb", "cccccccc", "f", "Joshua"},
		[]interface{}{nil, "I-Heart-Mira", 3.4512, true, 210702},
	}
	t.Logf("\r\n%v", FormatTable(data))
	t.Logf("\r\n%v", FormatTable(data, false))
}

func BenchmarkFormatTable(b *testing.B) {
	//内存分配的基准测试
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//xy = 10*10
		data := [][]interface{}{
			[]interface{}{"1", 2, "eree", "dsdsdsd", 8, 1992, false, math.Pi, math.Phi, 1300_000},
			[]interface{}{"1", 2, "eree", "dsdsdsd", 8, 1992, false, math.Pi, math.Phi, 1300_000},
			[]interface{}{"1", 2, "eree", "dsdsdsd", 8, 1992, false, math.Pi, math.Phi, 1300_000},
			[]interface{}{"1", 2, "eree", "dsdsdsd", 8, 1992, false, math.Pi, math.Phi, 1300_000},
			[]interface{}{"1", 2, "eree", "dsdsdsd", 8, 1992, false, math.Pi, math.Phi, 1300_000},
			// TODO [BUG-20181220]中文无效
			//[]interface{}{"中国", "贵州", "贵阳", "经开", ".."},
			{"xx", "yyy", "a", "abc", "success"},
			{"a", "bb", "cccccccc", "f", "Joshua"},
			{nil, "I-Heart-Mira", 3.4512, true, 210702},
			{nil, "I-Heart-Mira", 3.4512, true, 210702, nil, nil, nil, nil, false},
			{nil, "I-Heart-Mira", 3.4512, true, 210702, nil, nil, nil, nil, false},
		}
		//b.Logf("\r\n%v", FormatTable(data))
		FormatTable(data)
	}
}
