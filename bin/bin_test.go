package bin

import (
	"fmt"
	"testing"
)

// @Date：   2018/12/18 0018 16:47
// @Author:  Joshua Conero
// @Name:    bin 测试输出

func TestFormatStr(t *testing.T) {
	s := FormatStr("   ", [][]string{
		[]string{"c", "usage the \"c\""},
		[]string{"mmm", "usage the \"mm\""},
		[]string{"g", "usage the \"g\""},
		[]string{"wwwwww", "usage the \"wwwwww\""},
	})
	fmt.Println(s)
}
