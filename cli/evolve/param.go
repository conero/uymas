package evolve

import (
	"gitee.com/conero/uymas/v2/cli"
)

type Param struct {
	cli.Args
}

func NewParam(args ...string) cli.ArgsParser {
	arg := cli.NewArgs(args...)
	param := &Param{}
	param.ArgsParser = arg
	return param
}

func NewParamWith(config cli.ArgsConfig, args ...string) cli.ArgsParser {
	arg := cli.NewArgsWith(config, args...)
	param := &Param{}
	param.ArgsParser = arg
	return param
}
