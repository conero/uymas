package bin

import (
	"fmt"
	"testing"
)

// @Date：   2018/12/18 0018 16:47
// @Author:  Joshua Conero
// @Name:    bin 测试输出

func TestFormatStr(t *testing.T) {
	// 用于输出格式
	// 实际测试时将会忽略信息
	t.Skip()

	s := FormatStr("   ", [][]string{
		[]string{"c", "usage the \"c\""},
		[]string{"mmm", "usage the \"mm\""},
		[]string{"g", "usage the \"g\""},
		[]string{"wwwwww", "usage the \"wwwwww\""},
	})
	fmt.Println(s)
}

func TestFormatQue(t *testing.T) {
	// 用于输出格式
	// 实际测试时将会忽略信息
	t.Skip()

	data := []interface{}{}
	data = []interface{}{
		"中文", "letter", "letter", "letter", "letter", "letter", "letter",
		"中文", "letter", "letter", "letter", "letter", "letter", "letter",
		"中文", "letter", "letter", "letter", "letter", "letter", "letter",
	}
	s := FormatQue(data)
	fmt.Println(s)

	s = FormatQue(data, " .", "-")
	fmt.Println(s)
}

func TestFormatTable(t *testing.T) {
	// 用于输出格式
	// 实际测试时将会忽略信息
	//t.Skip()

	data := [][]interface{}{
		[]interface{}{"1", "2", "eree", "dsdsdsd", "8"},
		// TODO [BUG-20181220]中文无效
		//[]interface{}{"中国", "贵州", "贵阳", "经开", ".."},
		[]interface{}{"xx", "yyy", "a", "abc", "success"},
		[]interface{}{"a", "bb", "cccccccc", "f", "Joshua"},
	}
	fmt.Println(FormatTable(data))
}
