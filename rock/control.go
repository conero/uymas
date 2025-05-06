package rock

// If returns trueVal 'if' condition is true, otherwise returns falseVal.
//
// It is used to implement equivalent ternary sign operations.
func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}
