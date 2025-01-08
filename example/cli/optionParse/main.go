// option parse for cli
package main

import (
	"fmt"
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/cli/gen"
	"gitee.com/conero/uymas/v2/logger/lgr"
	"gitee.com/conero/uymas/v2/rock"
	"gitee.com/conero/uymas/v2/util"
)

type optionParse struct {
	Addr   string   `cmd:"addr,a default::12409 help:请求监听地址"`
	File   string   `cmd:"file,fl,f required help:输出文件指定文件"`
	Age    int16    `cmd:"age,a default:18 help:设置年龄，这是整形数展示"`
	Rage   float32  `cmd:"rage,r default:0.3145 help:设置占用比例，浮点数实例"`
	Off    bool     `cmd:"off default:True help:默认关闭，bool类型展示"`
	Index  []string `cmd:"index default:[index.html,index.htm] help:服务器支持索引文件"`
	Number []uint16 `cmd:"number,N default:[52,26,27] help:uint16切片类型测试"`
	optionHelpPlus
}

type subOptionX struct {
	Name    string `cmd:"name help:设置姓名"`
	Age     int16  `cmd:"age default:42 help:设置年纪"`
	IsCheck bool   `cmd:"check help:是否需要检查"`
}

type subOption struct {
	X subOptionX `cmd:"x structGen help:附加参数选项"`
}

type optionHelpPlus struct {
	NoHelp bool `cmd:"no-help help:- default:true"`
	IsHelp bool `cmd:"is-help default:false help:可设置no-help进行测试"`
}

func main() {
	app := cli.NewCli()

	app.Command(func(parser cli.ArgsParser) {
		var opt optionParse
		err := gen.ArgsDress(parser, &opt)
		if err != nil {
			lgr.Error(err.Error())
			return
		}

		fmt.Println("option 解析，各种只是得数据类型测试")
		fmt.Println()
		lgr.Info("地址: %s", opt.Addr)
		lgr.Info("年纪: %d", opt.Age)
		lgr.Info("文件: %s", opt.File)
		lgr.Info("占比: %f", opt.Rage)
		lgr.Info("是否关闭: %#v", opt.Off)
		lgr.Info("服务索引: %#v", opt.Index)
		lgr.Info("切片 u16: %#v", opt.Number)
		fmt.Println()
		lgr.Info("all option:\n%s", rock.FormatKv(util.StructToMap(opt)))

	}, "option", cli.Help("选项测试", gen.ArgsDecomposeMust(optionParse{})...))

	app.Command(func(parser cli.ArgsParser) {
		var opt subOption
		err := gen.ArgsDress(parser, &opt)
		if err != nil {
			lgr.Error(err.Error())
			return
		}
		fmt.Println("子选项测试: ")
		fmt.Printf("%#v\n", opt)
	}, "sub", cli.Help("命令子选项支持", gen.ArgsDecomposeMust(subOption{})...))
	lgr.ErrorIf(app.Run())
}
