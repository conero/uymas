package xini

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

// Marshal returns the ini encoding of v.
func Marshal(v any) ([]byte, error) {
	enc := NewEncoder()
	enc.Put(v)
	return enc.buff.Bytes(), nil
}

type Encoder struct {
	buff *bytes.Buffer
}

func NewEncoder() *Encoder {
	return &Encoder{
		buff: bytes.NewBufferString(""),
	}
}

func (c *Encoder) Put(v any) {
	if v == nil {
		return
	}
	rv := reflect.ValueOf(v)
	kind := rv.Kind()
	// the value by pointer point data.
	if kind == reflect.Pointer {
		rv = rv.Elem()
		kind = rv.Kind()
	}

	if kind == reflect.Map {
		c.putMap(v)
	} else if kind == reflect.Struct {
		c.putStruct(v)
	}
}

func (c *Encoder) putMap(v any) {
	rv := reflect.ValueOf(v)
	for _, key := range rv.MapKeys() {
		value := rv.MapIndex(key)
		c.buff.WriteString(fmt.Sprintf("%v = %s\n", key, valueToString(value)))
	}
}

func (c *Encoder) putStruct(v any) {
	rv := reflect.ValueOf(v)
	rt := reflect.TypeOf(v)
	for i := 0; i < rv.NumField(); i++ {
		sf := rt.Field(i)
		fieldName := sf.Name
		value := rv.FieldByName(fieldName)
		c.buff.WriteString(fmt.Sprintf("%v = %s\n", fieldName, valueToString(value)))
	}
}

func valueToString(v reflect.Value) string {
	if v.Kind() == reflect.Interface {
		v = reflect.ValueOf(v.Interface())
	}
	//fmt.Printf("%v => value: %v\n", v, v.Kind())
	switch v.Kind() {
	case reflect.String:
		return fmt.Sprintf(`"%s"`, v.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf(`%v`, v)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf(`%v`, v)
	case reflect.Bool:
		return fmt.Sprintf(`%v`, v)
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf(`%v`, v)
	case reflect.Slice, reflect.Array:
		vl := v.Len()
		var arr []string
		for i := 0; i < vl && vl > 1; i++ {
			el := v.Index(i)
			arr = append(arr, valueToString(el))
		}
		return strings.Join(arr, ",")
	}
	return fmt.Sprintf(`%v`, v)
}
