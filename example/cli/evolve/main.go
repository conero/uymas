package main

import (
	"fmt"
	"gitee.com/conero/uymas/v2"
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/cli/ansi"
	"gitee.com/conero/uymas/v2/cli/chest"
	"gitee.com/conero/uymas/v2/cli/evolve"
	"gitee.com/conero/uymas/v2/logger/lgr"
	"gitee.com/conero/uymas/v2/util/fs"
	"runtime"
	"time"
)

type test struct {
	evolve.Command
}

func (c *test) Demo() {
	fmt.Println("test demo, ha!")
	fmt.Println()
	fmt.Println("rootPath: " + fs.RootPath())
	fmt.Println("rootApp: " + fs.AppName())
}

func (c *test) DefHelp() {
	fmt.Println("测试命令，全局选项如：")
	fmt.Println(" --cmd [command]          判断命令是否存在")
	fmt.Println()
	fmt.Println(" demo                     普通测试用例")
}

func (c *test) DefIndex() {
	command := c.X.Args.Get("cmd")
	if command != "" {
		exist, err := chest.CmdExist(command)
		if err != nil {
			fmt.Printf("命令行检测错误，%s\n", err)
		} else if exist {
			fmt.Printf("%s：命令存在\n", command)
		} else {
			fmt.Printf("%s：%s\n", command, ansi.Style("命令不存在", ansi.Red))
		}
		return
	}
	fmt.Println("您在使用 test 命令")
}

func main() {
	evl := evolve.NewEvolve()
	testCmd := func() {
		fmt.Println("Evolution For Index.")
		fmt.Println()
		buildInfo := uymas.GetBuildInfo()
		if buildInfo != "" {
			buildInfo = "  " + buildInfo
		}
		fmt.Printf("版本信息 v%s/%s%s\n", uymas.Version, uymas.Release, buildInfo)
		fmt.Printf("build by %s\n", runtime.Version())
	}

	evl.Index(testCmd)
	evl.Command(testCmd, "index", cli.Help("索引测试命令"))
	evl.CommandList(new(test), []string{"test", "t"}, cli.Help("命令测试"))
	evl.Command(func(arg evolve.Param) {
		data := arg.Args.SubCommand()
		if data == "" {
			data = arg.Args.Get("data", "d")
		}
		if data == "" {
			data = "日志测试工具，" + time.Now().Format(time.DateTime) + "\n 命令格式 log [data]"
		}
		lgr.Trace(data)
		lgr.Debug(data)
		lgr.Warn(data)
		lgr.Error(data)
	}, "log", cli.Help("日志测试工具", cli.Option{
		Name:  "name",
		Alias: []string{"d"},
		Help:  "设置日志输出内容",
	}))
	//evl.Run("test", "demo")
	evl.Run()
}
