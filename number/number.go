package number

// @Date：   2018/12/20 0020 16:16
// @Author:  Joshua Conero
// @Name:    数字包

// 集合求和
// 支持格式： int, int8, int16, int32, int64
// 支持格式： uint, uint8, uint16, uint32, uint64
// 支持格式： float32, float64
func SumQueue(data []interface{}) interface{} {
	var a interface{} = nil
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

// 集合
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
