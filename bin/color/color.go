// Package color Experimental addition of command line text color characters.
//
// # ANSI escape code
//
// # Attempt to support powershell, bash
//
// ANSI
// # 前景色
// # 30-37 为标准颜色，38;5;编号 为256色模式
// # 示例：黑色、红色、绿色、黄色、蓝色、洋红、青色、白色
//
// Link: https://worktile.com/kb/ask/452398.html, linux命令行输出颜色文本
package color

import "fmt"

const (
	AnsiTextBlack     = 30
	AnsiTextRed       = 31
	AnsiTextGreen     = 32
	AnsiTextYellow    = 33
	AnsiTextBlue      = 34
	AnsiTextPurple    = 35
	AnsiTextCyan      = 36
	AnsiTextWhite     = 37
	AnsiBkgBlack      = 40
	AnsiBkgRed        = 41
	AnsiBkgGreen      = 42
	AnsiBkgYellow     = 43
	AnsiBkgBlue       = 44
	AnsiBkgPurple     = 45
	AnsiBkgCyan       = 46
	AnsiBkgWhite      = 47
	AnsiBoldFont      = 1
	AnsiDimFont       = 2
	AnsiItalicFont    = 3
	AnsiUnderlineFont = 4
	AnsiTwinkleFont   = 5
	AnsiReverseFont   = 6
	AnsiHideFont      = 7
)

func StyleByAnsi(ansi int, value any) string {
	return fmt.Sprintf("\033[%dm%v\033[0m", ansi, value)
}
