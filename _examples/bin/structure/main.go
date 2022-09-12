package main

import (
	"fmt"
	"gitee.com/conero/uymas/bin"
)

func main() {
	cli := bin.NewCLI()
	cli.RegisterApp(new(TypeCommand), "test")
	cli.RegisterApp(new(TypeCommand), "type")
	cli.Run()
}

type TypeCommand struct {
	bin.CliApp
}

func (tc *TypeCommand) Construct() {
}

func (tc *TypeCommand) Debug() {
	context := tc.Cc.Context()
	vList := context.GetCmdList()
	fmt.Printf("GetCmdList: %#v\n", vList)
}
