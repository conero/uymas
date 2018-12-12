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

func TestUcfirst(t *testing.T) {
	fn := func(s, sT string) {
		if ClearSpace(s) != sT {
			fmt.Println(s + " VS " + sT)
			t.Fail()
		}
	}
	fn(Ucfirst(" i am joshua conero"), "IAmJoshuaConero")
	fn(Ucfirst(" joshuaConero"), "JoshuaConero")
	fn(Ucfirst(" test "), "Test")
	fn(Ucfirst(" tEST "), "TEST")
}

func TestLcfirst(t *testing.T) {
	fn := func(s, sT string) {
		if ClearSpace(s) != sT {
			fmt.Println(s + " VS " + sT)
			t.Fail()
		}
	}
	fn(Lcfirst("Joshua Conero"), "joshuaconero")
	fn(Lcfirst("JOSHUA"), "jOSHUA")
	fn(Lcfirst("AA BB C D E"), "aAbBcde")
}

// 项目测试
func TestClearSpace(t *testing.T) {
	s := " s p a c e"
	sT := "space"
	if ClearSpace(s) != sT {
		t.Fail()
	}
	s = " "
	sT = ""
	if ClearSpace(s) != sT {
		t.Fail()
	}

	s = ""
	sT = ""
	if ClearSpace(s) != sT {
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
	fn := func(dd, real string, data interface{}) {
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
func TestReverse(t *testing.T) {
	tc := [][]string{
		[]string{"Joshua Conero.", ".orenoC auhsoJ"},
		[]string{"JoJ", "JoJ"},
		[]string{"", ""},
	}
	for _, ts := range tc {
		ref := ts[0]
		xs := ts[1]
		ys := Reverse(xs)
		if ref != ys {
			t.Fatal(fmt.Sprintf("%s != [Reverse(%s)] => %s", ref, xs, ys))
		}
	}
}
