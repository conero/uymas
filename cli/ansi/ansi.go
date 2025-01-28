// Package ansi color code implementation for command line.
//
// Base (original) library, independent of other packages (project library).
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
//
// 关于\033中的\0在ANSI颜色码中，\033是一个八进制转义序列，代表ASCII字符集中的ESC（Escape）字符。
// 在ASCII表中，ESC字符的值为27（十进制），对应的八进制表示为033。因此，\033实际上是转义序列的一部分，用于引入ANSI颜色码。
//
//	var s = "xxxx"
//	for i, e := range []rune(s) {
//		 fmt.Printf("%d -> %c\n", i, e)
//	}
//	\033 -> esc
package ansi

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	// Black list of text base colors supported by the system library
	Black  = 30
	Red    = 31
	Green  = 32
	Yellow = 33
	Blue   = 34
	Purple = 35
	Cyan   = 36
	White  = 37
	// BlackBr Font Brightening
	BlackBr  = 90
	RedBr    = 91
	GreenBr  = 92
	YellowBr = 93
	BlueBr   = 94
	PurpleBr = 95
	CyanBr   = 96
	WhiteBr  = 97
	// BkgBlack font background color
	BkgBlack  = 40
	BkgRed    = 41
	BkgGreen  = 42
	BkgYellow = 43
	BkgBlue   = 44
	BkgPurple = 45
	BkgCyan   = 46
	BkgWhite  = 47
	// BkgBlackBr font background color and flicker
	BkgBlackBr  = 100
	BkgRedBr    = 101
	BkgGreenBr  = 102
	BkgYellowBr = 103
	BkgBlueBr   = 104
	BkgPurpleBr = 105
	BkgCyanBr   = 106
	BkgWhiteBr  = 107
	// Bold typeface specific style
	Bold      = 1
	Dim       = 2
	Italic    = 3
	Underline = 4
	Twinkle   = 5
	Reverse   = 6
	Hide      = 7
)

func Style(value any, styles ...int) string {
	if value == nil {
		return ""
	}
	if len(styles) == 0 {
		return fmt.Sprintf("%v", value)
	}
	var ansiList []string
	for _, style := range styles {
		ansiList = append(ansiList, fmt.Sprintf("%d", style))
	}

	return fmt.Sprintf("\033[%sm%v\033[0m", strings.Join(ansiList, ";"), value)
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
	// "\033[31mThis is red text.\033[0m"
	//reg := regexp.MustCompile(`((\\{0, 1}033)|(\\{0, 1}x1b))\[\d+(;\d+)*m.*(((\\{0, 1}033)|(\\{0, 1}x1b))\[0m){0, 1}`)
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
