package bin

import (
	"bufio"
	"fmt"
	"gitee.com/conero/uymas/bin/parser"
	"os"
	"strings"
	"testing"
)

func ExampleNewCLI_func() {
	cli := NewCLI()
	cli.RegisterAny(func() {
		fmt.Println("Hello world, cli.")
	})
	cli.Run()

	// Output:
	// Hello world, cli.
}

type defaultAppHelloWorld struct {
	CliApp
}

func (w *defaultAppHelloWorld) DefaultIndex() {
	fmt.Println("Hello world, struct cli.")
}

func ExampleNewCLI_struct() {
	cli := NewCLI()
	// add the new struct instance.
	// type defaultAppHelloWorld struct {
	//	bin.CliApp
	// }
	//
	// func (w *defaultAppHelloWorld) DefaultIndex() {
	//	fmt.Println("Hello world, struct cli.")
	// }
	cli.RegisterAny(new(defaultAppHelloWorld))
	cli.Run()

	// Output:
	// Hello world, struct cli.
}

// to run Example func tpl
func TestNewCLI(t *testing.T) {
	//ExampleXXX
	//ExampleNewCLI_func()
	ExampleNewCLI_struct()
}

func TestNewCLI_Repl(t *testing.T) {
	ExampleNewCLI_repl()
}

func ExampleNewCLI_repl() {
	var input = bufio.NewScanner(os.Stdin)
	prefShow := func() {
		fmt.Print("$ uymas> ")
	}
	cli := NewCLI()
	cli.RegisterAny(new(defaultAppHelloWorld))
	prefShow()
	for input.Scan() {
		text := input.Text()
		text = strings.TrimSpace(text)
		switch text {
		case "exit":
			os.Exit(0)
		default:
			// to run struct command.
			var cmdsList = parser.NewParser(text)
			for _, cmdArgs := range cmdsList {
				cli.Run(cmdArgs...)
			}
		}
		fmt.Println()
		prefShow()
	}
}
