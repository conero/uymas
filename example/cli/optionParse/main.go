// option parse for cli
package main

import (
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/cli/gen"
	"gitee.com/conero/uymas/v2/logger/lgr"
)

type optionParse struct {
	Addr string `cmd:"addr,a default::12409 help:请求监听地址"`
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

		lgr.Info("地址: %s", opt.Addr)

	}, "option", cli.Help("选项测试", gen.ArgsDecomposeMust(optionParse{})...))
	lgr.ErrorIf(app.Run())
}
