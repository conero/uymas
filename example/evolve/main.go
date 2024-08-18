package main

import (
	"fmt"
	"gitee.com/conero/uymas/v2/cli/evolve"
	"gitee.com/conero/uymas/v2/util/fs"
)

type test struct {
	evolve.Command
}

func (c *test) Demo() {
	fmt.Println("test demo, ha!")
	fmt.Println()
	fmt.Println("rootPath: " + fs.RootDir())
	fmt.Println("rootApp: " + fs.AppName())
}

func main() {
	evl := evolve.NewEvolve()
	evl.Command(func() {
		fmt.Println("Evolution For Index.")
	}, "index")
	evl.Command(new(test), "test", "t")
	//evl.Run("test", "demo")
	evl.Run()
}
