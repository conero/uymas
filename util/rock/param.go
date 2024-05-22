package rock

import "gitee.com/conero/uymas/util/constraints"

// ExtractParam Extract indefinite parameters from functions and default code values
func ExtractParam[T constraints.Equable](defValue T, args ...T) T {
	if len(args) > 0 {
		defValue = args[0]
	}
	return defValue
}

// ExtractParamFunc Implementing parameter extraction through custom callback functions
func ExtractParamFunc[T constraints.Equable](defFunc func() T, args ...T) T {
	var defValue T
	if len(args) > 0 && defFunc != nil {
		defValue = defFunc()
	}
	return defValue
}

// ExtractParamIndex Extract indefinite parameters from functions and default code values and point index
//
// index => [1, ..]
func ExtractParamIndex[T constraints.Equable](defValue T, index int, args ...T) T {
	vLen := len(args)
	if vLen >= index {
		defValue = args[index-1]
	}
	return defValue
}
