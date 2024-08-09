package evolve

import (
	"fmt"
	"gitee.com/conero/uymas/v2"
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/str"
	"reflect"
)

type Evolve[T any] struct {
	config      cli.Config
	indexTodo   T
	lostTodo    T
	beforeHook  T
	endHook     T
	registerMap map[string]T
	param       *Param
}

func (e *Evolve[T]) Command(t T, commands ...string) cli.Application[T] {
	for _, cmd := range commands {
		e.registerMap[cmd] = t
	}
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
	switch fnVal.(type) {
	case func():
		fnVal.(func())()
		isSuccess = true
	case func(cli.ArgsParser):
		fnVal.(func(cli.ArgsParser))(e.param.Args)
		isSuccess = true
	case func(...cli.ArgsParser):
		fnVal.(func(...cli.ArgsParser))(e.param.Args)
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
		if sumCommand == "help" || (sumCommand == "" && args.Switch("help", "h")) {
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
	command := param.Args.Command()
	if command == "" {
		e.runIndex()
		return nil
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

func NewEvolve() cli.Application[any] {
	evl := &Evolve[any]{
		registerMap: map[string]any{},
	}
	return evl
}
