package main

import (
	"fmt"
	"github.com/conero/uymas/bin"
)

/**
 * @DATE        2019/5/2
 * @NAME        Joshua Conero
 * @DESCRIPIT   描述 descript
 **/

func main() {
	// 项目注册
	bin.RegisterFunc("name", func() {
		fmt.Println(" conero/uymas/bin example with Base.")
	})

	// 未知命令
	bin.UnfindFunc(func(cmd string) {
		fmt.Println(cmd + "unfind（functional）")
	})

	// 空函数
	bin.EmptyFunc(func() {
		fmt.Println("empty（functional）")
	})

	bin.Run()
}
