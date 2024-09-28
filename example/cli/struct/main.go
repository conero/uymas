package main

import (
	"fmt"
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/cli/gen"
	"gitee.com/conero/uymas/v2/logger/lgr"
)

func main() {
	app := gen.AsCommand(new(defaultApp))
	// User-defined external registration commands
	app.Command(func(parser cli.ArgsParser) {
		fmt.Println("用户自定义命令，JC")
	}, "cust", cli.Help("用户自定义命令"))
	lgr.ErrorIf(app.Run())
}
