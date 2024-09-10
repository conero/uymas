// Package evolve version Command line, which supports more registration types than cli. Adopting reflection.
package evolve

import (
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/logger/lgr"
	"gitee.com/conero/uymas/v2/rock"
	"gitee.com/conero/uymas/v2/str"
	"reflect"
)

// Evolve The following registration types are supported:
//
// 1. struct(inherited Command)
//
// 2. func(cli.ArgsParser)
//
// 3. func(...cli.ArgsParser)
//
// 4. func()
//
// 5. func(...string)
type Evolve[T any] struct {
	cli.Register[T]
}

type registerEvolveAttr[T any] struct {
	cli.CommandOptional
	runnable T
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
		callValue(e.Args())
		isSuccess = true
	case func(...cli.ArgsParser):
		callValue(e.Args())
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
		args := e.Args()
		sumCommand := args.SubCommand()
		runMth := func(name string) bool {
			mth := rv.MethodByName(name)
			if mth.IsValid() {
				return e.callFunc(mth)
			}
			return false
		}

		// set field
		field := vStruct.FieldByName(CmdFidArgs)
		if field.IsValid() {
			field.Set(reflect.ValueOf(args))
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

func NewEvolve(cfgs ...cli.Config) cli.Application[any] {
	evl := &Evolve[any]{}
	evl.Config = rock.Param(cli.DefaultConfig, cfgs...)
	evl.Call = func(fn any, parser cli.ArgsParser) {
		if fn == nil {
			return
		}
		if evl.toRunRg(fn) {
			return
		}
		lgr.Error("%s: 命令注册实例无效，其不可执行（%#v）", parser.Command(), fn)
	}
	evl.Help(evl.GenerateHelpFn)
	return evl
}
