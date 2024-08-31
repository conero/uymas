// Package number handler like unit cover.
package number

import (
	"fmt"
	"math"
	"strconv"
)

// @Date：   2018/12/20 0020 16:16
// @Author:  Joshua Conero
// @Name:    数字包

// SumQueue the sum value by any number type
// support type: int, int8, int16, int32, int64
// support type: uint, uint8, uint16, uint32, uint64
// support type: float32, float64
func SumQueue(data []any) any {
	var a any = nil
	for _, d := range data {
		switch d.(type) {
		case int:
			if a == nil {
				var tInt int = 0
				a = tInt
			}
			a = a.(int) + d.(int)
		case int8:
			if a == nil {
				var tInt int8 = 0
				a = tInt
			}
			a = a.(int8) + d.(int8)
		case int16:
			if a == nil {
				var tInt int16 = 0
				a = tInt
			}
			a = a.(int16) + d.(int16)
		case int32:
			if a == nil {
				var tInt int32 = 0
				a = tInt
			}
			a = a.(int32) + d.(int32)
		case int64:
			if a == nil {
				var tInt int64 = 0
				a = tInt
			}
			a = a.(int64) + d.(int64)
		case uint:
			if a == nil {
				var tInt uint = 0
				a = tInt
			}
			a = a.(uint) + d.(uint)
		case uint8:
			if a == nil {
				var tInt uint8 = 0
				a = tInt
			}
			a = a.(uint8) + d.(uint8)
		case uint32:
			if a == nil {
				var tInt uint32 = 0
				a = tInt
			}
			a = a.(uint32) + d.(uint32)
		case uint64:
			if a == nil {
				var tInt uint64 = 0
				a = tInt
			}
			a = a.(uint64) + d.(uint64)
		case float32:
			if a == nil {
				var tInt float32 = 0
				a = tInt
			}
			a = a.(float32) + d.(float32)
		case float64:
			if a == nil {
				var tInt float64 = 0
				a = tInt
			}
			a = a.(float64) + d.(float64)
		}
	}
	return a
}

// SumQInt the sum value from int array
func SumQInt(data []int) int {
	if data == nil {
		data = []int{}
	}
	var a int
	for _, n := range data {
		a += n
	}
	return a
}

func AnyInt64(v any) int64 {
	var i64 int64
	if v != nil {
		switch value := v.(type) {
		case int64:
			i64 = value
		case int32:
			i64 = int64(value)
		case int:
			i64 = int64(value)
		case int16:
			i64 = int64(value)
		case int8:
			i64 = int64(value)
		default:
			if i, er := strconv.ParseInt(fmt.Sprintf("%v", v), 10, 64); er == nil {
				i64 = i
			}
		}
	}
	return i64
}

// Factorial Supports factorial operations on natural numbers
// `n! = n*(n-1)*(n-2)*...*1`
func Factorial(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	var amass uint64 = 1
	for i := n; i > 0; i-- {
		amass *= i
	}
	return amass
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
