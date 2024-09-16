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
	// RunArgs run the command line program by setting Args
	RunArgs(args ArgsParser) error
}

// Fn command line registration function
type Fn = func(args ArgsParser)

func lostFn(arg ArgsParser) {
	fmt.Println()
	fmt.Printf("%s: We gotta lost, honey!\n    Uymas@%s/%s\n", arg.Command(), uymas.Version, uymas.Release)
	fmt.Println()
}

// Cli command line struct
type Cli struct {
	Register[Fn]
}

// NewCli the command line program is instantiated and the driver is as light as possible
func NewCli(cfgs ...Config) *Cli {
	app := &Cli{}
	app.Config = rock.Param(DefaultConfig, cfgs...)
	app.Call = func(fn Fn, parser ArgsParser) {
		if fn != nil {
			fn(parser)
			return
		}
	}
	app.lostTodo = lostFn
	app.indexTodo = func(parser ArgsParser) {
		app.Config.IndexDoc()
	}
	app.helpTodo = app.GenerateHelpFn
	return app
}
