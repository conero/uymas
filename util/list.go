package util

import "gitee.com/conero/uymas/util/constraints"

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
