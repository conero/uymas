package main

import (
	"fmt"
	"gitee.com/conero/uymas/v2"
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/cli/ansi"
	"gitee.com/conero/uymas/v2/cli/chest"
	"gitee.com/conero/uymas/v2/cli/evolve"
	"gitee.com/conero/uymas/v2/cli/gen"
	"gitee.com/conero/uymas/v2/logger/lgr"
	"gitee.com/conero/uymas/v2/number"
	"gitee.com/conero/uymas/v2/str"
	"gitee.com/conero/uymas/v2/util/fs"
	"gitee.com/conero/uymas/v2/util/tm"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"time"
)

type testArgs struct {
	Port int    `cmd:"port,p default:8080 help:设置端口号"`
	Addr string `cmd:"addr,a default::9000 help:设置外部请求地址"`
	Host string `cmd:"host,h required help:设置服务地址"`
}

type test struct {
	evolve.Command
}

func (c *test) Demo() {
	fmt.Println("test demo, ha!")
	fmt.Println()
	fmt.Println("rootPath: " + fs.RootPath())
	fmt.Println("rootApp: " + fs.AppName())
}

func (c *test) Test() {
	spendFn := tm.SpendFn()
	arg := c.X.Args
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
		lgr.Info("消耗时间：%s\n", spendFn())
		return
	}

	fmt.Println()
	fmt.Printf("消耗时间：%s\n", spendFn())
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

func (c *test) Arg() {
	param := testArgs{}
	err := gen.ArgsDress(c.X.Args, &param)
	if err != nil {
		lgr.Info("解析值错误，%v", err)
		return
	}

	lgr.Info("解析到的值如下：\n%#v", param)

}

func main() {
	argsConfig := cli.DefArgsConfig
	key := "UYMAS_CLI_SHORT"
	cliShort := os.Getenv(key)

	argsConfig.ShortOption = cliShort != ""
	args := cli.NewArgsWith(argsConfig)
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
		fmt.Println()
		fmt.Println("环境变量 " + key + "=true   用于设置使其支持，短选项")
	}

	evl.Index(testCmd)
	evl.Command(testCmd, "index", cli.Help("索引测试命令"))
	testArgsOpts, _ := gen.ArgsDecompose(testArgs{})
	evl.CommandList(new(test), []string{"test", "t"},
		cli.HelpSub("命令测试",
			cli.Help("测试工具",
				cli.Option{Name: "verbose", Alias: []string{"V"}, Help: "详细输出内容"},
				cli.OptionHelp("设置需读取选项名称", "option", "O"),
				cli.OptionHelp("生成用于测试的命令选项数", "make-number", "M"),
				cli.Option{}).NameAlias("test"),
			cli.Help("命令示例").NameAlias("demo"),
			cli.Help("参数值测试", testArgsOpts...).NameAlias("arg"),
		))
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
		Alias: []string{"data", "d"},
		Help:  "设置日志输出内容",
	}))

	evl.Command(func(arg evolve.Param) {
		flName := arg.Args.Get("file", "f")
		fi, err := os.Stat(flName)
		if err != nil {
			lgr.Error("文件读取错误，%s", err)
			return
		}
		lgr.Info("文件读取成功，主要信息如下：\n"+
			"文件大小：%s\n"+"mode：%s\n"+"修改日期：%s",
			number.Bytes(fi.Size()), fi.Mode(), fi.ModTime().Format(time.DateTime))
	}, "stat", cli.Help("文件信息查看",
		cli.Option{
			Name:    "file",
			Alias:   []string{"f"},
			Require: true,
			Help:    "指定文件名称",
		},
	))

	evl.Lost(func(arg cli.ArgsParser) {
		name := arg.Command()
		raw := args.Raw()
		raw = raw[1:]
		output, isSearch, err := chest.CmdSearchRun(name, raw, "bin/")
		if isSearch {
			if err != nil {
				lgr.Error(err.Error())
				return
			}
			lgr.Info("子级命令搜索到，输入内容如下：")
			fmt.Println(output)
			return
		}

		lgr.Warn("%s: 命令不存在", name)
	})

	//evl.Run("test", "demo")
	err := evl.RunArgs(args)
	if err != nil {
		lgr.Error(err.Error())
	}
}
