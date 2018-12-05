package str

import (
	"fmt"
	"testing"
)

// @Date：   2018/12/5 0005 11:03
// @Author:  Joshua Conero
// @Name:    名称描述

// 项目测试
func TestClearSpace(t *testing.T) {
	s := " s p a c e"
	sT := "space"
	if ClearSpace(s) != sT{
		t.Fail()
	}
	s = " "
	sT = ""
	if ClearSpace(s) != sT{
		t.Fail()
	}

	s = ""
	sT = ""
	if ClearSpace(s) != sT{
		t.Fail()
	}
}


func TestSplitSafe(t *testing.T) {
	fn := func(src, tar []string) {
		oS := fmt.Sprintf("%v", src)
		oT := fmt.Sprintf("%v", tar)
		fmt.Println(" 输出格式：", oS)
		if oS != oT{
			t.Fail()
		}
	}
	fn(SplitSafe("t est", ","), []string{"test"})
	fn(SplitSafe("t est, test2, m   3", ","), []string{"test", "test2", "m3"})
	fn(SplitSafe("t est,", ","), []string{"test", ""})
}