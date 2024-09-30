package gen

import (
	"errors"
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/rock"
	"gitee.com/conero/uymas/v2/str"
	"reflect"
	"strings"
)

func isStruct(value reflect.Value) bool {
	elem := value
	if value.Kind() == reflect.Ptr {
		elem = value.Elem()
	}
	return elem.Kind() == reflect.Struct
}

// MultiArgs Multi Args value parsing
func MultiArgs(args cli.ArgsParser, tgt any, params ...string) error {
	refVal := reflect.ValueOf(tgt)
	elem := refVal
	if refVal.Kind() == reflect.Ptr {
		elem = refVal.Elem()
	}

	if elem.Kind() != reflect.Struct {
		return errors.New("target is not struct type")
	}

	seq := rock.ParamIndex(1, ".", params...)
	pref := rock.ParamIndex(2, "", params...)
	var keyList []string
	if pref != "" {
		keyList = append(keyList, pref)
	}
	fieldNum := elem.NumField()
	refType := elem.Type()
	for i := 0; i < fieldNum; i++ {
		field := elem.Field(i)
		sf := refType.Field(i)
		name := str.Str(sf.Name).LowerStyle()
		key := strings.Join(append(keyList, name), seq)
		if isStruct(field) {
			MultiArgs(args, field, seq, key)
			continue
		}

		setValueByStr(field, []string{name}, args)
	}

	return nil
}
