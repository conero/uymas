package main

import (
	"bufio"
	"fmt"
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
	cli.Run()
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