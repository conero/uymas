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

	// to set Value
	toSetValue := func(keys []string, value string) {
		countKey := len(keys)
		//fmt.Printf("keys: %v, value: %v\n", keys, value)
		for cIdx, ccKey := range keys {
			//fmt.Printf("cIdx: %v, ccKey: %v\n", cIdx, ccKey)
			switch cIdx {
			case 0: // 顶级
				if countKey == 1 {
					continue
				}
				vMapKey := reflect.ValueOf(ccKey)
				vMapValue := elem.MapIndex(vMapKey)
				if !vMapValue.IsValid() || vMapValue.IsZero() || vMapValue.IsNil() {
					elem.SetMapIndex(vMapKey, reflect.ValueOf(map[string]any{}))
				}
			case 1:
				parentKey := reflect.ValueOf(keys[cIdx-1])
				vMapValue := elem.MapIndex(parentKey)
				if !vMapValue.IsValid() || vMapValue.IsZero() || vMapValue.IsNil() {
					elem.SetMapIndex(parentKey, reflect.ValueOf(map[string]any{}))
					vMapValue = elem.MapIndex(parentKey)
				}

				if countKey == 2 {
					//vMapValue.SetMapIndex(parentKey, reflect.ValueOf(value))
					if vMapValue.Kind() == reflect.Map {
						vMapValue.SetMapIndex(reflect.ValueOf(ccKey), reflect.ValueOf(value))
					} else if vMapValue.Kind() == reflect.Interface {
						instance := vMapValue.Elem()
						if instance.IsValid() {
							instance.SetMapIndex(reflect.ValueOf(ccKey), reflect.ValueOf(value))
							//fmt.Printf("instance: %v\n", instance)
							//elem.SetMapIndex(vMapKey, reflect.ValueOf(instance.Interface()))
							//elem.SetMapIndex(vMapKey, instance)
						}
					}
				}
			case 2:
				parentKey := reflect.ValueOf(keys[cIdx-1])
				parentElem := elem.MapIndex(reflect.ValueOf(keys[cIdx-2]))
				if parentElem.Kind() == reflect.Interface {
					parentElem = parentElem.Elem()
				}
				if parentElem.Kind() != reflect.Map {
					continue
				}
				vMapValue := parentElem.MapIndex(parentKey)
				if !vMapValue.IsValid() || vMapValue.IsZero() || vMapValue.IsNil() {
					parentElem.SetMapIndex(parentKey, reflect.ValueOf(map[string]any{}))
					vMapValue = parentElem.MapIndex(parentKey)
				}

				if countKey == 3 {
					//vMapValue.SetMapIndex(parentKey, reflect.ValueOf(value))
					if vMapValue.Kind() == reflect.Map {
						vMapValue.SetMapIndex(reflect.ValueOf(ccKey), reflect.ValueOf(value))
					} else if vMapValue.Kind() == reflect.Interface {
						instance := vMapValue.Elem()
						if instance.IsValid() {
							instance.SetMapIndex(reflect.ValueOf(ccKey), reflect.ValueOf(value))
							//fmt.Printf("instance: %v\n", instance)
							//elem.SetMapIndex(vMapKey, reflect.ValueOf(instance.Interface()))
							//elem.SetMapIndex(vMapKey, instance)
						}
					}
				}
			}
		}
	}

	values := args.Values()
	for key, _ := range values {
		fullKey := pref + key
		index := strings.Index(fullKey, seq)
		if index < 1 {
			continue
		}

		keys := strings.Split(fullKey, seq)
		value := args.Get(key)
		toSetValue(keys, value)
	}

	return nil
}
