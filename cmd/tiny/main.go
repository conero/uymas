package main

import (
	"fmt"
	"gitee.com/conero/uymas/bin"
)

func main() {
	cli := bin.NewCLI()
	cli.RegisterAny(func() {
		fmt.Println("Uymas for tinygo")
		fmt.Println()
		fmt.Println("Just Try, use tinygo >= 0.31.0")
	})

	// help
	cli.RegisterFunc(func(arg *bin.Arg) {
		fmt.Println("tinygo    https://github.com/tinygo-org/tinygo")
		fmt.Println()
		fmt.Println("It's help command.")
	}, "help", "?")
	cli.Run()
}
