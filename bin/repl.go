package bin

import (
	"bufio"
	"fmt"
	"gitee.com/conero/uymas/bin/parser"
	"gitee.com/conero/uymas/util"
	"os"
	"strings"
)

const (
	ReplContinue = iota
	ReplBreak
	ReplExit
)

type Repl struct {
	Name        string
	Title       string
	HandlerFunc func(string) int
	BeforeExit  func()   // 退出前回调
	Exit        []string // 退出列表，默认为 exit
}

// Run to start repl command
func (c Repl) Run(cli *CLI) {
	var input = bufio.NewScanner(os.Stdin)
	name := c.Name
	if name != "" {
		name = " " + name
	}
	tip := fmt.Sprintf("$%v> ", name)
	if c.Title != "" {
		fmt.Println(c.Title)
	}

	// 退出命令
	toExit := func() {
		if c.BeforeExit != nil {
			c.BeforeExit()
		}
		os.Exit(0)
	}

	fmt.Print(tip)
	for input.Scan() {
		text := input.Text()
		text = strings.TrimSpace(text)

		isBreak := false
		switch text {
		default:
			if c.Exit == nil && text == "exit" {
				toExit()
			} else if util.ListIndex(c.Exit, text) > -1 {
				toExit()
			}

			if c.HandlerFunc != nil {
				switch c.HandlerFunc(text) {
				case ReplBreak:
					isBreak = true
				case ReplExit:
					os.Exit(0)
				case ReplContinue:
					continue
				}
			}
			if isBreak {
				break
			}
			var cmdsList = parser.NewParser(text)
			for _, cmdArgs := range cmdsList {
				cli.Run(cmdArgs...)
			}
		}
		if isBreak {
			break
		}

		fmt.Println()
		fmt.Println()
		fmt.Print(tip)
	}
}

// ReplCommand create a repl cli sub command
func ReplCommand(name string, cli *CLI) {
	repl := Repl{
		Name: name,
	}
	repl.Run(cli)
}
