package util

import (
	"gitee.com/conero/uymas/util/constraints"
	"gitee.com/conero/uymas/util/rock"
)

// ExtractParam Extract indefinite parameters from functions and default code values
// Deprecated: please replace by `rock.ExtractParam`，will remove in future
func ExtractParam[T constraints.Equable](defValue T, args ...T) T {
	return rock.ExtractParam(defValue, args...)
}
