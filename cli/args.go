package cli

import "os"

// ArgsParser command line parameter parsing interface
type ArgsParser interface {
	// Values The original data type of the command
	Values() map[string][]string
	// Get command line data by key value
	Get(key ...string) string
	// GetDef get command line data by key value and specify default values
	GetDef(def string, key ...string) string
	// Switch determines whether the option specified by the key value exists
	Switch(key ...string) bool
	// Command get the command of the command line program
	Command() string
}

// Args command line program parameters
type Args struct {
	raw     []string
	command string
	ArgsParser
}

// parse data by args
func (c *Args) parse() {
	for i, arg := range c.raw {
		if i == 0 {
			c.command = arg
			continue
		}
	}
}

func (c *Args) Values() map[string][]string {
	return nil
}

func (c *Args) Get(key ...string) string {
	return ""
}

func (c *Args) GetDef(def string, key ...string) string {
	return ""
}

func (c *Args) Switch(key ...string) bool {
	return false
}

func (c *Args) Command() string {
	return ""
}

func NewArgs(args ...string) ArgsParser {
	if len(args) == 0 {
		args = os.Args[1:]
	}
	arg := &Args{
		raw: args,
	}
	arg.parse()
	return arg
}
