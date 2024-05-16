package main

import (
	"fmt"
	"gitee.com/conero/uymas/bin"
	"gitee.com/conero/uymas/bin/color"
	"gitee.com/conero/uymas/logger"
	"gitee.com/conero/uymas/logger/lgr"
	"gitee.com/conero/uymas/str"
	"os"
	"regexp"
	"runtime"
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
	fmt.Println()
	fmt.Println("color [text] 文本颜色码测试，不设置时默认")
	fmt.Println("  -v,--value       指定样式默认时为红色")
	fmt.Println("  -r,--raw         原始输出，用于原始命令控制如清屏之类")
	fmt.Println("  -m,--multi [..]  设置多种样式合并效果，组合样式")
	fmt.Println()
	fmt.Println("log [text]  日志测试")
	fmt.Println("  -l,--level 日志级别设置")

}

// Cal @todo 下一版本进行删除（next-remove）
//
// Deprecated: next major version remove
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
	rawArgs := []string{"raw", "r"}
	if c.Cc.CheckSetting(rawArgs...) {
		argStr := c.Cc.ArgRaw(rawArgs...)
		if argStr != "" {
			text = argStr
		}
		if text == "" {
			lgr.Info("请输入内容先！")
			return
		}
		fmt.Printf("\033[%s", text)
		fmt.Println()
		return
	}
	if text == "" {
		text = ":)- It's demo test text, default.\n    " + time.Now().Format(time.RFC3339)
	}

	multi := c.Cc.ArgIntSlice("multi", "m")

	if len(multi) > 0 {
		fmt.Println(color.StyleByAnsiMulti(text, multi...))
		return
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

func (c *defaultApp) Log() {
	text := c.Cc.SubCommand
	if text == "" {
		text = "日志测试文本，可输入内容显示不同内容。" + time.Now().Format(time.RFC3339) +
			"。\n    可通过 UYMAS_LGR_LEVEL 环境变量设置 lgr 的日志级别"
	}

	level := c.Cc.ArgRaw("l", "level")
	var lg *logger.Logger
	if level == "" {
		tmpLog := lgr.Log()
		lg = &tmpLog
	} else {
		lg = logger.NewLogger(logger.Config{
			Level: level,
		})
	}

	if level != "" {
		lg.Infof("当前设置 level 参数: %s", level)
	}

	lg.Errorf(text)
	lg.Warnf(text)
	lg.Infof(text)
	lg.Infof("go version: %s, Os: %s, pid: %d, ppid: %d", runtime.Version(), runtime.GOOS, os.Getpid(), os.Getppid())
	lg.Debugf(text)
	lg.Tracef(text)

	lg.Debugf("os args: %#v", os.Args)
	lg.Tracef("Getpid: %v, uid: %v", os.Getpid(), os.Getuid())
	fmt.Println()
}
