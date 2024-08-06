package rock

// Param Extract indefinite parameters from functions and default code values
func Param[T any](defValue T, args ...T) T {
	if len(args) > 0 {
		defValue = args[0]
	}
	return defValue
}

// ParamFunc Implementing parameter extraction through custom callback functions
func ParamFunc[T any](defFunc func() T, args ...T) T {
	var defValue T
	if len(args) > 0 && defFunc != nil {
		defValue = defFunc()
	}
	return defValue
}
