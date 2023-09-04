package main

import (
	"fmt"
	"gitee.com/conero/uymas/bin"
)

// uymas 子命令（测试）
func main() {
	plgCmd := bin.NewPluginCommand()
	plgCmd.Name = "demo"
	plgCmd.Descript = "插件式(热插拔)子命令 demo."

	plgCmd.RegisterFunc(func(arg *bin.Arg) {
		fmt.Printf("demo 子命令")
	})
	plgCmd.Run()
}
