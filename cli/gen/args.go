package gen

import (
	"errors"
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/data/convert"
	"gitee.com/conero/uymas/v2/rock"
	"gitee.com/conero/uymas/v2/str"
	"reflect"
	"regexp"
	"strings"
)

// ArgsTagName The struct tag is named cmd. If this parameter is not set,
// the struct tag is divided into lowercase cases, such as file_name, make_up.
//
// json tag is also supported when cmd is not set.
//
// syntax rules of name: "cmd>>json>>FileName".
const ArgsTagName = "cmd"
const ArgsCmdRequired = "required"
const ArgsCmdHelp = "help"
const ArgsCmdDefault = "default"
const ArgsTagOmit = "-"

func argsValueCheck(ref reflect.Value) (reflect.Value, error) {
	isStruct := ref.Kind() == reflect.Struct
	isPtr := false
	if ref.Kind() == reflect.Ptr {
		isStruct = ref.Elem().Kind() == reflect.Struct
		isPtr = true
	}

	if !isStruct {
		return reflect.Value{}, errors.New("data: the param of ArgsDress only support struct")
	}

	rValue := ref
	if isPtr {
		rValue = ref.Elem()
	}

	return rValue, nil
}

// ArgsDress Dress the command argument up on the specified data entity (struct)
func ArgsDress(args cli.ArgsParser, data any) error {
	ref := reflect.ValueOf(data)
	realValue, err := argsValueCheck(ref)
	if err != nil {
		return err
	}
	rtp := realValue.Type()

	for i := 0; i < rtp.NumField(); i++ {
		fieldType := rtp.Field(i)
		name := fieldType.Tag.Get(ArgsTagName)
		if name == "" {
			name = fieldType.Tag.Get("json")
		}
		if name == "" {
			name = str.Str(fieldType.Name).LowerStyle()
		}

		if name == ArgsTagOmit {
			continue
		}
		keys := getNameByTag(name)

		var vFiled = realValue.Field(i)
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
func ArgsDecompose(data any, excludes ...string) ([]cli.Option, error) {
	ref := reflect.ValueOf(data)
	realValue, err := argsValueCheck(ref)
	if err != nil {
		return nil, err
	}
	rtp := realValue.Type()

	var optionList []cli.Option
	for i := 0; i < rtp.NumField(); i++ {
		fieldType := rtp.Field(i)
		cmdTag := fieldType.Tag.Get(ArgsTagName)
		if cmdTag == ArgsTagOmit {
			continue
		}
		option := OptionTagParse(cmdTag)
		var name string
		if option == nil {
			if name == "" {
				name = fieldType.Tag.Get("json")
			}
			if name == "" {
				name = str.Str(fieldType.Name).LowerStyle()
			}
			if rock.InList(excludes, name) {
				continue
			}
			option = &cli.Option{
				Alias: []string{name},
			}
		} else if rock.InList(excludes, option.Name) {
			continue
		}

		optionList = append(optionList, *option)
	}
	return optionList, nil
}

func ArgsDecomposeMust(data any, excludes ...string) []cli.Option {
	opts, _ := ArgsDecompose(data, excludes...)
	return opts
}

// OptionTagParse Resolves the value of the tag into an option object
//
// syntax rules of tag: `"name,n required default:111 help:help msg"`.
func OptionTagParse(vTag string) *cli.Option {
	if vTag == "" {
		return nil
	}
	spaceList := regexp.MustCompile(`\s{2,}`)
	vTag = spaceList.ReplaceAllString(vTag, " ")
	vTag = strings.TrimSpace(vTag)
	if vTag == "" {
		return nil
	}

	option := &cli.Option{}
	for i, s := range strings.Split(vTag, " ") {
		if i == 0 && !strings.Contains(s, ":") {
			option.Alias = strings.Split(str.Str(s).ClearSpace(), ",")
			continue
		}
		if s == ArgsCmdRequired {
			option.Require = true
			continue
		}

		idx := strings.Index(s, ":")
		if idx > 0 {
			key := s[:idx]
			value := s[idx+1:]
			switch key {
			case ArgsCmdHelp:
				option.Help = str.Str(value).Unescape()
			case ArgsCmdDefault:
				option.DefValue = str.Str(value).Unescape()
			}
		}
	}

	return option
}

// Get the name by parsing the tag of struct, format like `cmd:"name,n"`
func getNameByTag(tag string) []string {
	tag = strings.TrimSpace(tag)
	if tag == "" {
		return nil
	}
	var name string
	for _, vs := range strings.Split(tag, " ") {
		vs = strings.TrimSpace(vs)
		if vs == "" {
			continue
		}
		if strings.Contains(vs, ":") {
			continue
		}

		name = vs
		break

	}

	if len(name) > 0 {
		return strings.Split(str.Str(name).ClearSpace(), ",")
	}
	return []string{tag}
}
