package util

import (
	"gitee.com/conero/uymas/util/constraints"
	"gitee.com/conero/uymas/util/rock"
)

// MapAssign Merge multiple map parameters, where the same key value is the forward overwrite value.
//
// And never return nil.
//
// Deprecated: only as a copy, has been moved to package util/rock, later deleted
func MapAssign[K constraints.KeyIterable, V constraints.ValueIterable](source map[K]V, more ...map[K]V) map[K]V {
	return rock.MapAssign(source, more...)
}
