package evolve

import (
	"gitee.com/conero/uymas/v2/cli"
)

type Param struct {
	Args cli.ArgsParser
}

func NewParam(args ...string) *Param {
	arg := cli.NewArgs(args...)
	param := &Param{}
	param.Args = arg
	return param
}

func NewArgs(args cli.ArgsParser) *Param {
	param := &Param{}
	param.Args = args
	return param
}

func NewParamWith(config cli.ArgsConfig, args ...string) *Param {
	arg := cli.NewArgsWith(config, args...)
	param := &Param{}
	param.Args = arg
	return param
}
