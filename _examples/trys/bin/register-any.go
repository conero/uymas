package main

import (
	"fmt"
	"github.com/conero/uymas/bin"
)

func main() {
	cli := bin.NewCLI()
	cli.RegisterAny(&RegisterAny{})
	cli.Run()
}


type RegisterAny struct {
	Cc bin.CliCmd
}

func (c *RegisterAny) Construct()  {
	fmt.Println(" Any-init ")
}

func (c *RegisterAny) Before()  {
	cc := c.Cc
	fmt.Println(" @Before ")
	fmt.Printf(" Data -> %v\r\n", cc.DataRaw)
}