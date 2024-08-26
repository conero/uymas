package main

import (
	"fmt"
	"gitee.com/conero/uymas/v2"
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/logger/lgr"
	"gitee.com/conero/uymas/v2/str"
	"gitee.com/conero/uymas/v2/util/tm"
	"log"
	"math/rand"
	"runtime"
	"strings"
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
		spendFn := tm.SpendFn()
		if arg.Switch("verbose", "V") {
			fmt.Println()
			fmt.Println("参数解析，数据如下")
			fmt.Println()
			fmt.Printf("value: %v\n", arg.Values())
			fmt.Printf("option: %v\n", arg.Option())
			fmt.Printf("CommandList: %v\n", arg.CommandList())
		}
		option := arg.List("option", "O")
		if len(option) > 0 {
			fmt.Printf("Read option: %v\n", arg.Get(option...))
		}

		vNumber := arg.Int("make-number", "M")
		if vNumber > 0 {
			var mkOptionList = []string{"uymas", "test"}
			for i := 0; i < vNumber; i++ {
				mkKey := str.RandStr.SafeStr(rand.Intn(40))
				mkQueue := []string{"--" + mkKey}
				if rand.Intn(4)%2 == 0 {
					mkQueue = append(mkQueue, fmt.Sprintf("%d", rand.Intn(999999)))
				}
				mkOptionList = append(mkOptionList, mkQueue...)
			}
			lgr.Info("创建生成测试命令如下：\n%s", strings.Join(mkOptionList, " "))
			fmt.Printf("消耗时间：%s\n", spendFn())
			return
		}

		fmt.Println()
		lgr.Info("使用时间：%s\n", spendFn())
	}, "test", cli.Help("参数解析测试命令",
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
	))

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

	err := app.Run()
	if err != nil {
		log.Fatalf("命令行执行错误，%v", err)
	}
}
