package util

import (
	"encoding/json"
	"fmt"
	"gitee.com/conero/uymas/v2/str"
	"reflect"
	"strconv"
	"strings"
)

var (
	// Obj the object instance for call Object method directly
	Obj Object
)

type Object struct {
}

// Assign Base of `reflect` to come true like javascript `Object.Assign`, target should be pointer best.
// It can be Multiple, only for `reflect.Map`. And support nested struct.
// @todo 将使用 reflect 与 未使用reflect从包去区分开
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

	// Automatically support value conversion when it is not struct or map.
	if tRefv.Kind() != reflect.Struct {
		obj.AssignCovert(target, source)
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
				obj.nestStructAssign(tField, sField)
			} else {
				tField.Set(sField)
			}
		}
	}

	return m
}

// Nested struct substructures alignment and assignment
func (obj Object) nestStructAssign(dst, src reflect.Value) {
	if dst.Kind() != src.Kind() || dst.Kind() != reflect.Struct {
		return
	}

	dstType := dst.Type()
	for i := 0; i < dst.NumField(); i++ {
		nestField := dst.Field(i)
		dFieldTy := dstType.Field(i)

		sField := src.FieldByName(dFieldTy.Name)
		if !sField.IsValid() || sField.IsZero() || sField.Kind() != nestField.Kind() {
			continue
		}

		if sField.Kind() == reflect.Struct {
			obj.nestStructAssign(nestField, sField)
		} else {
			nestField.Set(sField)
		}
	}
}

// AssignCovert Simple type automatic coverage, supporting cross type. So do not try to cover complex type.
// number covert
// any -> string
// string -> number
func (obj Object) AssignCovert(target any, source any) any {
	if source == nil {
		return target
	}

	tvf := reflect.ValueOf(target)
	// AssignCovert: The target parameter should provide a string
	if tvf.Kind() != reflect.Pointer {
		return target
	}
	tvf = tvf.Elem()

	svf := reflect.ValueOf(source)
	if svf.IsZero() {
		return target
	}

	ttf := reflect.TypeOf(tvf.Interface()) // The type of actual execution target
	tvfKind := ttf.Kind()

	if tvfKind == svf.Kind() { // same parameter type
		tvf.Set(svf)
	} else if tvfKind == reflect.String { // all types can covert into strings
		tvf.Set(reflect.ValueOf(fmt.Sprintf("%v", source)))
	} else if tvfKind == reflect.Bool { // Non null values are valid, can be true
		tvf.Set(reflect.ValueOf(true))
		//} else if obj.covertSameNumber(tvf, svf) {
		//	fmt.Printf("covertSameNumber/isOK, tvf: %v\n", tvf)
	} else if obj.stringCoverNumber(tvf, svf) {
	} else if svf.CanConvert(ttf) {
		tvf.Set(svf.Convert(ttf))
	}

	return target
}

// string type covert into a number
func (obj Object) stringCoverNumber(dst, src reflect.Value) bool {
	if src.Kind() != reflect.String {
		return false
	}
	vStr := src.String()
	dstKind := reflect.TypeOf(dst.Interface()).Kind()
	// init
	getIntFn := func() int64 {
		i64, err := strconv.ParseInt(vStr, 10, 64)
		if err == nil {
			return i64
		}

		u64, err := strconv.ParseUint(vStr, 10, 64)
		if err == nil {
			return int64(u64)
		}

		f64, err := strconv.ParseFloat(vStr, 64)
		if err == nil {
			return int64(f64)
		}
		return 0
	}
	// float
	getFloatFn := func() float64 {
		f64, err := strconv.ParseFloat(vStr, 64)
		if err == nil {
			return f64
		}

		i64, err := strconv.ParseInt(vStr, 10, 64)
		if err == nil {
			return float64(i64)
		}

		u64, err := strconv.ParseUint(vStr, 10, 64)
		if err == nil {
			return float64(u64)
		}

		return 0
	}
	// uint
	getUintFn := func() uint64 {
		u64, err := strconv.ParseUint(vStr, 10, 64)
		if err == nil {
			return u64
		}

		i64, err := strconv.ParseInt(vStr, 10, 64)
		if err == nil {
			return uint64(i64)
		}

		f64, err := strconv.ParseFloat(vStr, 64)
		if err == nil {
			return uint64(f64)
		}

		return 0
	}
	isOk := false
	switch dstKind {
	case reflect.Int:
		dst.Set(reflect.ValueOf(int(getIntFn())))
		isOk = true
	case reflect.Int8:
		dst.Set(reflect.ValueOf(int8(getIntFn())))
		isOk = true
	case reflect.Int16:
		dst.Set(reflect.ValueOf(int16(getIntFn())))
		isOk = true
	case reflect.Int32:
		dst.Set(reflect.ValueOf(int32(getIntFn())))
		isOk = true
	case reflect.Int64:
		dst.Set(reflect.ValueOf(getIntFn()))
		isOk = true
	case reflect.Uint:
		dst.Set(reflect.ValueOf(uint(getUintFn())))
		isOk = true
	case reflect.Uint8:
		dst.Set(reflect.ValueOf(uint8(getUintFn())))
		isOk = true
	case reflect.Uint16:
		dst.Set(reflect.ValueOf(uint16(getUintFn())))
		isOk = true
	case reflect.Uint32:
		dst.Set(reflect.ValueOf(uint32(getUintFn())))
		isOk = true
	case reflect.Uint64:
		dst.Set(reflect.ValueOf(getUintFn()))
		isOk = true
	case reflect.Float64:
		dst.Set(reflect.ValueOf(getFloatFn()))
		isOk = true
	case reflect.Float32:
		dst.Set(reflect.ValueOf(float32(getFloatFn())))
		isOk = true
	}
	return isOk
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
	if rv.Kind() != reflect.Struct {
		return nil
	}
	if rt == nil {
		rt = reflect.TypeOf(value)
	}
	vMap := map[string]any{}
	var toSetMapValueFn func(value reflect.Value)
	toSetMapValueFn = func(value reflect.Value) {
		if value.Kind() != reflect.Struct {
			return
		}
		valueTp := value.Type()
		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i)
			if !field.IsValid() {
				continue
			}
			// Notice: struct lower field also can be scanned, and ignore func/ptr.
			if vKind := field.Kind(); vKind != reflect.Func && vKind != reflect.Ptr {
				structField := valueTp.Field(i)
				name := structField.Name
				//ignore keys
				if str.InQuei(name, ignoreKeys) > -1 {
					continue
				}
				// determine whether it is a combination inheritance type
				if structField.Anonymous {
					toSetMapValueFn(field)
					continue
				}
				if !field.CanInterface() {
					continue
				}
				vMap[name] = field.Interface()
			}
		}
	}

	toSetMapValueFn(rv)
	return vMap
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
					if tagName == "-" {
						continue
					}
					if tagName != "" {
						name = tagName
						idx := strings.Index(name, ",")
						if idx > -1 {
							name = name[:idx]
						}
					}
				} else {
					name = str.Str(name).LowerStyle()
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
						name = str.Str(name).LowerStyle()
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

// MapToStruct use the map value to set structPtr
func MapToStruct(mapValue any, structPtr any, supportTags []string) error {
	vm := reflect.ValueOf(mapValue)
	if vm.Kind() != reflect.Map {
		return fmt.Errorf("mapValue is not belong to map type")
	}

	sp := reflect.ValueOf(structPtr)
	if sp.Kind() != reflect.Ptr {
		return fmt.Errorf("structPtr is not belong to ptr type")
	}

	sp = sp.Elem()
	if sp.Kind() != reflect.Struct {
		return fmt.Errorf("structPtr is not belong to ptr type of struct")
	}

	mapValKv := map[string]reflect.Value{}
	for _, mKey := range vm.MapKeys() {
		mVal := vm.MapIndex(mKey)
		keyName := mKey.String()
		mapValKv[keyName] = mVal
	}
	spt := sp.Type()
	for i := 0; i < sp.NumField(); i++ {
		field := sp.Field(i)
		sf := spt.Field(i)
		keyName := sf.Name

		// from file name
		if src, exist := mapValKv[keyName]; exist {
			TryAssignValue(src, field)
			continue
		}

		// from tag
		isContinue := false
		for _, tagName := range supportTags {
			// from json tag name
			jsonName, exist := sf.Tag.Lookup(tagName)
			if exist {
				if src, exist := mapValKv[jsonName]; exist {
					TryAssignValue(src, field)
					isContinue = true
					break
				}
			}
		}
		if isContinue {
			continue
		}
		// lc-style
		lcKey := str.Str(sf.Name).Lcfirst()
		if lcKey != sf.Name {
			if src, exist := mapValKv[lcKey]; exist {
				TryAssignValue(src, field)
				continue
			}
		}
	}

	return nil
}

// MapToStructViaJson use the map value to set structPtr
func MapToStructViaJson(mapValue any, structPtr any) error {
	return MapToStruct(mapValue, structPtr, []string{"json"})
}

// TryAssignValue try assign value to another value.
func TryAssignValue(src reflect.Value, tgt reflect.Value) bool {
	if !tgt.CanSet() {
		return false
	}

	// interface
	if src.CanInterface() {
		src = reflect.ValueOf(src.Interface())
	}

	tTy := tgt.Type()
	sTy := src.Type()
	// assign
	if sTy.AssignableTo(tTy) {
		tgt.Set(src)
		return true
	}

	// convert
	if sTy.ConvertibleTo(tTy) {
		tgt.Set(src.Convert(tTy))
		return true
	}

	// any only can be int, bool, float
	sKind := src.Kind()
	tKind := tgt.Kind()
	if tKind == reflect.String { // any -> string
		tgt.Set(reflect.ValueOf(fmt.Sprintf("%v", src)))
	} else if sKind == reflect.String { // string -> any.
		stringVal := src.String()
		if tgt.CanInt() { // string -> int
			i, err := strconv.ParseInt(stringVal, 10, 64)
			if err != nil {
				return false
			}

			tgt.SetInt(i)
			return true
		} else if tgt.CanUint() { // string -> int
			u, err := strconv.ParseUint(stringVal, 10, 64)
			if err != nil {
				return false
			}
			tgt.SetUint(u)
			return true
		} else if tgt.CanFloat() { // string -> int
			f, err := strconv.ParseFloat(stringVal, 64)
			if err != nil {
				return false
			}
			tgt.SetFloat(f)
			return true
		} else if tKind == reflect.Bool {
			if strings.ToLower(stringVal) == "true" {
				tgt.SetBool(true)
				return true
			} else if strings.ToLower(stringVal) == "false" {
				tgt.SetBool(false)
				return true
			}
			return false
		}
	} else if tKind == reflect.Bool { // any -> bool
		tgt.SetBool(!src.IsZero())
		return true
	}
	return false
}
