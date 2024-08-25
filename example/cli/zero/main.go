package main

import (
	"gitee.com/conero/uymas/v2/cli"
)

func main() {
	app := cli.NewCli()
	_ = app.Run()
}
