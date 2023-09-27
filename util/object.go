package util

import (
	"encoding/json"
	"gitee.com/conero/uymas/str"
	"reflect"
)

type Object struct {
}

// Assign @todo
//
//	Base of `reflect` to come true like javascript `Object.Assign`, target should be pointer best.
//	It can be Multiple, only for `reflect.Map`.
func (obj Object) Assign(target any, source any) any {
	var m = target
	tReft := reflect.TypeOf(target)
	if tReft.Kind() == reflect.Ptr {
		tReft = tReft.Elem()
	}
	tRefv := reflect.ValueOf(target)
	if tRefv.Kind() == reflect.Ptr {
		tRefv = tRefv.Elem()
	}
	//if it's map that can add field
	isMap := tReft.Kind() == reflect.Map
	if isMap {
		obj.AssignMap(target, source)
		return target
	}

	if tReft.Kind() != reflect.Struct {
		return target
	}
	sRefv := reflect.ValueOf(source)
	num := tReft.NumField()
	for i := 0; i < num; i++ {
		field := tReft.Field(i)
		sField := sRefv.FieldByName(field.Name)
		tField := tRefv.Field(i)
		if sField.IsValid() && !sField.IsZero() && sField.Kind() == tField.Kind() {
			if sField.Kind() == reflect.Struct { // Nesting Assign
				//Structure nesting handler
				//@todo <Nesting Assign>
				//panic: reflect: Elem of invalid type reflect.Value
				//fmt.Println(field.Name)
				//if tField.CanAddr() {
				//fmt.Printf("Nest->tField %#v\n", tField)
				//fmt.Printf("Nest->sField %#v\n", sField)
				//obj.Assign(tField.Addr(), sField)
				//obj.Assign(tField.Addr(), sField)
				//}

				tfValue := tField.Interface()
				sfValue := sField.Interface()
				//fmt.Printf("tfValue: %#v, sfValue:%#v\n", tfValue, sfValue)
				obj.Assign(&tfValue, sfValue)
				//fmt.Printf("tfValue: %#v, sfValue:%#v\n", tfValue, sfValue)
				tField.Set(reflect.ValueOf(tfValue))
			} else {
				tField.Set(sField)
			}
		}
	}

	return m
}

// AssignMap Assign Map/Struct to map
func (obj Object) AssignMap(targetMap any, srcMapOrStruct any) {
	tVal := reflect.ValueOf(targetMap)
	sVal := reflect.ValueOf(srcMapOrStruct)
	tKind := tVal.Kind()
	if tKind == reflect.Map {
		sKind := sVal.Kind()
		if tKind == sKind {
			rg := sVal.MapRange()
			for rg.Next() {
				sk := rg.Key()
				sV := rg.Value()
				if !sV.IsNil() {
					tVal.SetMapIndex(sk, sV)
				}
			}
		} else if sKind == reflect.Struct {
			sVal = sVal.Elem()
			num := sVal.NumField()
			sTp := reflect.TypeOf(srcMapOrStruct)
			for i := 0; i < num; i++ {
				field := sVal.Field(i)
				fieldKind := field.Kind()
				tField := sTp.Elem()
				if fieldKind != reflect.Struct && fieldKind != reflect.Map {
					tVal.SetMapIndex(reflect.ValueOf(tField.Name()), field)
				}
			}
		}
	}
}

// Keys get keys from map or struct.
//
// Notice: map keys maybe disorder.
func (obj Object) Keys(value any) []string {
	rv := reflect.ValueOf(value)
	var isPtr = false
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
		isPtr = true
	}

	var keys []string
	if rv.Kind() == reflect.Map {
		it := rv.MapRange()
		for it.Next() {
			keys = append(keys, it.Key().String())
		}
	} else if rv.Kind() == reflect.Struct {
		rt := reflect.TypeOf(value)
		if isPtr {
			rt = rt.Elem()
		}
		vn := rv.NumField()
		for i := 0; i < vn; i++ {
			field := rt.Field(i)
			// Get the tag JSON parameter first
			key := field.Tag.Get("json")
			if key == "" {
				key = field.Name
			}
			keys = append(keys, key)
		}
	}
	return keys
}

// StructToMap convert Struct field to by Map, support the Ptr
func StructToMap(value any, ignoreKeys ...string) map[string]any {
	rv := reflect.ValueOf(value)
	var rt reflect.Type
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
		rt = rv.Type()
	}
	if rv.Kind() == reflect.Struct {
		if rt == nil {
			rt = reflect.TypeOf(value)
		}
		vMap := map[string]any{}
		for i := 0; i < rv.NumField(); i++ {
			field := rv.Field(i)
			if !field.IsValid() {
				continue
			}
			// Notice: struct lower field also can be scan, and ignore func/ptr.
			if vKind := field.Kind(); vKind != reflect.Func && vKind != reflect.Ptr {
				name := rt.Field(i).Name
				//ignore keys
				if str.InQuei(name, ignoreKeys) > -1 {
					continue
				}
				vMap[name] = field.Interface()
			}
		}
		return vMap
	}
	return nil
}

// StructToMapLStyle convert Struct field to by Map and key is Lower style, key support `JSON.TAG`.
// Notice: reflect field num not contain inherit struct.
func StructToMapLStyle(value any, ignoreKeys ...string) map[string]any {
	rv := reflect.ValueOf(value)
	var rt reflect.Type
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
		rt = rv.Type()
	}
	if rv.Kind() == reflect.Struct {
		if rt == nil {
			rt = reflect.TypeOf(value)
		}
		vMap := map[string]any{}
		for i := 0; i < rv.NumField(); i++ {
			field := rv.Field(i)
			if !field.IsValid() {
				continue
			}
			if vKind := field.Kind(); vKind != reflect.Func && vKind != reflect.Ptr {
				tField := rt.Field(i)
				// Support JSON/TAG
				name := tField.Name
				//ignore keys
				if str.InQuei(name, ignoreKeys) > -1 {
					continue
				}
				if tagName, isOk := tField.Tag.Lookup("json"); isOk {
					if tagName != "-" && tagName != "" {
						name = tagName
					}
				} else {
					name = str.LowerStyle(name)
				}
				vMap[name] = field.Interface()
			}
		}
		return vMap
	}
	return nil
}

// ToMapLStyleIgnoreEmpty convert Struct field to by Map and key is Lower style and ignore empty.
// StructToMapViaJson is slower than StructToMapLStyle by Benchmark
func ToMapLStyleIgnoreEmpty(value any, ignoreKeys ...string) map[string]any {
	rv := reflect.ValueOf(value)
	var rt reflect.Type
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
		rt = rv.Type()
	}
	if rv.Kind() == reflect.Struct {
		if rt == nil {
			rt = reflect.TypeOf(value)
		}
		vMap := map[string]any{}
		for i := 0; i < rv.NumField(); i++ {
			field := rv.Field(i)
			if !field.IsValid() {
				continue
			}
			if vKind := field.Kind(); vKind != reflect.Func && vKind != reflect.Ptr {
				if !field.IsZero() {
					tField := rt.Field(i)
					name := tField.Name
					//ignore keys
					if str.InQuei(name, ignoreKeys) > -1 {
						continue
					}

					// Support JSON/TAG
					if tagName, isOk := tField.Tag.Lookup("json"); isOk {
						if tagName != "-" && tagName != "" {
							name = tagName
						}
					} else {
						name = str.LowerStyle(name)
					}
					vMap[name] = field.Interface()
				}
			}
		}
		return vMap
	}
	return nil
}

// StructToMapViaJson convert map via json Marshal/Unmarshal
// StructToMapViaJson is slower than StructToMapLStyle by Benchmark
func StructToMapViaJson(value any, ignoreKeys ...string) map[string]any {
	var newVal map[string]any
	marshal, err := json.Marshal(value)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(marshal, &newVal)
	if err != nil {
		return nil
	}
	if len(ignoreKeys) > 0 && newVal != nil {
		for key, _ := range newVal {
			//ignore keys
			if str.InQuei(key, ignoreKeys) > -1 {
				delete(newVal, key)
			}
		}
	}
	return newVal
}
