package main

import (
	"fmt"
	"gitee.com/conero/uymas/v2/cli/evolve"
)

type test struct {
	evolve.Command
}

func (c *test) Demo() {
	fmt.Println("test demo, ha!")
}

func main() {
	evl := evolve.NewEvolve()
	evl.Command(func() {
		fmt.Println("Evolution For Index.")
	}, "index")
	evl.Command(new(test), "test", "t")
	evl.Run()
}
