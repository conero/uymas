package gen

import (
	"errors"
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/data/convert"
	"gitee.com/conero/uymas/v2/str"
	"reflect"
	"strings"
)

// ArgsTagName The struct tag is named cmd. If this parameter is not set,
// the struct tag is divided into lowercase cases, such as file_name, make_up.
//
// json tag is also supported when cmd is not set.
//
// syntax rules: "cmd>>json>>FileName".
const ArgsTagName = "cmd"

// ArgsDress Dress the command argument up on the specified data entity (struct)
func ArgsDress(args cli.ArgsParser, data any) error {
	ref := reflect.ValueOf(data)
	isStruct := ref.Kind() == reflect.Struct
	isPtr := false
	if ref.Kind() == reflect.Ptr {
		isStruct = ref.Elem().Kind() == reflect.Struct
		isPtr = true
	}

	if !isStruct {
		return errors.New("data: the param of ArgsDress only support struct")
	}

	rtp := ref.Type()
	if isPtr {
		rtp = ref.Elem().Type()
	}

	for i := 0; i < rtp.NumField(); i++ {
		fieldType := rtp.Field(i)
		name := fieldType.Tag.Get(ArgsTagName)
		if name == "" {
			name = fieldType.Tag.Get("json")
		}
		if name == "" {
			name = str.Str(fieldType.Name).LowerStyle()
		}

		var vFiled reflect.Value
		if isPtr {
			vFiled = ref.Elem().Field(i)
		} else {
			vFiled = ref.Field(i)
		}

		keys := strings.Split(str.Str(name).ClearSpace(), ",")
		vfKind := vFiled.Kind()

		if vfKind == reflect.Bool {
			vFiled.SetBool(args.Switch(keys...))
			continue
		}
		if vfKind == reflect.Slice {
			convert.SetByStrSlice(vFiled, args.List(keys...))
			continue
		}

		value := args.Get(keys...)
		convert.SetByStr(vFiled, value)
	}

	return nil
}

// ArgsDecompose Decompose the structure into an option list
// todo needTodo
func ArgsDecompose(data any) []cli.Option {
	panic("todo")
	//return nil
}
