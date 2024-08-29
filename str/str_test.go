package str

import (
	"fmt"
	"testing"
)

// @Date：   2018/12/5 0005 11:03
// @Author:  Joshua Conero
// @Name:    名称描述

func compareStr(expect, real string) bool {
	if expect != real {
		fmt.Println(expect + " VS " + real)
		return false
	}
	return true
}

// 字符串比较
func compareStrFunc(expect func() string, real string) bool {
	s := expect()
	return compareStr(s, real)
}

func TestStr_Lcfirst(t *testing.T) {
	fn := func(s, sT string) {
		if Str(s).ClearSpace() != sT {
			fmt.Println(s + " VS " + sT)
			t.Fail()
		}
	}
	fn(Str("Joshua Conero").Lcfirst(), "joshuaconero")
	fn(Str("JOSHUA").Lcfirst(), "jOSHUA")
	fn(Str("AA BB C D E").Lcfirst(), "aAbBcde")
}

// 项目测试
func TestStr_ClearSpace(t *testing.T) {
	s := " s p a c e"
	sT := "space"
	if Str(s).ClearSpace() != sT {
		t.Fail()
	}
	s = " "
	sT = ""
	if Str(s).ClearSpace() != sT {
		t.Fail()
	}

	s = ""
	sT = ""
	if Str(s).ClearSpace() != sT {
		t.Fail()
	}
}

func TestSplitSafe(t *testing.T) {
	fn := func(src, tar []string) {
		oS := fmt.Sprintf("%v", src)
		oT := fmt.Sprintf("%v", tar)
		//fmt.Println(" 输出格式：", oS)
		if oS != oT {
			t.Fail()
		}
	}
	fn(SplitSafe("t est", ","), []string{"test"})
	fn(SplitSafe("t est, test2, m   3", ","), []string{"test", "test2", "m3"})
	fn(SplitSafe("t est,", ","), []string{"test", ""})
}

func TestRender(t *testing.T) {
	fn := func(dd, real string, data any) {
		if !compareStrFunc(func() string {
			c, _ := Render(dd, data)
			return c
		}, real) {
			t.Fail()
		}
	}
	fn("Tell me your name: {{.name}}", "Tell me your name: Joshua Conero", map[string]string{
		"name": "Joshua Conero",
	})
}

// 反转测试
func TestStr_Reverse(t *testing.T) {
	tc := [][]string{
		{"Joshua Conero.", ".orenoC auhsoJ"},
		{"JoJ", "JoJ"},
		{"", ""},
	}
	for _, ts := range tc {
		ref := ts[0]
		xs := ts[1]
		ys := Str(xs).Reverse()
		if ref != ys {
			t.Fatal(fmt.Sprintf("%s != [Reverse(%s)] => %s", ref, xs, ys))
		}
	}
}

func _strTestCompare(data [][]string, t *testing.T) {
	for _, dd := range data {
		out := dd[1]
		compare := dd[0]
		if compare != out {
			t.Fatal(fmt.Sprintf("%s != %s(compare vs out)", compare, out))
		}
	}
}

func TestPadLeft(t *testing.T) {
	_strTestCompare([][]string{
		[]string{"000001", PadLeft("1", "0", 6)},
		[]string{"-=-=11", PadLeft("11", "-=", 6)},
		[]string{"-*-*-*-ivu", PadLeft("ivu", "*-", 10)},
	}, t)
}

func TestPadRight(t *testing.T) {
	_strTestCompare([][]string{
		[]string{"100000", PadRight("1", "0", 6)},
		[]string{"11-=-=", PadRight("11", "-=", 6)},
		[]string{"ivu*-*-*-*", PadRight("ivu", "*-", 10)},
	}, t)
}

func TestStr_LowerStyle(t *testing.T) {
	// Case 1
	vStr := "FirstName"
	rStr := "first_name"
	gStr := Str(vStr).LowerStyle()
	if gStr != rStr {
		t.Fatalf("%v --> %v VS %v", vStr, gStr, rStr)
	}

	// Case 2
	vStr = "getHeightWidthRate"
	rStr = "get_height_width_rate"
	gStr = Str(vStr).LowerStyle()
	if gStr != rStr {
		t.Fatalf("%v --> %v VS %v", vStr, gStr, rStr)
	}

	// Case 2
	vStr = "_stringIsLowerStyleAndNeedTrimWithoutFuncButFieldIsAlpha2Email0519"
	rStr = "_string_is_lower_style_and_need_trim_without_func_but_field_is_alpha2_email0519"
	gStr = Str(vStr).LowerStyle()
	if gStr != rStr {
		t.Fatalf("%v --> %v VS %v", vStr, gStr, rStr)
	}
}

// `first_name` 			-> `FirstName`,
// `get_height_width_rate` 	-> `GetHeightWidthRate`
func TestStr_CamelCase(t *testing.T) {
	// Case 1
	vStr := "first_name"
	rStr := "FirstName"
	gStr := Str(vStr).CamelCase()
	if gStr != rStr {
		t.Fatalf("%v --> %v VS %v", vStr, gStr, rStr)
	}

	// Case 2
	vStr = "get_height_width_rate"
	rStr = "GetHeightWidthRate"
	gStr = Str(vStr).CamelCase()
	if gStr != rStr {
		t.Fatalf("%v --> %v VS %v", vStr, gStr, rStr)
	}

}

func TestStr_Unescape(t *testing.T) {
	line := `空格字符测试\sj\so\ss\sh\su\sa\s.`
	ref := `空格字符测试 j o s h u a .`
	rel := Str(line).Unescape()

	// case
	testFn := func() {
		if ref != rel {
			t.Errorf("解析错误：%#v ≠ %#v", rel, ref)
		}
	}
	testFn()

	//case
	line = `换行符 j\nc.`
	ref = "换行符 j\nc."
	rel = Str(line).Unescape()
	testFn()

	//case
	line = `\s`
	ref = " "
	rel = Str(line).Unescape()
	testFn()

	//case
	line = `\n`
	ref = "\n"
	rel = Str(line).Unescape()
	testFn()
}

func BenchmarkRandString_SafeStr(b *testing.B) {
	b.ResetTimer()
	bit := 35
	for i := 0; i < b.N; i++ {
		if i < 2 {
			continue
		}
		rss := repeatRandStringSafeStr(bit, i, nil)
		b.Logf("重复率未满足百分之百 => %v/%v，重复率：%.4f", rss.uniqueCount, rss.Max, rss.Rate)
	}
}

func TestRandString_SafeStr(t *testing.T) {
	// case1
	rss := repeatRandStringSafeStr(35, 100, func(s string) {
		t.Logf("%v", s)
	})
	if rss.uniqueCount != rss.Max {
		t.Fatalf("重复率未满足百分之百 => %v/%v，重复率：%.4f", rss.uniqueCount, rss.Max, rss.Rate)
	}

	// case2
	rss = repeatRandStringSafeStr(50, 500, nil)
	if rss.uniqueCount != rss.Max {
		t.Fatalf("重复率未满足百分之百 => %v/%v，重复率：%.4f", rss.uniqueCount, rss.Max, rss.Rate)
	}
}

type repeatRsss struct {
	uniqueCount int
	Max         int
	Rate        float64
}

// 安全数随机生成
func repeatRandStringSafeStr(bit, max int, scanFn func(s string)) repeatRsss {
	var lastStrMap = map[string]int{}
	for i := 0; i < max; i++ {
		ss := RandStr.SafeStr(bit)
		lastStrMap[ss] = 1
		if scanFn != nil {
			scanFn(fmt.Sprintf("%v => %v", i, ss))
		}
	}

	uniqueCtt := len(lastStrMap)
	return repeatRsss{
		uniqueCount: uniqueCtt,
		Max:         max,
		Rate:        float64(max-uniqueCtt) / float64(max),
	}
}
