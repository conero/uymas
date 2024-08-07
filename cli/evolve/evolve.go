package evolve

import "gitee.com/conero/uymas/v2/cli"

type Evolve[T any] struct {
}

func (e *Evolve[T]) Command(t T, commands ...string) cli.Application[T] {
	//TODO implement me
	panic("implement me")
}

func (e *Evolve[T]) Index(t T) cli.Application[T] {
	//TODO implement me
	panic("implement me")
}

func (e *Evolve[T]) Lost(t T) cli.Application[T] {
	//TODO implement me
	panic("implement me")
}

func (e *Evolve[T]) Before(t T) cli.Application[T] {
	//TODO implement me
	panic("implement me")
}

func (e *Evolve[T]) End(t T) cli.Application[T] {
	//TODO implement me
	panic("implement me")
}

func (e *Evolve[T]) Run(args ...string) error {
	//TODO implement me
	panic("implement me")
}

func NewEvolve() cli.Application[int] {
	evl := &Evolve[int]{}
	return evl
}
