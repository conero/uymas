// Package cli basic command line definition and processing tools.
//
// Simple command lines that do not apply to package reflect, only functional route definition is supported.
package cli

import (
	"fmt"
	"gitee.com/conero/uymas/v2"
	"gitee.com/conero/uymas/v2/logger/lgr"
	"gitee.com/conero/uymas/v2/rock"
	"gitee.com/conero/uymas/v2/str"
	"sort"
	"strings"
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
type Fn = func(ArgsParser)

func lostFn(arg ArgsParser) {
	fmt.Println()
	fmt.Printf("%s: We gotta lost, honey!\n    Uymas@%s/%s\n", arg.Command(), uymas.Version, uymas.Release)
	fmt.Println()
}

type registerAttr[T any] struct {
	CommandOptional
	runnable T
}

// Cli command line struct
type Cli struct {
	config       Config
	args         ArgsParser
	entryFn      Fn
	lostFn       Fn
	beforeFn     Fn
	endFn        Fn
	helpFn       Fn
	registerAttr map[string]registerAttr[Fn]
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
	optional.Keys = commands

	c.registerAttr[name] = registerAttr[Fn]{
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

func (c *Cli) RunArgs(args ArgsParser) error {
	c.args = args
	return c.router()
}

func (c *Cli) Help(t Fn) Application[Fn] {
	c.helpFn = t
	return c
}

func (c *Cli) getRegister(name string) (registerAttr[Fn], bool) {
	register, isMatch := c.registerAttr[name]
	if isMatch {
		return register, true
	}

	for _, reg := range c.registerAttr {
		if rock.InList(reg.Alias, name) {
			return reg, true
		}
	}

	return registerAttr[Fn]{}, false
}

func (c *Cli) generateHelpFn(arg ArgsParser) {
	cmdName := arg.HelpCmd()
	helpMsg, isFind := c.GetHelp(cmdName)
	if isFind {
		fmt.Println(helpMsg)
		fmt.Println()
		return
	}

	if cmdName != "" {
		lgr.Warn("命令 " + cmdName + " 不存在")
	}
}

func (c *Cli) generateEntryFn(arg ArgsParser) {
	config := c.config
	config.IndexDoc()
}

func (c *Cli) GetHelp(cmd string) (helpMsg string, exits bool) {
	if cmd == "" {
		var lines []string
		keys := rock.MapKeys(c.registerAttr)
		maxLen := str.QueueMaxLen(keys)
		sort.Strings(keys)
		for _, name := range keys {
			reg := c.registerAttr[name]
			cmdHelp := reg.Help
			if cmdHelp == "" {
				cmdHelp = "这是 " + name + " 命令（默认）"
			}
			if len(reg.Alias) > 0 {
				cmdHelp += "，别名支持 " + strings.Join(reg.Alias, ",")
			}
			line := fmt.Sprintf("%-"+(fmt.Sprintf("%d", maxLen+8))+"s", name) + cmdHelp
			optionHelp := reg.OptionHelpMsg()
			if optionHelp != "" {
				line += "\n" + optionHelp
			}
			lines = append(lines, line)
		}
		helpMsg = strings.Join(lines, "\n")
		exits = true
		return
	}
	reg, hasCmd := c.registerAttr[cmd]
	if !hasCmd {
		for fName, fReg := range c.registerAttr {
			if rock.InList(fReg.Alias, cmd) {
				reg = fReg
				cmd = fName
				hasCmd = true
				break
			}
		}
	}
	if !hasCmd {
		return
	}

	helpMsg = "命令 " + cmd + "，帮助信息如下：\n\n" + reg.Help
	if helpMsg == "" {
		helpMsg = "这是命令"
	}
	if len(reg.Alias) > 0 {
		helpMsg += "，别名 " + strings.Join(reg.Alias, ",")
	}
	optionHelp := reg.OptionHelpMsg()
	if optionHelp != "" {
		helpMsg += "\n" + optionHelp
	}
	commandHelp := reg.SubCommandHelpMsg()
	if commandHelp != "" {
		helpMsg += "\n\n子级命令菜单如下\n" + commandHelp
	}
	exits = true
	return
}

func (c *Cli) router() error {
	args := c.args
	command := args.Command()
	helpCall := c.helpFn
	if helpCall == nil {
		helpCall = c.generateHelpFn
	}
	cfg := c.config
	if !cfg.DisableHelp && command == "" && args.Switch("help", "h", "?") {
		helpCall(args)
	} else if !cfg.DisableHelp && (command == "help" || command == "?") {
		helpCall(args)
	} else if command != "" {
		reg, isFind := c.getRegister(command)
		if isFind {
			args.SetOptional(&reg.CommandOptional)
			if c.beforeFn != nil {
				c.beforeFn(args)
			}
			if !cfg.DisableHelp {
				invalidMsg := reg.InvalidMsg(args)
				if invalidMsg != "" {
					lgr.Error(invalidMsg)
					return nil
				}
			}
			reg.runnable(args)
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
	app := &Cli{
		config:       rock.Param(DefaultConfig, cfgs...),
		lostFn:       lostFn,
		registerAttr: map[string]registerAttr[Fn]{},
	}
	app.entryFn = app.generateEntryFn
	return app
}
