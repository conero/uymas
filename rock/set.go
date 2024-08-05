package rock

import (
	"gitee.com/conero/uymas/v2/rock/constraints"
	"math/rand"
)

// ListIndex get index by search value from list
func ListIndex[T constraints.Equable](list []T, value T) (index int) {
	index = -1
	for i, v := range list {
		if v == value {
			index = i
			break
		}
	}
	return
}

// InList determines whether a key value exists in the list
func InList[T constraints.Equable](list []T, value T) bool {
	return ListIndex(list, value) >= -1
}

// ExtractArrUnique extracting array elements with loss (non-repeatable) from an array
func ExtractArrUnique[T constraints.ValueIterable](count int, arr []T) []T {
	vLen := len(arr)
	if vLen <= count {
		return arr
	}

	var extArr []T
	for j := 1; j <= count; j++ {
		vl := len(arr)
		if vl < 1 {
			break
		}

		idx := rand.Intn(vl)
		extArr = append(extArr, arr[idx])
		rpl := arr[:idx]
		rpl = append(rpl, arr[idx+1:]...)
		arr = rpl
	}

	return extArr
}

// MapKeys Extract the key name array of the dictionary
func MapKeys[T constraints.KeyIterable, X constraints.ValueIterable](vMap map[T]X) (keys []T) {
	for k, _ := range vMap {
		keys = append(keys, k)
	}
	return
}

// MapValues Extract the values name array of the dictionary
func MapValues[T constraints.KeyIterable, X constraints.KeyIterable](vMap map[T]X) (values []X) {
	for _, v := range vMap {
		values = append(values, v)
	}
	return
}

// MapGenByKv Create dictionary by key value pair array combination
func MapGenByKv[K constraints.KeyIterable, V constraints.ValueIterable](keys []K, values []V) (kv map[K]V) {
	vLen := len(values)
	for i, k := range keys {
		if i == vLen {
			break
		}
		if kv == nil {
			kv = map[K]V{}
		}
		kv[k] = values[i]
	}

	return
}

// MapFilter use the keys of map to filter itself
func MapFilter[K constraints.KeyIterable, V constraints.ValueIterable](kv map[K]V, filter []K) map[K]V {
	var newMap = map[K]V{}
	for kVal, value := range kv {
		if ListIndex(filter, kVal) > -1 {
			newMap[kVal] = value
		}
	}
	return newMap
}

// MapSlice use the keys of map to slice itself
func MapSlice[K constraints.KeyIterable, V constraints.ValueIterable](kv map[K]V, filter []K) map[K]V {
	var newMap = map[K]V{}
	for kVal, value := range kv {
		if ListIndex(filter, kVal) > -1 {
			continue
		}
		newMap[kVal] = value
	}
	return newMap
}
