package main

import (
	"fmt"
	"gitee.com/conero/uymas/v2/cli/evolve"
)

type defaultApp struct {
	evolve.Command
}

func (c *defaultApp) DefIndex() {
	fmt.Println("Hello World, Struct.")
}

func (c *defaultApp) Demo() {
	fmt.Println("Demo command")
}
