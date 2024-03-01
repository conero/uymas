package main

import (
	"fmt"
	"gitee.com/conero/uymas/v2/bin"
	"gitee.com/conero/uymas/v2/bin/butil"
	"strings"
)

// App 版本测试
type App struct {
	bin.CliApp
}

func (c *App) Construct() {
	fmt.Println("App 初始化：Construct")
}

// Test 测试命令
func (c *App) Test() {
	cc := c.Cc
	fmt.Println(" 命令行测试")
	fmt.Printf("  SubCommand: %v \n", cc.SubCommand)
	fmt.Printf("  Option: %v \n", cc.Setting)
	fmt.Printf("  DataRaw: %v \n", cc.DataRaw)
	fmt.Printf("  Data: %#v \n", cc.Data)
	fmt.Printf("  Input: %#v \n", strings.Join(cc.Raw, " "))
	fmt.Printf("  Current next: %#v \n", cc.Next())
	fmt.Printf("  Is CmdApp : %#v \n", cc.CmdType() == int(bin.CmdApp))
	fmt.Printf("  Basedir : %v \n", butil.Basedir())
	fmt.Println()
}
