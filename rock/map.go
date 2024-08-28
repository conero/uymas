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
	for k, _ := range data {
		keys = append(keys, fmt.Sprintf("%v", k))
	}
	return keys
}
