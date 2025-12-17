package main

import (
	"fmt"

	"gitee.com/conero/uymas/v2/cli/chest"
)

func main() {
	fmt.Println("----- InputOption -----")
	name := chest.InputOption("请输入姓名：", "")
	fmt.Printf("您输出姓名：%s\n", name)

	fmt.Println()
	fmt.Println("----- InputRequireDef -----")
	name = chest.InputRequireDef("请输入姓名：", "Jo")
	fmt.Printf("您输出姓名：%s\n", name)
}
