package main

import (
	"fmt"
	"github.com/conero/uymas/bin"
	"github.com/conero/uymas/bin/butil"
)

//the example about clear the clis
func main() {
	base()
}



func base(){
	cli := bin.NewCLI()

	// clear the cli app
	cli.RegisterFunc(func(cc *bin.CliCmd) {
		er := butil.Clear()
		if er != nil{
			fmt.Println(er)
		}
	}, "clear")

	cli.RegisterFunc(func(cc *bin.CliCmd) {
		fmt.Println("the example about call the system, use clear")
	})

	cli.Run()
}
