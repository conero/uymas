package repl

import (
	"bufio"
	"fmt"
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/logger/lgr"
	"gitee.com/conero/uymas/v2/rock"
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
func (c *Repl) Run(cliApp *cli.Cli) {
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
			} else if rock.ListIndex(c.Exit, text) > -1 {
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
			var cmdsList = strings.Split(text, " ")
			err := cliApp.Run(cmdsList...)
			if err != nil {
				lgr.Error("命令执行错误，%v", err)
			}
		}
		if isBreak {
			break
		}

		fmt.Println()
		fmt.Print(tip)
	}
}

// Command create a repl cli sub command
func Command(name string, cliApp *cli.Cli) {
	repl := &Repl{
		Name: name,
	}
	repl.Run(cliApp)
}
