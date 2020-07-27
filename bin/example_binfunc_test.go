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
	cli := NewCLI()

	// 项目注册
	cli.RegisterFunc(func(cc *CliCmd) {
		fmt.Println(" conero/uymas/bin example with Base.")
	}, "name")

	// 未知命令
	cli.RegisterUnfind(func(cmd string, cc *CliCmd) {
		fmt.Println(cmd + "unfind（functional）")
	})

	cli.RegisterEmpty(func(cc *CliCmd) {
		fmt.Println("empty（functional）")
	})

	cli.Run()
}
