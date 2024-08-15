// Package evolve version Command line, which supports more registration types than cli. Adopting reflection.
package evolve

import (
	"fmt"
	"gitee.com/conero/uymas/v2"
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/rock"
	"gitee.com/conero/uymas/v2/str"
	"log"
	"reflect"
)

type Evolve[T any] struct {
	config        cli.Config
	indexTodo     T
	lostTodo      T
	helpTodo      T
	beforeHook    T
	endHook       T
	registerMap   map[string]T
	registerAlias map[string][]string
	param         *Param
	namingMap     map[string]any
}

// Command When registering a method you must specify commands to run more than one.
// We agreed that the second and subsequent commands should be aliases for the first command.
func (e *Evolve[T]) Command(t T, commands ...string) cli.Application[T] {
	vNum := len(commands)
	if vNum == 0 {
		log.Fatal("Evolve.Command: when registering a method you must specify commands to run more than one. ")
		return e
	}

	for _, cmd := range commands {
		e.registerMap[cmd] = t
	}
	if vNum == 1 {
		return e
	}

	// remember the command of alias.
	mainCmd := commands[0]
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

	fmt.Println()
	fmt.Println("-------------- Uymas/Evolution -----------------")
	fmt.Println("Welcome to our world")
	fmt.Printf(":)- %s/%s\n", uymas.Version, uymas.Release)
	fmt.Println()
}

func (e *Evolve[T]) routerCli() error {
	param := e.param
	config := e.config
	args := param.Args

	if !config.DisableHelp {
		command := args.Command()
		isHelp := command == "help" || command == "?"
		if isHelp || args.Switch("help", "h") {
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

	rg, match := e.registerMap[command]
	if match {
		e.toRunRg(e.beforeHook)
		if e.toRunRg(rg) {
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

	args := e.param.Args
	command := args.Command()
	cmdName := args.HelpCmd()
	if cmdName != "" {
		command = "<" + command + " " + cmdName + ">"
	}
	fmt.Printf("Default Help: we should add the help information for command %s here, honey!\n\n", command)
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

func NewEvolve() cli.Application[any] {
	evl := &Evolve[any]{
		registerMap:   map[string]any{},
		registerAlias: map[string][]string{},
		namingMap:     map[string]any{},
	}
	return evl
}
