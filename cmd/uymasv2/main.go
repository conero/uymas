package main

import (
	"flag"
	"fmt"
	"gitee.com/conero/uymas/v2/cli"
	"log"
)

// v2 版本临时程序
//
// @todo 后期移除（稳定后）
func main() {
	// 系统内容分命令行测试
	check := flag.Bool("check", false, "A boolean flag")
	flag.Parse()
	fmt.Printf("bool: %v\n", *check)
	fmt.Println()

	app := cli.NewCli()
	err := app.Run()
	if err != nil {
		log.Fatalf("命令行执行错误，%v", err)
	}
}
