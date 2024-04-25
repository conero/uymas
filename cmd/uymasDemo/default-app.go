package main

import (
	"fmt"
	"gitee.com/conero/uymas/bin"
	"gitee.com/conero/uymas/bin/color"
	"gitee.com/conero/uymas/logger/lgr"
	"gitee.com/conero/uymas/str"
	"regexp"
	"time"
)

type defaultApp struct {
	bin.CliApp
}

func (c *defaultApp) DefaultIndex() {
	fmt.Printf("demo 子命令")
}

func (c *defaultApp) DefaultHelp() {
	fmt.Println("cal [equal]  计算器")
	fmt.Println("  -V,--verbose     详细显示")
	fmt.Println("color [text] 文本颜色码测试，不设置时默认")
	fmt.Println("  -v,--value       指定样式默认时为红色")
}

func (c *defaultApp) Cal() {
	equal := c.Cc.SubCommand
	if equal != "" {
		spanReg := regexp.MustCompile(`\s+`)
		equal = spanReg.ReplaceAllString(equal, "")
	}
	if equal == "" {
		lgr.Error("请输入等式符号！")
		return
	}

	calc := str.NewCalc(equal)
	calc.Count()

	if c.Cc.CheckSetting("V", "verbose") {
		lgr.Info("输入等式：%s\n    => %v", c.Cc.SubCommand, calc)
		return
	}
	fmt.Println(calc)
}

// Color 测试命令行文本颜色
func (c *defaultApp) Color() {
	text := c.Cc.SubCommand
	if text == "" {
		text = ":)- It's demo test text, default.\n    " + time.Now().Format(time.RFC3339)
	}
	value := c.Cc.ArgInt("value", "v")
	if value < 1 {
		value = color.AnsiTextRed
	}

	fmt.Println(color.StyleByAnsi(value, text))
}

func (c *defaultApp) DefaultUnmatched() {
	lgr.Error("命令 %v 不存在", c.Cc.Command)
}
