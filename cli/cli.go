// Package cli basic command line definition and processing tools.
//
// Simple command lines that do not apply to package reflect, only functional route definition is supported.
package cli

import (
	"fmt"
	"gitee.com/conero/uymas/v2"
	"gitee.com/conero/uymas/v2/rock"
)

// Application the command line program routes or parses the interface
type Application[T any] interface {
	// Command add command line
	Command(t T, command string, optionals ...CommandOptional) Application[T]
	CommandList(t T, commands []string, optionals ...CommandOptional) Application[T]

	// Index command line entry method
	Index(t T) Application[T]

	// Lost command line arguments cannot be routed to
	Lost(t T) Application[T]

	// Before hook for command line run command to do something before
	Before(t T) Application[T]

	// End hook for command line run command to do something in end
	End(t T) Application[T]

	Help(t T) Application[T]

	// Run execute the command parser
	Run(args ...string) error
}

// Fn command line registration function
type Fn = func(ArgsParser)

func entryFn(ArgsParser) {
	fmt.Println()
	fmt.Println("-------------- Uymas -----------------")
	fmt.Println("Welcome to our world")
	fmt.Printf(":)- %s/%s\n", uymas.Version, uymas.Release)
	fmt.Println()
}
func lostFn(arg ArgsParser) {
	fmt.Println()
	fmt.Printf("%s: We gotta lost, honey!\n    Uymas@%s/%s\n", arg.Command(), uymas.Version, uymas.Release)
	fmt.Println()
}

func helpFn(arg ArgsParser) {
	command := arg.Command()
	cmdName := arg.HelpCmd()
	if cmdName != "" {
		command = "<" + command + " " + cmdName + ">"
	}
	if command != "" {
		command += " "
	}
	fmt.Printf("Default Help: we should add the help information for command %shere, honey!\n\n", command)
}

type registerMap[T any] struct {
	CommandOptional
	runnable T
}

// Cli command line struct
type Cli struct {
	config      Config
	args        ArgsParser
	entryFn     Fn
	lostFn      Fn
	beforeFn    Fn
	endFn       Fn
	helpFn      Fn
	registerMap map[string]registerMap[Fn]
}

func (c *Cli) Command(t Fn, cmd string, optionals ...CommandOptional) Application[Fn] {
	c.CommandList(t, []string{cmd}, optionals...)
	return c
}

func (c *Cli) CommandList(t Fn, commands []string, optionals ...CommandOptional) Application[Fn] {
	if len(commands) == 0 {
		panic("CommandList: a valid parameter commands must be specified, but you provided an empty array")
	}
	name := commands[0]
	optional := rock.Param(CommandOptional{}, optionals...)
	optional.Alias = commands[1:]

	c.registerMap[name] = registerMap[Fn]{
		CommandOptional: optional,
		runnable:        t,
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

func (c *Cli) Before(t Fn) Application[Fn] {
	c.beforeFn = t
	return c
}

func (c *Cli) End(t Fn) Application[Fn] {
	c.endFn = t
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

func (c *Cli) Help(t Fn) Application[Fn] {
	c.helpFn = t
	return c
}

func (c *Cli) getCall(name string) Fn {
	register, isMatch := c.registerMap[name]
	if isMatch {
		return register.runnable
	}

	for _, reg := range c.registerMap {
		if rock.InList(reg.Alias, name) {
			return reg.runnable
		}
	}

	return nil
}

func (c *Cli) router() error {
	args := c.args
	command := args.Command()
	helpCall := c.helpFn
	if helpCall == nil {
		helpCall = helpFn
	}
	cfg := c.config
	if !cfg.DisableHelp && command == "" && args.Switch("help", "h", "?") {
		helpCall(args)
	} else if !cfg.DisableHelp && (command == "help" || command == "?") {
		helpCall(args)
	} else if command != "" {

		fn := c.getCall(command)
		if fn != nil {
			if c.beforeFn != nil {
				c.beforeFn(args)
			}
			fn(args)
			if c.endFn != nil {
				c.endFn(args)
			}
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
		config:      config,
		entryFn:     entryFn,
		lostFn:      lostFn,
		registerMap: map[string]registerMap[Fn]{},
	}
	return app
}
