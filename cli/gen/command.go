package gen

import (
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/cli/evolve"
	"gitee.com/conero/uymas/v2/rock"
	"gitee.com/conero/uymas/v2/str"
	"reflect"
)

type Runnable interface {
	cli.Fn | func()
}

type StructCmdAttr struct {
	Title     string
	Option    []cli.Option
	FieldName string
}

type StructCmd struct {
	indexRn      any
	lostRn       any
	initRn       any
	commandList  map[string]any
	commandAttr  map[string]StructCmdAttr
	globalOption []cli.Option
	contextVal   reflect.Value
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

func (s *StructCmd) parseStructOption() {
	if s.commandAttr == nil {
		s.commandAttr = map[string]StructCmdAttr{}
	}
	rv := s.contextVal
	rlStruct := rv
	if rv.Kind() == reflect.Ptr {
		rlStruct = rv.Elem()
	}
	num := rlStruct.NumField()
	rType := rlStruct.Type()
	for i := 0; i < num; i++ {
		field := rlStruct.Field(i)
		sf := rType.Field(i)

		cmdStr := sf.Tag.Get(ArgsTagName)
		if cmdStr == "" {
			continue
		}

		opt := OptionTagParse(cmdStr)
		if opt == nil {
			continue
		}

		if rock.InList(opt.List, ArgsGlobalOwner) {
			s.globalOption = append(s.globalOption, *opt)
			continue
		}

		owner := opt.Owner
		if owner == "" {
			continue
		}

		s.commandAttr[owner] = StructCmdAttr{
			Title:     opt.Help,
			Option:    StructDress(field),
			FieldName: sf.Name,
		}
		//if field.Kind() == reflect.Struct {
		//	fmt.Println("Child Struct")
		//}
	}
}

func (s *StructCmd) GetOptions(cmd string) *StructCmdAttr {
	options := s.globalOption
	vAttr, exist := s.commandAttr[cmd]
	if exist {
		if len(options) > 0 {
			vAttr.Option = append(vAttr.Option, options...)
		}
		return &vAttr
	}
	return nil
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

	// method parse handler
	num := rv.NumMethod()
	rType := rv.Type()
	sc := &StructCmd{
		commandList: map[string]any{},
	}
	sc.contextVal = rv
	for i := 0; i < num; i++ {
		method := rType.Method(i)
		methodValue := rv.Method(i)
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

	// option parse handler
	sc.parseStructOption()
	return sc
}

// SetArgs 设置option
func (s *StructCmd) SetArgs(args cli.ArgsParser) {
	if s.contextVal.IsNil() || !s.contextVal.IsValid() || s.contextVal.IsZero() {
		return
	}
	value := s.contextVal
	if s.contextVal.Kind() == reflect.Ptr {
		value = s.contextVal.Elem()
	}
	s.setOption(args, value)
	field := value.FieldByName(evolve.CmdFidArgs)
	if !field.CanSet() {
		return
	}
	field.Set(reflect.ValueOf(args))
}

func (s *StructCmd) setOption(args cli.ArgsParser, target reflect.Value) {
	// set global option
	for _, opt := range s.globalOption {
		if opt.FieldName == "" {
			continue
		}
		vField := target.FieldByName(opt.FieldName)
		if !vField.IsValid() {
			continue
		}

		setValueByOption(vField, &opt, args, nil)
	}

	command := args.Command()
	if command == "" {
		return
	}

	attr, exist := s.commandAttr[command]
	if !exist {
		return
	}

	vField := target.FieldByName(attr.FieldName)
	if !vField.IsValid() {
		return
	}
	setToStruct(vField, args)
}

func AsCommand(vStruct any, cfgs ...cli.Config) cli.Application[any] {
	pCmd := ParseStruct(vStruct)
	if pCmd == nil {
		panic("vStruct is not struct, and parse fail")
	}
	evl := evolve.NewEvolve(cfgs...)
	evl.Lost(pCmd.Lost())
	evl.Index(pCmd.Index())
	evl.RouterBefore(func(args cli.ArgsParser) {
		pCmd.SetArgs(args)
	})

	for vCmd, runnable := range pCmd.commandList {
		vAttr := pCmd.GetOptions(vCmd)
		if vAttr != nil {
			evl.Command(runnable, vCmd, cli.CommandOptional{
				Help:    vAttr.Title,
				Options: vAttr.Option,
			})
		} else {
			evl.Command(runnable, vCmd)
		}

	}
	return evl
}
