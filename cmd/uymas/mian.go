package main

import (
	"bufio"
	"fmt"
	"github.com/conero/uymas"
	"github.com/conero/uymas/bin"
	"github.com/conero/uymas/bin/parser"
	"github.com/conero/uymas/culture/pinyin"
	"os"
	"strings"
)

var (
	cli         *bin.CLI
	pinyinCache *pinyin.Pinyin = nil
)

//the cli app tools
func application() {
	cli = bin.NewCLI()

	//pinyin
	cli.RegisterFunc(func(cc *bin.CliCmd) {
		words := cc.SubCommand
		if words != "" {
			pinyinCache = getPinyin()
			fmt.Println(pinyinCache.GetPyTone(words))
		} else {
			fmt.Println("请输入: $ pinyin <汉字>")
		}
	}, "pinyin")

	//empty
	cli.RegisterEmpty(func(cc *bin.CliCmd) {
		fmt.Printf(" wecolme use the <%v>. \r\n", uymas.Name)
		fmt.Printf(" %v/%v\r\n", uymas.Version, uymas.Release)
		fmt.Printf(" Power by %v.\r\n", uymas.Author)
	})

	//uls,uymas ls
	cli.RegisterFunc(func(cc *bin.CliCmd) {
		fmt.Println("  " + strings.Join(cli.GetCmdList(), "\r\n  "))
	}, "uls", "uymas-ls")

	//REPL
	cli.RegisterFunc(func(cc *bin.CliCmd) {
		var input = bufio.NewScanner(os.Stdin)
		fmt.Println("驻留式命令行程序")
		fmt.Print("$ uymas>")

		for input.Scan() {
			text := input.Text()
			text = strings.TrimSpace(text)

			switch text {
			case "exit":
				os.Exit(0)
			default:
				var cmdsList = parser.NewParser(text)
				for _, cmdArgs := range cmdsList {
					cli.Run(cmdArgs...)
				}
			}

			fmt.Println()
			fmt.Println()
			fmt.Print("$ uymas>")
		}
	}, "repl")

	cli.Run()
}

func getPinyin() *pinyin.Pinyin {
	if pinyinCache == nil {
		pinyinCache = pinyin.NewPinyin("./resource/culture/pinyin.txt")
	}
	return pinyinCache
}

//the uymas cmd message
func main() {
	application()
}
