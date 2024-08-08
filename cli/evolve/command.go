package evolve

import "fmt"

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
	fmt.Printf("%s: welcome use the command\n", command)
}

// DefHelp help/reference command definition
func (c *Command) DefHelp() {}

// DefLost No command definition exists
func (c *Command) DefLost() {
	args := c.X.Args
	command := args.SubCommand()
	fmt.Println()
	fmt.Printf("%s: sub-command is not exist\n", command)
}
