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

func MustNoPanic(noPanic bool) {
	globalMustPanic = !noPanic
}
