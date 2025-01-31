package rock

import (
	"fmt"
	"gitee.com/conero/uymas/v2/rock/constraints"
)

// MapAssign Merge multiple map parameters, where the same key value is the forward overwrite value.
//
// And never return nil.
func MapAssign[K constraints.KeyIterable, V any](source map[K]V, more ...map[K]V) map[K]V {
	if source == nil {
		source = map[K]V{}
	}
	for _, mr := range more {
		for k, v := range mr {
			source[k] = v
		}
	}
	return source
}

func MapKeysString[K constraints.KeyIterable, V any](data map[K]V) []string {
	var keys []string
	for k := range data {
		keys = append(keys, fmt.Sprintf("%v", k))
	}
	return keys
}

// MapKeys Extract the key name array of the dictionary
func MapKeys[T constraints.KeyIterable, X constraints.ValueIterable](vMap map[T]X) (keys []T) {
	for k := range vMap {
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
