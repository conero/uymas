package main

import (
	"fmt"
	"gitee.com/conero/uymas/bin"
)

func main() {
	app := &bin.App{
		Title:       "uymas-struct-map",
		Description: "结构体映射命令行实例代码",
	}
	app.Append(bin.AppCmd{
		Name:  "test",
		Title: "命令行测试工具！",
		Register: func() {
			fmt.Printf("test 名命令")
		},
	})
	app.Append(bin.AppCmd{
		Name:  "add",
		Title: "数值计算！",
		Register: func(cmd *bin.CliCmd) {
			fmt.Printf("数值计算！")
		},
	})
	app.Run()
}
