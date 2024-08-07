package main

import (
	"fmt"
	"gitee.com/conero/uymas/v2/cli/evolve"
)

func main() {
	evl := evolve.NewEvolve()
	evl.Command(func() {
		fmt.Println("Evolution For Index.")
	}, "index")
	evl.Run()
}
