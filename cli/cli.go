// Package cli basic command line definition and processing tools.
//
// Simple command lines that do not apply to package reflect, only functional route definition is supported.
package cli

import (
	"fmt"
	"gitee.com/conero/uymas/v2"
)

// Application the command line program routes or parses the interface
type Application[T any] interface {
	// Command add command line
	Command(t T, commands ...string) Application[T]

	// Index command line entry method
	Index(t T) Application[T]

	// Lost command line arguments cannot be routed to
	Lost(t T) Application[T]

	// Run execute the command parser
	Run(args ...string) error
}

// Fn command line registration function
type Fn = func(...ArgsParser)

func entryFn(...ArgsParser) {
	fmt.Println()
	fmt.Println("-------------- Uymas -----------------")
	fmt.Println("Welcome to our world")
	fmt.Printf(":)- %s/%s\n", uymas.Version, uymas.Release)
	fmt.Println()
}
func lostFn(args ...ArgsParser) {
	arg := args[0]
	fmt.Println()
	fmt.Printf("%s: We gotta lost, honey!\n    Uymas@%s/%s\n", arg.Command(), uymas.Version, uymas.Release)
	fmt.Println()
}

// Cli command line struct
type Cli struct {
	config     Config
	args       ArgsParser
	entryFn    Fn
	lostFn     Fn
	registerFn map[string]Fn
}

func (c *Cli) Command(t Fn, commands ...string) Application[Fn] {
	for _, cmd := range commands {
		c.registerFn[cmd] = t
	}
	return c
}

func (c *Cli) Index(t Fn) Application[Fn] {
	c.entryFn = t
	return c
}

func (c *Cli) Lost(t Fn) Application[Fn] {
	c.lostFn = t
	return c
}

// Run execute command line as a program entry
func (c *Cli) Run(args ...string) error {
	var arg ArgsParser
	if c.config.ArgsConfig != nil {
		arg = NewArgsWith(*c.config.ArgsConfig, args...)
	} else {
		arg = NewArgs(args...)
	}
	c.args = arg
	return c.router()
}

func (c *Cli) router() error {
	args := c.args
	command := args.Command()
	if command != "" {
		fn, match := c.registerFn[args.Command()]
		if match {
			fn(args)
		} else {
			c.lostFn(args)
		}
	} else {
		c.entryFn(args)
	}

	return nil
}

// NewCli the command line program is instantiated and the driver is as light as possible
func NewCli(cfgs ...Config) *Cli {
	config := DefaultConfig
	if len(cfgs) > 0 {
		config = cfgs[0]
	}
	app := &Cli{
		config:     config,
		entryFn:    entryFn,
		lostFn:     lostFn,
		registerFn: map[string]Fn{},
	}
	return app
}
