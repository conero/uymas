package evolve

import (
	"fmt"
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/logger/lgr"
)

const (
	CmdMtdInit  = "Init"
	CmdMtdIndex = "DefIndex"
	CmdMtdHelp  = "DefHelp"
	CmdMtdLost  = "DefLost"
	CmdFidArgs  = "Args"
)

// Command line basic (minimum) structure standard
type Command struct {
	Args cli.ArgsParser
}

// Init command initialization
func (c *Command) Init() {}

// DefIndex entry definition
func (c *Command) DefIndex() {
	args := c.Args
	command := args.Command()
	fmt.Println()
	fmt.Printf("%s: help information for command here\n", command)
}

// DefHelp help/reference command definition
//
// support: `$ help [command]` or `$ -h [command]`
func (c *Command) DefHelp() {
	command := c.Args.Command()
	cmdName := c.Args.HelpCmd()
	if cmdName != "" {
		command = "<" + command + " " + cmdName + ">"
	}
	if command != "" {
		command += " "
	}
	fmt.Printf("Default Help: we should add the help information for command %shere, honey!\n\n", command)
}

// DefLost No command definition exists
func (c *Command) DefLost() {
	args := c.Args
	command := args.SubCommand()
	lgr.Error("%s: sub-command is not exist", command)
	fmt.Println()
}
