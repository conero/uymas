package evolve

import (
	"fmt"
	"gitee.com/conero/uymas/v2"
	"gitee.com/conero/uymas/v2/cli"
	"reflect"
)

type Evolve[T any] struct {
	config      cli.Config
	indexTodo   T
	lostTodo    T
	beforeHook  T
	endHook     T
	registerMap map[string]T
	param       cli.ArgsParser
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

// to run register instance
func (e *Evolve[T]) toRunRg(rg T) bool {
	rv := reflect.ValueOf(rg)
	if !rv.IsValid() || rv.IsZero() || rv.IsNil() {
		return false
	}

	if !rv.CanInterface() {
		return false
	}
	vAny := rv.Interface()
	isRun := false
	switch vAny.(type) {
	case func():
		vAny.(func())()
		isRun = true
	case func(cli.ArgsParser):
		vAny.(func(cli.ArgsParser))(e.param)
		isRun = true
	case func(...cli.ArgsParser):
		vAny.(func(...cli.ArgsParser))(e.param)
		isRun = true
	}
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
	command := param.Command()
	if command == "" {
		e.runIndex()
		return nil
	}
	return nil
}

func NewEvolve() cli.Application[any] {
	evl := &Evolve[any]{
		registerMap: map[string]any{},
	}
	return evl
}
