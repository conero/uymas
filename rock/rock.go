// Package rock generic based common package processing
package rock

import "gitee.com/conero/uymas/v2/rock/constraints"

// ListNoRepeat filter duplicate elements in list
func ListNoRepeat[T constraints.Equable](list []T) []T {
	var noRepeat []T
	var tmpMap = map[T]bool{}

	for _, v := range list {
		if _, exist := tmpMap[v]; exist {
			continue
		}

		noRepeat = append(noRepeat, v)
		tmpMap[v] = true
	}
	return noRepeat
}
