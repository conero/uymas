package cli

import (
	"fmt"
	"gitee.com/conero/uymas/v2/logger/lgr"
	"gitee.com/conero/uymas/v2/rock"
	"log"
	"sort"
	"strings"
)

type RegisterMeta[T any] struct {
	Command  CommandOptional
	Runnable T
}

type Handler = func(ArgsParser) (bool, error)

// Register An experimental generic registry that supports different types for the underlying registration implementation
//
// todo Replace duplicate registration definitions, cli.Cli/evolve.Evolve
type Register[T any] struct {
	register      map[string]RegisterMeta[T]
	registerAlias map[string][]string
	indexTodo     T
	lostTodo      T
	helpTodo      T
	beforeHook    T
	endHook       T
	args          ArgsParser
	Config        Config
	Handler       Handler
	Call          func(T, ArgsParser)
}

// Command When registering a method you must specify commands to run more than one.
// We agreed that the second and subsequent commands should be aliases for the first command.
func (r *Register[T]) Command(t T, command string, optionals ...CommandOptional) Application[T] {
	r.CommandList(t, []string{command}, optionals...)
	return r
}

// CommandList When registering a method you must specify commands to run more than one.
// We agreed that the second and subsequent commands should be aliases for the first command.
func (r *Register[T]) CommandList(t T, commands []string, optionals ...CommandOptional) Application[T] {
	vNum := len(commands)
	if vNum == 0 {
		log.Fatal("Evolve.Command: when registering a method you must specify commands to run more than one. ")
		return r
	}

	if r.register == nil {
		r.register = map[string]RegisterMeta[T]{}
	}
	if r.registerAlias == nil {
		r.registerAlias = map[string][]string{}
	}

	optional := rock.Param(CommandOptional{}, optionals...)
	optional.Alias = commands[1:]
	optional.Keys = commands
	attr := RegisterMeta[T]{
		Command:  optional,
		Runnable: t,
	}
	// remember the command of alias.
	mainCmd := commands[0]

	// repetitive testing
	_, exist := r.register[mainCmd]
	if exist {
		panic(fmt.Sprintf("%s: please do not repeat the registration command", mainCmd))
	}

	r.register[mainCmd] = attr
	if vNum == 1 {
		return r
	}

	r.registerAlias[mainCmd] = commands[1:]
	return r
}

func (r *Register[T]) Index(t T) Application[T] {
	r.indexTodo = t
	return r
}

func (r *Register[T]) Lost(t T) Application[T] {
	r.lostTodo = t
	return r
}

func (r *Register[T]) Before(t T) Application[T] {
	r.beforeHook = t
	return r
}

func (r *Register[T]) End(t T) Application[T] {
	r.endHook = t
	return r
}

func (r *Register[T]) Help(t T) Application[T] {
	r.helpTodo = t
	return r
}

func (r *Register[T]) helpCmdName() (list []string, keys []string, maxLen int) {
	keys = rock.MapKeys(r.register)
	sort.Strings(keys)
	maxLen = 0
	for _, key := range keys {
		reg := r.register[key]
		opt := reg.Command.DataOption()
		docName := key
		getName := ""
		if opt != nil {
			getName = opt.GetName()
		}
		if opt != nil && getName != "" {
			docName += " [" + opt.GetName() + "]"
		}
		list = append(list, docName)
		vLen := len(docName)
		if vLen > maxLen {
			maxLen = vLen
		}
	}
	return
}

func (r *Register[T]) GetHelp(cmd string) (helpMsg string, exits bool) {
	if cmd == "" {
		var lines []string
		list, keys, maxLen := r.helpCmdName()
		for i, name := range keys {
			if name == "" {
				continue
			}
			docName := list[i]
			meta := r.register[name]
			reg := meta.Command
			cmdHelp := reg.Help
			if cmdHelp == "" {
				cmdHelp = "这是 " + name + " 命令（默认）"
			}
			if len(reg.Alias) > 0 {
				cmdHelp += "，别名支持 " + strings.Join(reg.Alias, ",")
			}
			line := fmt.Sprintf("%-"+(fmt.Sprintf("%d", maxLen+8))+"s", docName) + cmdHelp
			optionHelp := reg.OptionHelpMsg()
			if optionHelp != "" {
				line += "\n" + optionHelp
			}
			subHelpMsg := reg.SubCommandHelpMsg(2)
			if subHelpMsg != "" {
				line += "\n" + subHelpMsg
			}
			lines = append(lines, line)
		}
		helpMsg = strings.Join(lines, "\n")
		exits = true
		return
	}
	meta, hasCmd := r.register[cmd]
	reg := meta.Command
	if !hasCmd {
		for fName, fMeta := range r.register {
			fReg := fMeta.Command
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

func (r *Register[T]) GenerateHelpFn(arg ArgsParser) {
	cmdName := arg.HelpCmd()
	helpMsg, isFind := r.GetHelp(cmdName)
	if isFind {
		fmt.Println(helpMsg)
		fmt.Println()
		return
	}

	if cmdName != "" {
		lgr.Warn("命令 " + cmdName + " 不存在")
	}
}

func (r *Register[T]) Args() ArgsParser {
	return r.args
}

func (r *Register[T]) router(args ArgsParser) error {
	command := args.Command()
	cfg := r.Config
	helpCall := r.helpTodo

	isHelp := !cfg.DisableHelp && command == "" && args.Switch("help", "h", "?")
	isHelp = isHelp || (!cfg.DisableHelp && (command == "help" || command == "?"))
	if isHelp {
		r.Call(r.beforeHook, args)
		r.Call(helpCall, args)
		r.Call(r.endHook, args)
		return nil
	}

	if command == "" {
		r.Call(r.beforeHook, args)
		r.Call(r.indexTodo, args)
		r.Call(r.endHook, args)
		return nil
	}

	meta, isFind := r.register[command]
	if !isFind {
		for _, m := range r.register {
			com := m.Command
			if rock.InList(com.Keys, command) {
				meta = m
				isFind = true
				break
			}
		}
	}

	if isFind {
		r.Call(r.beforeHook, args)
		metaCmd := meta.Command
		if !r.Config.DisableVerify && !metaCmd.OffValid {
			invalidMsg := metaCmd.InvalidMsg(args)
			if invalidMsg != "" {
				lgr.Error(invalidMsg)
				return nil
			}
		}
		r.Call(meta.Runnable, args)
		r.Call(r.endHook, args)
		return nil
	}

	r.Call(r.lostTodo, args)
	return nil
}

func (r *Register[T]) Run(args ...string) error {
	if r.Config.ArgsConfig != nil {
		r.args = NewArgsWith(*r.Config.ArgsConfig, args...)
	} else {
		r.args = NewArgs(args...)
	}
	if r.registerAlias != nil {
		r.registerAlias = map[string][]string{}
	}

	return r.router(r.args)
}

func (r *Register[T]) RunArgs(args ArgsParser) error {
	r.args = args
	return r.Run()
}
