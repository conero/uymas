package util

import "gitee.com/conero/uymas/v2/rock/constraints"

// MapAssign Merge multiple map parameters, where the same key value is the forward overwrite value.
//
// And never return nil.
func MapAssign[K constraints.KeyIterable, V constraints.ValueIterable](source map[K]V, more ...map[K]V) map[K]V {
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
