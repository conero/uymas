package util

import (
	"gitee.com/conero/uymas/util/constraints"
	"gitee.com/conero/uymas/util/rock"
)

// ListIndex get index by search value from list
//
// Deprecated: only as a copy, has been moved to package util/rock, later deleted
func ListIndex[T constraints.Equable](list []T, value T) (index int) {
	return rock.ListIndex(list, value)
}

// ListNoRepeat filter duplicate elements in list
//
// Deprecated: only as a copy, has been moved to package util/rock, later deleted
func ListNoRepeat[T constraints.Equable](list []T) []T {
	return rock.ListNoRepeat(list)
}

// ExtractArrUnique extracting array elements with loss (non-repeatable) from an array
//
// Deprecated: only as a copy, has been moved to package util/rock, later deleted
func ExtractArrUnique[T constraints.ValueIterable](count int, arr []T) []T {
	return rock.ExtractArrUnique(count, arr)
}

// MapKeys Extract the key name array of the dictionary
//
// Deprecated: only as a copy, has been moved to package util/rock, later deleted
func MapKeys[T constraints.KeyIterable, X constraints.ValueIterable](vMap map[T]X) (keys []T) {
	return rock.MapKeys(vMap)
}

// MapValues Extract the values name array of the dictionary
//
// Deprecated: only as a copy, has been moved to package util/rock, later deleted
func MapValues[T constraints.KeyIterable, X constraints.KeyIterable](vMap map[T]X) (values []X) {
	return rock.MapValues(vMap)
}

// MapGenByKv Create dictionary by key value pair array combination
//
// Deprecated: only as a copy, has been moved to package util/rock, later deleted
func MapGenByKv[K constraints.KeyIterable, V constraints.ValueIterable](keys []K, values []V) (kv map[K]V) {
	return rock.MapGenByKv(keys, values)
}

// MapFilter use the keys of map to filter itself
//
// Deprecated: only as a copy, has been moved to package util/rock, later deleted
func MapFilter[K constraints.KeyIterable, V constraints.ValueIterable](kv map[K]V, filter []K) map[K]V {
	return rock.MapFilter(kv, filter)
}

// MapSlice use the keys of map to slice itself
//
// Deprecated: only as a copy, has been moved to package util/rock, later deleted
func MapSlice[K constraints.KeyIterable, V constraints.ValueIterable](kv map[K]V, filter []K) map[K]V {
	return rock.MapSlice(kv, filter)
}
