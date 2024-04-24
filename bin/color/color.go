// Package color Experimental addition of command line text color characters.
//
// # Attempt to support powershell, bash
//
// ANSI
// # 前景色
// # 30-37 为标准颜色，38;5;编号 为256色模式
// # 示例：黑色、红色、绿色、黄色、蓝色、洋红、青色、白色
package color

import "fmt"

const (
	AnsiRed = 31
)

func ColorByAnsi(ansi int, value any) string {
	return fmt.Sprintf("\\033[%dm%v\\033[0m", ansi, value)
}
