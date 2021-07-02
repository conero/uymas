// Package test unit.
package unit

import (
	"fmt"
)

// @Date：   2018/12/12 0012 14:29
// @Author:  Joshua Conero
// @Name:    单元测试

// StrSingLine reversal test
// data [][]string{expected, input, output}
// bool 表通过与否； string 错误消息
// args[0]  = "{expected} {input}, {output}  "
func StrSingLine(data [][]string, args ...string) interface{} {
	for _, ts := range data {
		expected := ts[0]
		input := ts[1]
		output := ts[2]
		if expected != output {
			if args != nil && len(args) > 0 {
				return fmt.Sprintf(args[0], expected, input, output)
			} else {
				return false
			}
		}
	}
	return true
}
