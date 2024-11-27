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
	TextBlack  = 30
	TextRed    = 31
	TextGreen  = 32
	TextYellow = 33
	TextBlue   = 34
	TextPurple = 35
	TextCyan   = 36
	TextWhite  = 37
	// TextBlackBr Font Brightening
	TextBlackBr   = 90
	TextRedBr     = 91
	TextGreenBr   = 92
	TextYellowBr  = 93
	TextBlueBr    = 94
	PurpleBr      = 95
	TextCyanBr    = 96
	TextWhiteBr   = 97
	BkgBlack      = 40
	BkgRed        = 41
	BkgGreen      = 42
	BkgYellow     = 43
	BkgBlue       = 44
	BkgPurple     = 45
	BkgCyan       = 46
	BkgWhite      = 47
	BkgBlackBr    = 100
	BkgRedBr      = 101
	BkgGreenBr    = 102
	BkgYellowBr   = 103
	BkgBlueBr     = 104
	BkgPurpleBr   = 105
	BkgCyanBr     = 106
	BkgWhiteBr    = 107
	BoldFont      = 1
	DimFont       = 2
	ItalicFont    = 3
	UnderlineFont = 4
	TwinkleFont   = 5
	ReverseFont   = 6
	HideFont      = 7
)

func Style(ansi int, value any) string {
	return fmt.Sprintf("\033[%dm%v\033[0m", ansi, value)
}

func Styles(value any, ansi ...int) string {
	var aList []string
	for _, a := range ansi {
		aList = append(aList, fmt.Sprintf("%d", a))
	}
	return fmt.Sprintf("\033[%sm%v\033[0m", strings.Join(aList, ";"), value)
}

func StyleString(ansi string, value any) string {
	return fmt.Sprintf("\033[%sm%v\033[0m", ansi, value)
}

// Clear ansi color code from text string
//
// Format: `\033[<parameter1>;<parameter2>...<parameterN><letter>`
func Clear(ansiColor string) string {
	clearFn := ClearFn()
	return clearFn(ansiColor)
}

// ClearFn clear ansi with anonymous function
func ClearFn() func(string) string {
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
