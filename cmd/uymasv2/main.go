package main

import (
	"fmt"
	"gitee.com/conero/uymas/v2/cli"
	"log"
)

// v2 版本临时程序
//
// @todo 后期移除（稳定后）
func main() {
	app := cli.NewCli()
	//app.Index(func(parser ...cli.ArgsParser) {
	//	fmt.Println("欢饮您使用 Uymas v2")
	//})
	app.Command(func(parser ...cli.ArgsParser) {
		arg := parser[0]
		fmt.Println()
		fmt.Println("参数解析，数据如下")
		fmt.Println()
		fmt.Printf("value: %v\n", arg.Values())
		fmt.Printf("option: %v\n", arg.Option())
		fmt.Printf("CommandList: %v\n", arg.CommandList())
		fmt.Println()
	}, "test")
	err := app.Run()
	if err != nil {
		log.Fatalf("命令行执行错误，%v", err)
	}
}
