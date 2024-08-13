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
// 格式： `\033[<parameter1>;<parameter2>...<parameterN><letter>`
//
// Link: https://worktile.com/kb/ask/452398.html, linux命令行输出颜色文本
//
// Link: http://blog.lujinkai.cn/%E8%BF%90%E7%BB%B4/%E5%9F%BA%E7%A1%80/%E6%96%87%E6%9C%AC%E5%A4%84%E7%90%86/ANSI%E8%BD%AC%E4%B9%89%E5%BA%8F%E5%88%97/
package color

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	AnsiTextBlack  = 30
	AnsiTextRed    = 31
	AnsiTextGreen  = 32
	AnsiTextYellow = 33
	AnsiTextBlue   = 34
	AnsiTextPurple = 35
	AnsiTextCyan   = 36
	AnsiTextWhite  = 37
	// AnsiTextBlackBr Font Brightening
	AnsiTextBlackBr   = 90
	AnsiTextRedBr     = 91
	AnsiTextGreenBr   = 92
	AnsiTextYellowBr  = 93
	AnsiTextBlueBr    = 94
	AnsiTextPurpleBr  = 95
	AnsiTextCyanBr    = 96
	AnsiTextWhiteBr   = 97
	AnsiBkgBlack      = 40
	AnsiBkgRed        = 41
	AnsiBkgGreen      = 42
	AnsiBkgYellow     = 43
	AnsiBkgBlue       = 44
	AnsiBkgPurple     = 45
	AnsiBkgCyan       = 46
	AnsiBkgWhite      = 47
	AnsiBkgBlackBr    = 100
	AnsiBkgRedBr      = 101
	AnsiBkgGreenBr    = 102
	AnsiBkgYellowBr   = 103
	AnsiBkgBlueBr     = 104
	AnsiBkgPurpleBr   = 105
	AnsiBkgCyanBr     = 106
	AnsiBkgWhiteBr    = 107
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

func StyleByAnsiMulti(value any, ansi ...int) string {
	var aList []string
	for _, a := range ansi {
		aList = append(aList, fmt.Sprintf("%d", a))
	}
	return fmt.Sprintf("\033[%sm%v\033[0m", strings.Join(aList, ";"), value)
}

func StyleByAnsiString(ansi string, value any) string {
	return fmt.Sprintf("\033[%sm%v\033[0m", ansi, value)
}

// AnsiClear clear ansi color code from text string
//
// Format: `\033[<parameter1>;<parameter2>...<parameterN><letter>`
func AnsiClear(ansiColor string) string {
	clearFn := AnsiClearFn()
	return clearFn(ansiColor)
}

// AnsiClearFn clear ansi with anonymous function
func AnsiClearFn() func(string) string {
	reg := regexp.MustCompile(`.*\033\[\d+(;\d+)*m.*(\033\[0m).*`)
	headReg := regexp.MustCompile(`\033\[\d+(;\d+)*m`)
	endReg := regexp.MustCompile(`\033\[0m`)

	return func(s string) string {
		if s == "" {
			return ""
		}
		if reg.MatchString(s) {
			s = headReg.ReplaceAllString(s, "")
			s = endReg.ReplaceAllString(s, "")
		}
		return s
	}
}
