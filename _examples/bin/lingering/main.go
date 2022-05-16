package main

import (
	"bufio"
	"fmt"
	"gitee.com/conero/uymas"
	"gitee.com/conero/uymas/bin"
	"gitee.com/conero/uymas/bin/butil"
	"gitee.com/conero/uymas/bin/parser"
	"os"
	"strings"
)

func main() {
	//oldBin()
	//新版本，命令行程序实现
	newBin()
}

func newBin()  {
	cli := bin.NewCLI()

	cli.RegisterFunc(func(cmd *bin.CliCmd) {
		fmt.Println("this is help command.")
	}, "help", "?")

	cli.RegisterFunc(func(cmd *bin.CliCmd) {
		fmt.Println(uymas.Version + "/" + uymas.Release)
	}, "version")

	//empty data.
	cli.RegisterFunc(func(cc *bin.CliCmd) {
	})

	//the empty data
	cli.RegisterEmpty(func(cmd *bin.CliCmd) {
		fmt.Println("welcome the new BIN.")
		newLingering(cmd, cli)
		//fmt.Println(cmd.Raw)
		//fmt.Println(cmd.Setting)
		//fmt.Println(cmd.DataRaw)
		//fmt.Println(cli.GetCmdList())
	})

	//clear the system
	cli.RegisterFunc(func(cmd *bin.CliCmd) {
		er := butil.Clear()
		if er != nil{
			fmt.Printf(" ERROR: %v", er)
		}
	}, "clear")

	//empty data.

	cli.RegisterAny(func(cmd string, cc *bin.CliCmd){
		fmt.Println("  通用命令解析: ")
		fmt.Printf("    command: %v\n", cmd)
		if cc.SubCommand != ""{
			fmt.Printf("    sub_command: %v\n",cc.SubCommand)
		}
		fmt.Printf("    data_raw: %v\n", cc.DataRaw)
		fmt.Printf("    data: %v\n", cc.Data)
		fmt.Printf("    setting: %v\n", cc.Setting)
	})

	cli.RegisterFunc(func(cc *bin.CliCmd) {
		cList := cli.GetCmdList()
		for _, c := range cList{
			//打印数据列表
			fmt.Printf("%v       %v\n", c, cli.GetDescribe(c))
		}
	}, "list", "ls")


	cli.RegisterCommand(bin.Cmd{
		Command:  "author",
		Alias:    nil,
		Describe: "print the person who build the project.",
		Handler: func(cc *bin.CliCmd) {
			fmt.Println("I am Joshua Conero.")
		},
		Options:  nil,
	})

	cli.RegisterApp(&TestCmd{}, "test")
	//cli.Run("test", "verify")
	cli.Run("-xyz", "--name", "'Joshua Conero'", "--first=emma", "--list", "A", "B", "c", "table.name")
}

func newLingering(cc *bin.CliCmd, cli *bin.CLI)  {
	var input = bufio.NewScanner(os.Stdin)
	fmt.Println("驻留式命令行程序")
	fmt.Print("$ uymas>")

	for input.Scan() {
		text := input.Text()
		text = strings.TrimSpace(text)

		switch text {
		default:
			var cmdsList = parser.NewParser(text)
			for _, cmdArgs := range cmdsList{
				cli.Run(cmdArgs...)
			}
		}

		fmt.Println()
		fmt.Println()
		fmt.Print("$ uymas>")
	}
}

// TestCmd the command of `test`.
type TestCmd struct {
	bin.CliApp
}

// Construct need the construct
func (tc *TestCmd) Construct()  {
	//tc.DoRouter()
}

// Verify cmd `test verify`.
func (tc *TestCmd) Verify()  {
	cc := tc.Cc
	fmt.Println(cc.Command)
	fmt.Println(cc.SubCommand)
}