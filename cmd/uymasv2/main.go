package main

import (
	"gitee.com/conero/uymas/v2/cli"
	"log"
)

// v2 版本临时程序
//
// @todo 后期移除（稳定后）
func main() {
	app := cli.NewCli()
	err := app.Run()
	if err != nil {
		log.Fatalf("命令行执行错误，%v", err)
	}
}
