package main

import (
	"fmt"
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/logger/lgr"
	"gitee.com/conero/uymas/v2/number"
	"gitee.com/conero/uymas/v2/rock"
	"os"
	"time"
)

func demoCmd(parser cli.ArgsParser) {
	fmt.Println("Hello, demo.")
	vData := parser.Get("data", "d")
	if vData != "" {
		lgr.Info("输入数据：%s", vData)
	}
	list := parser.List("list", "l")
	if len(list) > 0 {
		lgr.Info("输入的列表数据层如：\n%s", rock.FormatList(list))
	}
}

func statCmd(parser cli.ArgsParser) {
	flName := parser.Get("file", "f")
	fi, err := os.Stat(flName)
	if err != nil {
		lgr.Error("文件读取错误，%s", err)
		return
	}
	lgr.Info("文件读取成功，主要信息如下：\n"+
		"文件大小：%s\n"+"mode：%s\n"+"修改日期：%s",
		number.Bytes(fi.Size()), fi.Mode(), fi.ModTime().Format(time.DateTime))

}

func main() {
	app := cli.NewCli()
	app.Command(demoCmd, "demo", cli.Help("示例命令工具柜",
		cli.Option{
			Name:  "data",
			Alias: []string{"d"},
			Help:  "支持字符串数据输入",
		},
		cli.Option{
			Name:  "list",
			Alias: []string{"l"},
			Help:  "支持列表数据支持",
		},
	))
	app.Command(statCmd, "stat", cli.Help("文件信息查看",
		cli.Option{
			Name:    "file",
			Alias:   []string{"f"},
			Require: true,
			Help:    "指定文件名称",
		},
	))
	_ = app.Run()
}