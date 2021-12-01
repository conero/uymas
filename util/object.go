package util

import (
	"reflect"
)

type Object struct {
}

// Assign @todo
//	Base of `reflect` to come true like javascript `Object.Assign`, target should be pointer best.
//	It can be Multiple, only for `reflect.Map`.
func (obj Object) Assign(target interface{}, source interface{}) interface{} {
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
				if tField.CanAddr() {
					//fmt.Printf("Nest->tField %#v\n", tField)
					//fmt.Printf("Nest->sField %#v\n", sField)
					//obj.Assign(tField.Addr(), sField)
					//obj.Assign(tField.Addr(), sField)
				}
			} else {
				tField.Set(sField)
			}
		}
	}

	return m
}

// AssignMap Assign Map/Struct to map
func (obj Object) AssignMap(targetMap interface{}, srcMapOrStruct interface{}) {
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
