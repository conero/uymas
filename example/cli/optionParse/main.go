// option parse for cli
package main

import (
	"fmt"
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/cli/gen"
	"gitee.com/conero/uymas/v2/logger/lgr"
)

type optionParse struct {
	Addr  string   `cmd:"addr,a default::12409 help:请求监听地址"`
	File  string   `cmd:"file,fl,f required help:输出文件指定文件"`
	Age   int16    `cmd:"age,a default:18 help:设置年龄，这是整形数展示"`
	Rage  float32  `cmd:"rage,r default:0.3145 help:设置占用比例，浮点数实例"`
	Off   bool     `cmd:"off default:True help:默认关闭，bool类型展示"`
	Index []string `cmd:"index default:[index.html,index.htm] help:服务器支持索引文件"`
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
		fmt.Println()

	}, "option", cli.Help("选项测试", gen.ArgsDecomposeMust(optionParse{})...))
	lgr.ErrorIf(app.Run())
}
