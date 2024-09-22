package main

import (
	"fmt"
	"gitee.com/conero/uymas/v2"
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/cli/ansi"
	"gitee.com/conero/uymas/v2/cli/gen"
	"gitee.com/conero/uymas/v2/culture/ganz"
	"gitee.com/conero/uymas/v2/logger/lgr"
	"log"
	"runtime"
	"strconv"
	"time"
)

type globalOption struct {
	IsVerbose bool `cmd:"verbose,vv help:是否详细输出"`
}

func cmdGanz(args cli.ArgsParser) {
	year := args.SubCommand()
	if year == "" {
		lgr.Info("请输入年份，来计算干支纪元法。默认为当年")
	}

	y, _ := strconv.Atoi(year)
	if y <= 0 {
		y = time.Now().Year()
	}

	gz, zod := ganz.CountGzAndZodiac(y)

	fmt.Printf("  %d年，干支纪元%s年，属%s.\n", y, gz, zod)
	fmt.Printf("\n天干：%s\n地支：%s\n属相：%s\n",
		ansi.Style(ganz.TianGan, ansi.CyanBr),
		ansi.Style(ganz.DiZhi, ansi.CyanBr),
		ansi.Style(ganz.DiZhi, ansi.CyanBr))
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

	app.CommandList(cmdPinyin, []string{"pinyin", "py"},
		cli.Help("生成汉字拼音", gen.ArgsDecomposeMust(pinyinOption{})...))

	app.Command(cmdCalc, "cal", cli.Help("等式计算器"))
	app.Command(cmdGanz, "ganz", cli.Help("计算给定年的干支纪元，默认今年"))
	app.Command(cmdHash, "hash", cli.Help("读取指定文件hash码列表", gen.ArgsDecomposeMust(cmdHashOpt{})...))
	app.CommandList(cmdDigit, []string{"digit", "dg"},
		cli.Help("阿拉伯数字转中文（默认大写）", gen.ArgsDecomposeMust(digitOption{})...))
	app.CommandList(cmdDatediff, []string{"datediff", "dd"},
		cli.Help("日期运算", gen.ArgsDecomposeMust(ddOption{})...))
	app.End(func(cli.ArgsParser) {
		fmt.Println()
	})
	err := app.Run()
	if err != nil {
		log.Fatalf("命令行执行错误，%v", err)
	}
}
