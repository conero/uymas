package rock

import "fmt"

// If returns trueVal 'if' condition is true, otherwise returns falseVal.
//
// It is used to implement equivalent ternary sign operations.
func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

// Must dimensionality reduction is performed on binary parameters
func Must[T any](value T, err error) T {
	if !globalMustPanic {
		return value
	}
	if err != nil {
		panic(fmt.Sprintf("Must: the value=%#v show must by set, and no error", value))
	}
	return value
}

// MustFunc it is used to implement dimensionality reduction for binary arrays with errors and supports callback functions.
//
// The handlerErrFn that if the callback function returns false, it will return vacancy.
func MustFunc[T any](handlerErrFn func(error) bool) func(T, error) T {
	return func(value T, err error) T {
		if handlerErrFn == nil {
			return value
		}
		if err != nil {
			if !handlerErrFn(err) {
				var zeroValue T
				return zeroValue
			}
		}
		return value
	}
}

// MustNoPanic to set the function Must don't to panic if exist error
func MustNoPanic(noPanic bool) {
	globalMustPanic = !noPanic
}
