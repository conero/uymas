package main

import (
	"bufio"
	"fmt"
	"github.com/conero/uymas"
	"github.com/conero/uymas/bin"
	"github.com/conero/uymas/bin/butil"
	"os"
	"strings"
)

func main() {
	//oldBin()
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

	//empty data.
	//cli.RegisterFunc(func(cc *bin.CliCmd) {
	//})

	cli.RegisterApp(&TestCmd{}, "test")
	cli.Run("test", "verify")
	//cli.Run("-xyz", "--name", "'Joshua Conero'", "--first=emma", "--list", "A", "B", "c", "table.name")
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
			tmpArgs := butil.StringToArgs(text)
			//fmt.Println(tmpArgs)
			cli.Run(tmpArgs...)
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