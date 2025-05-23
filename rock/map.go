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
func MapKeys[T constraints.KeyIterable, X any](vMap map[T]X) (keys []T) {
	for k := range vMap {
		keys = append(keys, k)
	}
	return
}

// MapValues Extract the values name array of the dictionary
func MapValues[T constraints.KeyIterable, X any](vMap map[T]X) (values []X) {
	for _, v := range vMap {
		values = append(values, v)
	}
	return
}

// MapValuesRely Produce a key value array in the order of the provided key name table,
// and provide a callback function when it does not exist.
//
// Applied to the problem of unordered Map key values
func MapValuesRely[T constraints.KeyIterable, X any](vMap map[T]X, keys []T, noExistFn func(key T) X) (values []X) {
	for _, key := range keys {
		if v, ok := vMap[key]; ok {
			values = append(values, v)
		} else if noExistFn != nil {
			values = append(values, noExistFn(key))
		} else {
			var defVal X
			values = append(values, defVal)
		}
	}
	return
}

// MapGenByKv Create dictionary by key value pair array combination
func MapGenByKv[K constraints.KeyIterable, V any](keys []K, values []V) (kv map[K]V) {
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

// MapGenFlat Create dictionary by key value pair array combination and fixed value
func MapGenFlat[K constraints.KeyIterable, V any](keys []K, value V) (kv map[K]V) {
	for _, k := range keys {
		if kv == nil {
			kv = map[K]V{}
		}
		kv[k] = value
	}

	return
}

// MapGenFlatFn Create dictionary by key value pair array combination and fixed value from callback
func MapGenFlatFn[K constraints.KeyIterable, V any](keys []K, defFn func(K) V) (kv map[K]V) {
	var initValue V
	for _, k := range keys {
		if kv == nil {
			kv = map[K]V{}
		}
		if defFn == nil {
			kv[k] = initValue
			continue
		}
		kv[k] = defFn(k)
	}

	return
}

// MapFilter use the keys of map to filter itself
func MapFilter[K constraints.KeyIterable, V any](kv map[K]V, filter []K) map[K]V {
	var newMap = map[K]V{}
	for kVal, value := range kv {
		if ListIndex(filter, kVal) > -1 {
			newMap[kVal] = value
		}
	}
	return newMap
}

// MapSlice use the keys of map to slice itself
func MapSlice[K constraints.KeyIterable, V any](kv map[K]V, filter []K) map[K]V {
	var newMap = map[K]V{}
	for kVal, value := range kv {
		if ListIndex(filter, kVal) > -1 {
			continue
		}
		newMap[kVal] = value
	}
	return newMap
}
