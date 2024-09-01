// Package evolve version Command line, which supports more registration types than cli. Adopting reflection.
package evolve

import (
	"fmt"
	"gitee.com/conero/uymas/v2"
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/logger/lgr"
	"gitee.com/conero/uymas/v2/rock"
	"gitee.com/conero/uymas/v2/str"
	"log"
	"reflect"
	"sort"
	"strings"
)

type Evolve[T any] struct {
	config        cli.Config
	indexTodo     T
	lostTodo      T
	helpTodo      T
	beforeHook    T
	endHook       T
	registerAttr  map[string]registerEvolveAttr[T]
	registerAlias map[string][]string
	param         *Param
	namingMap     map[string]any
}

type registerEvolveAttr[T any] struct {
	cli.CommandOptional
	runnable T
}

// Command When registering a method you must specify commands to run more than one.
// We agreed that the second and subsequent commands should be aliases for the first command.
func (e *Evolve[T]) Command(t T, command string, optionals ...cli.CommandOptional) cli.Application[T] {
	e.CommandList(t, []string{command}, optionals...)
	return e
}

// CommandList When registering a method you must specify commands to run more than one.
// We agreed that the second and subsequent commands should be aliases for the first command.
func (e *Evolve[T]) CommandList(t T, commands []string, optionals ...cli.CommandOptional) cli.Application[T] {
	vNum := len(commands)
	if vNum == 0 {
		log.Fatal("Evolve.Command: when registering a method you must specify commands to run more than one. ")
		return e
	}

	optional := rock.Param(cli.CommandOptional{}, optionals...)
	optional.Alias = commands[1:]
	optional.Keys = commands
	attr := registerEvolveAttr[T]{
		CommandOptional: optional,
		runnable:        t,
	}
	// remember the command of alias.
	mainCmd := commands[0]
	e.registerAttr[mainCmd] = attr
	if vNum == 1 {
		return e
	}

	e.registerAlias[mainCmd] = commands[1:]
	return e
}

func (e *Evolve[T]) Index(t T) cli.Application[T] {
	e.indexTodo = t
	return e
}

func (e *Evolve[T]) Lost(t T) cli.Application[T] {
	e.lostTodo = t
	return e
}

func (e *Evolve[T]) Before(t T) cli.Application[T] {
	e.beforeHook = t
	return e
}

func (e *Evolve[T]) End(t T) cli.Application[T] {
	e.endHook = t
	return e
}

func (e *Evolve[T]) Run(args ...string) error {
	e.param = NewParam(args...)
	return e.routerCli()
}

func (e *Evolve[T]) RunArgs(args cli.ArgsParser) error {
	e.param = NewArgs(args)
	return e.routerCli()
}

func (e *Evolve[T]) callFunc(fn reflect.Value) bool {
	fnVal := fn.Interface()
	isSuccess := false
	switch callValue := fnVal.(type) {
	case func():
		callValue()
		isSuccess = true
	case func(...string):
		callValue()
		isSuccess = true
	case func(cli.ArgsParser):
		callValue(e.param.Args)
		isSuccess = true
	case func(...cli.ArgsParser):
		callValue(e.param.Args)
		isSuccess = true
	}
	return isSuccess
}

// to run register instance
func (e *Evolve[T]) toRunRg(rg T) bool {
	rv := reflect.ValueOf(rg)
	if !rv.IsValid() || rv.IsZero() || rv.IsNil() {
		return false
	}

	vStruct := rv
	if rv.Kind() == reflect.Ptr {
		vStruct = rv.Elem()
	}

	if vStruct.Kind() == reflect.Struct {
		args := e.param.Args
		sumCommand := args.SubCommand()
		runMth := func(name string) bool {
			mth := rv.MethodByName(name)
			if mth.IsValid() {
				return e.callFunc(mth)
			}
			return false
		}

		// set field
		field := vStruct.FieldByName(CmdFidX)
		if field.IsValid() {
			field.Set(reflect.ValueOf(e.param))
		}

		runMth(CmdMtdInit)
		isHelpCmd := sumCommand == "help" || sumCommand == "?"
		if isHelpCmd || (sumCommand == "" && args.Switch("help", "h", "?")) {
			runMth(CmdMtdHelp)
		} else if sumCommand == "" {
			runMth(CmdMtdIndex)
		} else {
			if !runMth(str.Str(sumCommand).Ucfirst()) {
				runMth(CmdMtdLost)
			}
		}

		return true
	}

	if !rv.CanInterface() {
		return false
	}
	isRun := e.callFunc(rv)
	return isRun
}

func (e *Evolve[T]) runIndex() {
	if e.toRunRg(e.indexTodo) {
		return
	}

	config := e.config
	config.IndexDoc()
}

// find reg support both name and alias.
func (e *Evolve[T]) findReg(name string) (reg registerEvolveAttr[T], isFind bool) {
	reg, isFind = e.registerAttr[name]
	if isFind {
		return
	}

	for key, alias := range e.registerAlias {
		if rock.InList(alias, name) {
			isFind = true
			reg = e.registerAttr[key]
			return
		}
	}
	return
}

func (e *Evolve[T]) routerCli() error {
	param := e.param
	config := e.config
	args := param.Args

	if !config.DisableHelp {
		command := args.Command()
		isHelp := command == "help" || command == "?"
		if !isHelp {
			helpSwitch := args.Switch("help", "h", "?")
			isHelp = command == "" && helpSwitch
		}
		if isHelp {
			e.runHelp()
			return nil
		}
	}

	command := args.Command()
	if command == "" {
		e.runIndex()
		return nil
	}
	naming := e.NamingFind()
	if naming != "" {
		command = naming
	}

	rg, match := e.findReg(command)
	if match {
		args.SetOptional(&rg.CommandOptional)
		e.toRunRg(e.beforeHook)
		if !config.DisableHelp {
			invalidMsg := rg.InvalidMsg(args)
			if invalidMsg != "" {
				lgr.Error(invalidMsg)
				return nil
			}
		}
		if e.toRunRg(rg.runnable) {
			e.toRunRg(e.endHook)
		}
		return nil
	}

	if e.toRunRg(e.lostTodo) {
		return nil
	}

	fmt.Println()
	fmt.Printf("%s: We gotta lost, honey!\n    Uymas@%s/%s\n", command, uymas.Version, uymas.Release)
	fmt.Println()
	return nil
}

func (e *Evolve[T]) Help(t T) cli.Application[T] {
	e.helpTodo = t
	return e
}

func (e *Evolve[T]) runHelp() {
	if e.toRunRg(e.helpTodo) {
		return
	}

	cmdName := e.param.Args.HelpCmd()
	helpMsg, isFind := e.GetHelp(cmdName)
	if isFind {
		fmt.Println(helpMsg)
		fmt.Println()
		return
	}

	if cmdName != "" {
		lgr.Warn("命令 " + cmdName + " 不存在")
	}
}

// Naming manually set the named mapping to be non-alias, v support: `string`/`func(Param) string`
func (e *Evolve[T]) Naming(name string, v any) *Evolve[T] {
	e.namingMap[name] = v
	return e
}

// NamingFind default (when no parameters are specified) top-level command level
func (e *Evolve[T]) NamingFind(cmds ...string) string {
	param := e.param
	if param == nil {
		return ""
	}
	args := param.Args
	name := rock.Param(args.Command(), cmds...)
	if name == "" {
		return ""
	}

	value, exist := e.namingMap[name]
	if !exist {
		return ""
	}

	switch vRel := value.(type) {
	case string:
		return vRel
	case func(Param) string:
		return vRel(*param)
	}

	return ""
}

func (e *Evolve[T]) GetHelp(cmd string) (helpMsg string, exits bool) {
	if cmd == "" {
		var lines []string
		keys := rock.MapKeys(e.registerAttr)
		maxLen := str.QueueMaxLen(keys)
		sort.Strings(keys)
		for _, key := range keys {
			reg := e.registerAttr[key]
			cmdHelp := reg.Help
			if cmdHelp == "" {
				cmdHelp = "这是 " + key + " 命令（默认）"
			}
			if len(reg.Alias) > 0 {
				cmdHelp += "，别名支持 " + strings.Join(reg.Alias, ",")
			}
			line := fmt.Sprintf("%-"+(fmt.Sprintf("%d", maxLen+8))+"s", key) + cmdHelp
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
	reg, hasCmd := e.registerAttr[cmd]
	if !hasCmd {
		for fName, fReg := range e.registerAttr {
			if rock.InList(fReg.Keys, cmd) {
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

func NewEvolve(cfgs ...cli.Config) cli.Application[any] {
	evl := &Evolve[any]{
		config:        rock.Param(cli.DefaultConfig, cfgs...),
		registerAlias: map[string][]string{},
		registerAttr:  map[string]registerEvolveAttr[any]{},
		namingMap:     map[string]any{},
	}
	return evl
}
