package rock

import "gitee.com/conero/uymas/util/constraints"

// ExtractParam Extract indefinite parameters from functions and default code values
func ExtractParam[T constraints.Equable](defValue T, args ...T) T {
	if len(args) > 0 {
		defValue = args[0]
	}
	return defValue
}
