package main

import (
	"fmt"
	"gitee.com/conero/uymas/bin"
	"strings"
)

func main() {
	cli := bin.NewCLI()
	cli.RegisterApp(&App{}, "app")
	cli.RegisterAny(&RegisterAny{})
	cli.Run()
}

type RegisterAny struct {
	Cc bin.CliCmd
}

func (c *RegisterAny) Construct() {
	fmt.Println(" Any-init ")
}

func (c *RegisterAny) Before() {
	cc := c.Cc
	fmt.Println(" @Before ")
	fmt.Printf(" Data -> %v\n", cc.DataRaw)
}

// DefaultUnmatched unmatched
func (c *RegisterAny) DefaultUnmatched() {
	fmt.Println(" @DefaultUnmatched --> Oh, the method is not match. ")
}

func (c *RegisterAny) DefaultHelp() {
	fmt.Println(" @DefaultHelp --> isHelp Command. ")
	fmt.Println(" command like: before.")
}

func (c *RegisterAny) DefaultIndex() {
	fmt.Println(" @DefaultIndex --> Default index. ")
}

func (c *RegisterAny) Test() {
	cc := c.Cc
	fmt.Println(" 命令行测试")
	fmt.Printf("  SubCommand: %v \n", cc.SubCommand)
	fmt.Printf("  Option: %v \n", cc.Setting)
	fmt.Printf("  DataRaw: %v \n", cc.DataRaw)
	fmt.Printf("  Data: %#v \n", cc.Data)
	fmt.Printf("  Input: %#v \n", strings.Join(cc.Raw, " "))
	fmt.Println()
}

type App struct {
	bin.CliApp
}

func (c *App) Construct() {
	fmt.Println(" @App.Construct")
}

func (c *App) DefaultIndex() {
	fmt.Println(" @App.DefaultIndex")
}

func (c *App) DefaultUnmatched() {
	fmt.Println(" @App.DefaultUnmatched")
}

func (c *App) DefaultHelp() {
	fmt.Println(" @App.DefaultHelp")
}

func (c *App) Test() {
	cc := c.Cc
	fmt.Println(" 命令行测试")
	fmt.Printf("  SubCommand: %v \n", cc.SubCommand)
	fmt.Printf("  Option: %v \n", cc.Setting)
	fmt.Printf("  DataRaw: %v \n", cc.DataRaw)
	fmt.Printf("  Data: %#v \n", cc.Data)
	fmt.Printf("  Input: %#v \n", strings.Join(cc.Raw, " "))
	fmt.Println()
}
