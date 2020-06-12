package main

import "github.com/conero/uymas/bin"

var (
	cli *bin.CLI
)

//the cli app tools
func application() {
	cli = bin.NewCLI()
	cli.Run()
}

//the uymas cmd message
func main() {
	application()
}
