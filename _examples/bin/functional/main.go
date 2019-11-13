package main

// @DATE   2019/8/27
// @AUTHOR Joshua Conero<conero@163.com>
// @NAME   函数式命令测试

import (
	"fmt"
	"github.com/conero/uymas/bin"
)

func main() {
	// 测试
	//bin.InjectArgs("--part=8080")

	// 空命令
	bin.EmptyFunc(func(a *bin.App) {
		fmt.Println(a.Data)
		//fmt.Println(bin.Args())
	})

	// 运行命令行程序
	bin.Run()
}
