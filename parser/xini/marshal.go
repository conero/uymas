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
	// 先处理非section部门，由于map无的序的
	for _, key := range rv.MapKeys() {
		keyStr := fmt.Sprintf("%v", key)
		isSec := strings.Index(keyStr, baseSecRegPref) == 0
		if isSec {
			continue
		}
		value := rv.MapIndex(key)
		valueStr, isMt := marshalToString(value, "")
		if isMt {
			valueStr = fmt.Sprintf("{\n%s\n}\n", valueStr)
			c.buff.WriteString(fmt.Sprintf("\n%s = %s\n", keyStr, valueStr))
		} else {
			c.buff.WriteString(fmt.Sprintf("%s = %s\n", keyStr, valueStr))
		}
	}

	// Section, 处理节部分
	for _, key := range rv.MapKeys() {
		keyStr := fmt.Sprintf("%v", key)
		isSec := strings.Index(keyStr, baseSecRegPref) == 0
		if !isSec {
			continue
		}

		keyStr = keyStr[len(baseSecRegPref):]
		value := rv.MapIndex(key)
		valueStr, isMt := marshalToString(value, "")
		if isMt {
			c.buff.WriteString(fmt.Sprintf("\n[%s]\n%s\n", keyStr, valueStr))
		} else {
			c.buff.WriteString(fmt.Sprintf("%s = %s\n", keyStr, valueStr))
		}
	}
}

func (c *Encoder) putStruct(v any) {
	rv := reflect.ValueOf(v)
	rt := reflect.TypeOf(v)
	for i := 0; i < rv.NumField(); i++ {
		sf := rt.Field(i)
		fieldName := sf.Name
		value := rv.FieldByName(fieldName)
		valueStr, isMt := marshalToString(value, "")
		if isMt {
			valueStr = fmt.Sprintf("{\n%s\n}", valueStr)
			c.buff.WriteString(fmt.Sprintf("\n%v = %s\n", fieldName, valueStr))
		} else {
			c.buff.WriteString(fmt.Sprintf("%v = %s\n", fieldName, valueStr))
		}
	}
}

// 返回渲染的内容，以及返回内容是否多行
func marshalToString(v reflect.Value, parentKey string) (string, bool) {
	if !v.IsValid() {
		return "", false
	}

	if v.Kind() == reflect.Interface {
		v = reflect.ValueOf(v.Interface())
	}
	switch v.Kind() {
	case reflect.String:
		return fmt.Sprintf(`"%s"`, v.String()), false
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf(`%v`, v), false
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf(`%v`, v), false
	case reflect.Bool:
		return fmt.Sprintf(`%v`, v), false
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf(`%v`, v), false
	case reflect.Slice, reflect.Array:
		vl := v.Len()
		var arr []string
		for i := 0; i < vl && vl > 1; i++ {
			el := v.Index(i)
			valueStr, mtlLn := marshalToString(el, "")
			if mtlLn {
				valueStr = "{ " + valueStr + " }"
			}
			arr = append(arr, valueStr)
		}
		return strings.Join(arr, ","), false
	case reflect.Map: // support section
		var queue []string
		for _, key := range v.MapKeys() {
			value := v.MapIndex(key)
			valueStr, isMt := marshalToString(value, "")
			var mapLn string
			if parentKey != "" {
				if isMt {
					mapLn = fmt.Sprintf("\n%s.%s = {\n%v\n}", parentKey, key, valueStr)
				} else {
					mapLn = fmt.Sprintf("%s.%s = %v", parentKey, key, valueStr)
				}

			} else {
				if isMt {
					mapLn = fmt.Sprintf("\n%s = {\n%v\n}", key, valueStr)
				} else {
					mapLn = fmt.Sprintf("%s = %v", key, valueStr)
				}
			}
			queue = append(queue, mapLn)
		}
		startStr := ""
		if len(queue) > 0 {
			startStr = "    "
		}
		return startStr + strings.Join(queue, "\n    "), true
	}
	return fmt.Sprintf(`%v`, v), false
}
