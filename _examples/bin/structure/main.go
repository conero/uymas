package main

import (
	"fmt"
	"github.com/conero/uymas/bin"
)

func main() {
	bin.Register("test", new(TypeCommand))
	bin.Run()
}


type TypeCommand struct {
	bin.Command
}

func (tc *TypeCommand) Init()  {
	tc.Command.Init()
}


func (tc *TypeCommand) Run()  {
	fmt.Println(tc.App.DataRaw)
}