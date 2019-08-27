package bin

import (
	"fmt"
)

/**
 * @DATE        2019/6/3
 * @NAME        Joshua Conero
 * @DESCRIPIT   描述 descript
 **/

func main() {
	// 项目注册
	RegisterFunc("name", func() {
		fmt.Println(" conero/uymas/bin example with Base.")
	})

	// 未知命令
	UnfindFunc(func(cmd string) {
		fmt.Println(cmd + "unfind（functional）")
	})

	// 空函数
	EmptyFunc(func() {
		fmt.Println("empty（functional）")
	})

	Run()
}