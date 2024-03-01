package main

import (
	"fmt"
	"gitee.com/conero/uymas/v2/bin"
	"gitee.com/conero/uymas/v2/bin/butil"
)

// the example about clear the clis
func main() {
	base()
}

func base() {
	cli := bin.NewCLI()

	// clear the cli app
	cli.RegisterFunc(func(cc *bin.Arg) {
		er := butil.Clear()
		if er != nil {
			fmt.Println(er)
		}
	}, "clear")

	cli.RegisterFunc(func(cc *bin.Arg) {
		fmt.Println("the example about call the system, use clear")
	})

	cli.Run()
}
