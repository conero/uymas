package util

import (
	"gitee.com/conero/uymas/v2/util/constraints"
	"gitee.com/conero/uymas/v2/util/rock"
)

// ExtractParam Extract indefinite parameters from functions and default code values
// Deprecated: please replace by `rock.ExtractParam`，will remove 1.4
func ExtractParam[T constraints.Equable](defValue T, args ...T) T {
	return rock.ExtractParam(defValue, args...)
}
