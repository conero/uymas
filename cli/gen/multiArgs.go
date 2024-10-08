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
//
// MultiArgs(args cli.ArgsParser, target any, seq string='.', pref string=”): support all param like
func MultiArgs(args cli.ArgsParser, target any, params ...string) error {
	if target == nil {
		return errors.New("target is nil")
	}
	var refVal reflect.Value
	tgtVal, isVal := target.(reflect.Value)
	if isVal {
		refVal = tgtVal
	} else {
		refVal = reflect.ValueOf(target)
	}
	elem := refVal
	if refVal.Kind() == reflect.Ptr {
		elem = refVal.Elem()
	}

	elKind := elem.Kind()
	if elKind == reflect.Map {
		return MultiArgsMap(args, target, params...)
	}
	if elKind != reflect.Struct {
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
			err := MultiArgs(args, field, seq, key)
			if err != nil {
				return err
			}
			continue
		}

		setValueByStr(field, []string{key}, args)
	}

	return nil
}

// MultiArgsMap Assign cli.ArgsParser to map
func MultiArgsMap(args cli.ArgsParser, mapTgt any, params ...string) error {
	seq := rock.ParamIndex(1, ".", params...)
	pref := rock.ParamIndex(2, "", params...)

	refVal := reflect.ValueOf(mapTgt)
	elem := refVal
	if refVal.Kind() == reflect.Ptr {
		elem = refVal.Elem()
	}
	if elem.Kind() != reflect.Map {
		return errors.New("mapTgt is not valid Map")
	}

	values := args.Values()
	for key, _ := range values {
		index := strings.Index(key, seq)
		if index < 1 {
			continue
		}

		// eg `[key].[subkey]`, '-test.name Jc'
		fullKey := pref + key
		mapKey := fullKey[:index]
		subKey := fullKey[index+1:]

		vMapKey := reflect.ValueOf(mapKey)
		vMapValue := elem.MapIndex(vMapKey)
		if !vMapValue.IsValid() || vMapValue.IsNil() || vMapValue.IsZero() {
			var subMap = map[string]string{
				subKey: args.Get(key),
			}
			elem.SetMapIndex(vMapKey, reflect.ValueOf(subMap))
			continue
		}

		if vMapValue.Kind() != reflect.Map {
			continue
		}
		vMapValue.SetMapIndex(reflect.ValueOf(subKey), reflect.ValueOf(args.Get(key)))

	}

	return nil
}
