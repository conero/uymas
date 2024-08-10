package evolve

import (
	"fmt"
)

const (
	CmdMtdInit  = "Init"
	CmdMtdIndex = "DefIndex"
	CmdMtdHelp  = "DefHelp"
	CmdMtdLost  = "DefLost"
	CmdFidX     = "X"
)

// Command line basic (minimum) structure standard
type Command struct {
	X *Param
}

// Init command initialization
func (c *Command) Init() {}

// DefIndex entry definition
func (c *Command) DefIndex() {
	args := c.X.Args
	command := args.Command()
	fmt.Println()
	fmt.Printf("%s: help information for command  here\n", command)
}

// DefHelp help/reference command definition
//
// support: `$ help [command]` or `$ -h [command]`
func (c *Command) DefHelp(cmds ...string) {
	command := c.X.Args.Command()
	cmdName := c.X.Args.HelpCmd()
	if cmdName != "" {
		command = "<" + command + " " + cmdName + ">"
	}
	fmt.Printf("Default Help: we should add the help information for command %s here, honey!\n\n", command)
}

// DefLost No command definition exists
func (c *Command) DefLost() {
	args := c.X.Args
	command := args.SubCommand()
	fmt.Println()
	fmt.Printf("%s: sub-command is not exist\n", command)
}
