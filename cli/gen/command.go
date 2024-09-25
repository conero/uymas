package gen

import (
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/cli/evolve"
	"gitee.com/conero/uymas/v2/str"
	"reflect"
)

type Runnable interface {
	cli.Fn | func()
}

type StructCmd struct {
	indexRn     any
	lostRn      any
	initRn      any
	commandList map[string]any
}

func (s *StructCmd) Index() any {
	return s.indexRn
}

func (s *StructCmd) Lost() any {
	return s.lostRn
}

func (s *StructCmd) Init() any {
	return s.initRn
}

func ParseStruct(vStruct any) *StructCmd {
	if vStruct == nil {
		return nil
	}
	rv := reflect.ValueOf(vStruct)
	realVal := rv
	if rv.Kind() == reflect.Ptr {
		realVal = rv.Elem()
	}
	if realVal.Kind() != reflect.Struct {
		return nil
	}

	num := realVal.NumMethod()
	rType := realVal.Type()
	sc := &StructCmd{}
	for i := 0; i < num; i++ {
		method := rType.Method(i)
		methodValue := realVal.Method(i)
		if !methodValue.CanInterface() {
			continue
		}
		name := method.Name
		vAny := methodValue.Interface()
		switch name {
		case evolve.CmdMtdIndex:
			sc.indexRn = vAny
		case evolve.CmdMtdLost:
			sc.lostRn = vAny
		case evolve.CmdMtdInit:
			sc.initRn = vAny
		default:
			sc.commandList[str.Str(name).Lcfirst()] = vAny
		}
	}

	return sc
}
