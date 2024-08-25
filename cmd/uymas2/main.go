package main

import (
	"fmt"
	"gitee.com/conero/uymas/v2"
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/logger/lgr"
	"log"
	"runtime"
	"time"
)

// v2 版本临时程序
//
// @todo 后期移除（稳定后）
func main() {
	app := cli.NewCli()
	app.Index(func(cli.ArgsParser) {
		fmt.Println("欢饮您使用 Uymas v2")
		fmt.Println()
		buildInfo := uymas.GetBuildInfo()
		if buildInfo != "" {
			buildInfo = "  " + buildInfo
		}
		fmt.Printf("版本信息 v%s/%s%s\n", uymas.Version, uymas.Release, buildInfo)
		fmt.Printf("build by %s\n", runtime.Version())
	})

	app.Command(func(arg cli.ArgsParser) {
		fmt.Println()
		fmt.Println("参数解析，数据如下")
		fmt.Println()
		fmt.Printf("value: %v\n", arg.Values())
		fmt.Printf("option: %v\n", arg.Option())
		fmt.Printf("CommandList: %v\n", arg.CommandList())
		fmt.Println()
	}, "test", cli.Help("参数解析测试命令"))

	app.Command(func(arg cli.ArgsParser) {
		data := arg.SubCommand()
		if data == "" {
			data = "日志测试工具，" + time.Now().Format(time.DateTime) + "\n 命令格式 log [data]"
		}
		lgr.Trace(data)
		lgr.Debug(data)
		lgr.Warn(data)
		lgr.Error(data)
	}, "log", cli.Help("日志输出测试"))

	app.Command(func(parser cli.ArgsParser) {
		if parser.Switch("simple", "s") {
			fmt.Println("v" + uymas.Version)
			return
		}
		buildInfo := uymas.GetBuildInfo()
		if buildInfo != "" {
			buildInfo = "  " + buildInfo
		}
		fmt.Printf("版本信息 v%s/%s%s\n", uymas.Version, uymas.Release, buildInfo)
		fmt.Printf("build by %s\n", runtime.Version())

	}, "version", cli.Help("版本信息", cli.Option{
		Name:  "simple",
		Help:  "输出简单版本",
		Alias: []string{"s"},
	}))

	err := app.Run()
	if err != nil {
		log.Fatalf("命令行执行错误，%v", err)
	}
}
