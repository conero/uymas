package main

import (
	"fmt"
	"gitee.com/conero/uymas/v2"
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/cli/gen"
	"gitee.com/conero/uymas/v2/logger/lgr"
	"log"
	"runtime"
	"time"
)

type globalOption struct {
	IsVerbose bool `cmd:"verbose,vv help:是否详细输出"`
}

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

	app.Command(cmdTest, "test", cli.Help("参数解析测试命令",
		cli.Option{
			Alias: []string{"option", "O"},
			Help:  "测试读取选项值",
		},
		cli.Option{
			Alias: []string{"verbose", "V"},
			Help:  "是否详细输出",
		},
		cli.Option{
			Alias: []string{"make-number", "M"},
			Help:  "创建的个数，用于测试",
		},
	).NoValid())

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
		Name:    "simple",
		Help:    "输出简单版本",
		Require: false,
		Alias:   []string{"s"},
	}))

	app.End(func(cli.ArgsParser) {
		fmt.Println()
	})

	app.CommandList(cmdPinyin, []string{"pinyin", "py"},
		cli.Help("生成汉字拼音", gen.ArgsDecomposeMust(pinyinOption{})...))

	err := app.Run()
	if err != nil {
		log.Fatalf("命令行执行错误，%v", err)
	}
}
