// Package util implements other tool more, like type cover, type value check.
package util

import (
	"math"
	"reflect"
	"time"
)

// @Date：   2018/10/30 0030 13:26
// @Author:  Joshua Conero
// @Name:    工具栏

// InQue exist value exist in array, if not exists will return -1
func InQue(val any, que []any) int {
	idx := -1
	for i, v := range que {
		if v == val {
			idx = i
			break
		}
	}
	return idx
}

// InQueAny Check keys if exist in Array Or Slice.
func InQueAny(que any, keys ...any) int {
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

// SpendTimeDiff Get the program spend time for any format.
func SpendTimeDiff() func() time.Duration {
	now := time.Now()
	return func() time.Duration {
		return time.Since(now)
	}
}

// Round String method processing float equal length data specified digits
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
	num = float64(int(num)) + float64(n2t)/math.Pow10(b)
	return num
}

// DecT36 Data conversion
func DecT36(num int) string {
	return (&Decimal{num}).T36()
}

// DecT62 Data conversion
func DecT62(num int) string {
	return (&Decimal{num}).T62()
}

// NullDefault null value handler to default.
func NullDefault(value, def any) any {
	if ValueNull(value) {
		return def
	}
	return value
}

// ValueNull to find if is null
func ValueNull(value any) bool {
	if nil == value {
		return true
	}
	v := reflect.ValueOf(value)
	return v.IsZero()
}
