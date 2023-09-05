package main

import (
	"fmt"
	"gitee.com/conero/uymas/bin"
	"gitee.com/conero/uymas/logger/lgr"
	"gitee.com/conero/uymas/str"
	"regexp"
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
