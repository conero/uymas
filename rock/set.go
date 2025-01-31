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
	return ListIndex(list, value) > -1
}

// ListSubset determine whether the given subarray is accurate
func ListSubset[T constraints.Equable](list []T, subset []T) bool {
	for _, sub := range subset {
		if !InList(list, sub) {
			return false
		}
	}
	return true
}

// ListEq Determines whether an array is equal
func ListEq[T constraints.Equable](list []T, compare []T) bool {
	if len(list) != len(compare) {
		return false
	}
	for _, sub := range compare {
		if !InList(list, sub) {
			return false
		}
	}
	return true
}

// ListRemove removes the specified element from the list
func ListRemove[T constraints.Equable](list []T, removes ...T) []T {
	var newList []T
	for _, v := range list {
		if InList(removes, v) {
			continue
		}
		newList = append(newList, v)
	}
	return newList
}

// ListGetOr Read element values through index or use default substitution (when not present)
func ListGetOr[T any](list []T, index int, def T) T {
	if index < 0 {
		return def
	}
	if len(list) > index {
		return list[index]
	}
	return def
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

// ListAny data slice convert to any slice
func ListAny[T constraints.KeyIterable](vList []T) []any {
	var anyList []any
	for _, v := range vList {
		anyList = append(anyList, v)
	}
	return anyList
}
