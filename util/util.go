//Package util implements other tool more, like type cover, type value check.
package util

import (
	"fmt"
	"github.com/conero/uymas/str"
	"math"
	"reflect"
	"time"
)

// @Date：   2018/10/30 0030 13:26
// @Author:  Joshua Conero
// @Name:    工具栏

// 数组中是否存在
// 不存在返回 -1
func InQue(val interface{}, que []interface{}) int {
	idx := -1
	if que != nil {
		for i, v := range que {
			if v == val {
				idx = i
				break
			}
		}
	}
	return idx
}

//Check keys if exist in Array Or Slice.
func InQueAny(que interface{}, keys ...interface{}) int {
	idx := -1

	vt := reflect.ValueOf(que)
	//Only Array And Slice.
	if vt.Kind() == reflect.Array || vt.Kind() == reflect.Slice {
		vLen := vt.Len()
		for i := 0; i < vLen; i++ {
			value := vt.Index(i)
			for j := 0; j < len(keys); j++ {
				vsKey := keys[j]
				if reflect.DeepEqual(value.Interface(), vsKey) {
					idx = i
					break
				}
			}
			if idx != -1 {
				break
			}
		}
	}

	return idx
}

//返回秒用于计算程序用时,参数为0时返回当前的毫秒，否则返回计算后的秒差
func Sec(start float64) float64 {
	t := time.Now()
	ns := float64(t.Nanosecond())
	ms := ns / math.Pow10(6) //1ms = 10^6ns
	if start == 0 {
		return ms
	}
	ds := (ms - start) / math.Pow10(3)
	ds = Round(ds, 5)
	return ds
}

// 返回运算秒的秒数
func SecCall() func() float64 {
	start := Sec(0)
	return func() float64 {
		return Sec(start)
	}
}

// 返回字符串式的运行毫秒
func SecCallStr() func() string {
	start := time.Now()
	return func() string {
		return fmt.Sprintf("%s", time.Since(start))
	}
}

//字符串方法处理float等长数据 规定位数
func Round(num float64, b int) float64 {
	if b == 0 {
		return float64(int(num))
	}
	n2t := int(num * math.Pow10(b))    //num转换数
	base := int(num * math.Pow10(b+1)) //四舍五入的最后一位数
	base = int(math.Abs(float64(base - n2t*10)))
	if base > 5 {
		n2t += 1
	}
	num = float64(int(num)) + float64(n2t)/float64(math.Pow10(b))
	return num
}

// 数据进制转换
func DecT36(num int) string {
	return (&Decimal{num}).T36()
}

// 数据进制转换
func DecT62(num int) string {
	return (&Decimal{num}).T62()
}

//null value handler to default.
func NullDefault(value, def interface{}) interface{} {
	if ValueNull(value) {
		return def
	}
	return value
}

//to find if is null
func ValueNull(value interface{}) bool {
	if nil == value {
		return true
	}
	v := reflect.ValueOf(value)
	return v.IsZero()
}

// convert Struct field to by Map
func StructToMap(value interface{}) map[string]interface{} {
	rv := reflect.ValueOf(value)
	if rv.Kind() == reflect.Struct {
		rt := reflect.TypeOf(value)
		vMap := map[string]interface{}{}
		for i := 0; i < rv.NumField(); i++ {
			field := rv.Field(i)
			if field.Kind() != reflect.Func {
				name := rt.Field(i).Name
				vMap[name] = field.Interface()
			}
		}
		return vMap
	}
	return nil
}

// convert Struct field to by Map and key is Lower style.
func StructToMapLStyle(value interface{}) map[string]interface{} {
	rv := reflect.ValueOf(value)
	if rv.Kind() == reflect.Struct {
		rt := reflect.TypeOf(value)
		vMap := map[string]interface{}{}
		for i := 0; i < rv.NumField(); i++ {
			field := rv.Field(i)
			if field.Kind() != reflect.Func {
				name := rt.Field(i).Name
				vMap[str.LowerStyle(name)] = field.Interface()
			}
		}
		return vMap
	}
	return nil
}
