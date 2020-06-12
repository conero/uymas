package main

import (
	"bufio"
	"fmt"
	"github.com/conero/uymas"
	"github.com/conero/uymas/bin"
	"github.com/conero/uymas/bin/butil"
	"github.com/conero/uymas/bin/parser"
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
	//cli.RegisterFunc(func(cc *bin.CliCmd) {
	//})

	cli.RegisterUnfind(func(cmd string, cc *bin.CliCmd){
		fmt.Println("  通用命令解析: ")
		fmt.Printf("    command: %v\r\n", cmd)
		if cc.SubCommand != ""{
			fmt.Printf("    sub_command: %v\r\n",cc.SubCommand)
		}
		fmt.Printf("    data_raw: %v\r\n", cc.DataRaw)
		fmt.Printf("    data: %v\r\n", cc.Data)
		fmt.Printf("    setting: %v\r\n", cc.Setting)
	})

	cli.RegisterFunc(func(cc *bin.CliCmd) {
		cList := cli.GetCmdList()
		for _, c := range cList{
			//打印数据列表
			fmt.Printf("%v       %v\r\n", c, cli.GetDescribe(c))
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

//the old bin construct
func oldBin()  {
	bin.UnfindFunc(func(a *bin.App, cmd string) {
		fmt.Printf("    command: %v\r\n", cmd)
		if a.SubCommand != ""{
			fmt.Printf("    sub_command: %v\r\n", a.SubCommand)
		}
		fmt.Printf("    data_raw: %v\r\n", a.DataRaw)
		fmt.Printf("    data: %v\r\n", a.Data)
		fmt.Printf("    setting: %v\r\n", a.Setting)
	})
	lingering()
}

//驻留式命令行程序
func lingering()  {
	var input = bufio.NewScanner(os.Stdin)
	fmt.Println("驻留式命令行程序")
	fmt.Print("$ uymas>")

	for input.Scan() {
		text := input.Text()
		text = strings.TrimSpace(text)

		switch text {
		default:
			tmpArgs := butil.StringToArgs(text)
			fmt.Println(tmpArgs)
			bin.InjectArgs(tmpArgs...)
			bin.Run()
		}

		fmt.Println()
		fmt.Println()
		fmt.Print("$ uymas>")
	}
}


// the command of `test`.
type TestCmd struct {
	bin.CliApp
}

// need the construct
func (tc *TestCmd) Construct()  {
	//tc.DoRouter()
}

//cmd `test verify`.
func (tc *TestCmd) Verify()  {
	cc := tc.Cc
	fmt.Println(cc.Command)
	fmt.Println(cc.SubCommand)
}