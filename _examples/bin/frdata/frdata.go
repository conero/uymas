package main

import (
	"fmt"
	"github.com/conero/uymas/bin"
)

const (
	Version = "1.0.0"			// 版本
)

func main() {
	var frdata []bin.FRdata

	//结构体注册
	frdata = []bin.FRdata{
		bin.FRdata{
			Cmd:     "version",
			Alias:   nil,
			Todo: CmdVersion,
			Opts: []bin.Option{
				{
					Key:         "verbose",
					Description: "打印详情的版本",
					Logogram:    "v",
				},
			},
		},
		{
			Cmd:     bin.FuncRegisterEmpty,
			Alias:   nil,
			Todo: CmdEmpty,
			OptDick: bin.OptionDick{},
			Opts:    nil,
		},
	}


	//注册并运行
	bin.NewFRdata(frdata)
	bin.Run()
}

func CmdEmpty(a *bin.App)  {
	fmt.Println("新版函数式数据注册实现")
}

func CmdVersion(a *bin.App)  {
	fmt.Print(Version)
}