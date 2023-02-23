package bin

import (
	"bufio"
	"fmt"
	"gitee.com/conero/uymas/bin/parser"
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

	fmt.Print(tip)

	for input.Scan() {
		text := input.Text()
		text = strings.TrimSpace(text)

		isBreak := false
		switch text {
		case "exit":
			os.Exit(0)
		default:
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
