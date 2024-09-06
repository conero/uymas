package cli

import (
	"fmt"
	"gitee.com/conero/uymas/v2/rock"
	"log"
)

type RegisterMeta[T any] struct {
	Command  CommandOptional
	Runnable T
}

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
	Args          ArgsParser
}

func (r *Register[T]) Command(t T, command string, optionals ...CommandOptional) Application[T] {
	r.CommandList(t, []string{command}, optionals...)
	return r
}

func (r *Register[T]) CommandList(t T, commands []string, optionals ...CommandOptional) Application[T] {
	vNum := len(commands)
	if vNum == 0 {
		log.Fatal("Evolve.Command: when registering a method you must specify commands to run more than one. ")
		return r
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

func (r *Register[T]) Run(args ...string) error {
	r.Args = NewArgs(args...)
	if r.registerAlias != nil {
		r.registerAlias = map[string][]string{}
	}
	return nil
}

func (r *Register[T]) RunArgs(args ArgsParser) error {
	r.Args = args
	return r.Run()
}
