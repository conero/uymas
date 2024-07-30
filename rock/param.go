package rock

import "gitee.com/conero/uymas/v2/rock/constraints"

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
