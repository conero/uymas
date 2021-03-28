package main

import (
	"fmt"
	"github.com/conero/uymas/bin"
	"strings"
)

//App 版本测试
type App struct {
	bin.CliApp
}

func (c *App) Construct() {
}

//测试命令
func (c *App) Test() {
	cc := c.Cc
	fmt.Println(" 命令行测试")
	fmt.Printf("  SubCommand: %v \r\n", cc.SubCommand)
	fmt.Printf("  Option: %v \r\n", cc.Setting)
	fmt.Printf("  DataRaw: %v \r\n", cc.DataRaw)
	fmt.Printf("  Data: %#v \r\n", cc.Data)
	fmt.Printf("  Input: %#v \r\n", strings.Join(cc.Raw, " "))
	fmt.Printf("  Current next: %#v \r\n", cc.Next())
	fmt.Printf("  Is CmdApp : %#v \r\n", cc.CmdType() == int(bin.CmdApp))
	fmt.Println()
}
