package str

import (
	"gitee.com/conero/uymas/util/rock"
	"regexp"
	"strings"
)

// Calc String equality operator, which calculates the result of input string equality
// @todo should be completed in v1.3.x
type Calc struct {
	equality string
	result   float64
}

func NewCalc(equality string) *Calc {
	return &Calc{
		equality: equality,
	}
}

func (c *Calc) Count(args ...string) float64 {
	equality := rock.ExtractParam(c.equality, args...)
	c.equality = equality
	return c.result
}

// FloatSimple Floating-point number string beautification
func FloatSimple(fv string) string {
	potSig := "."
	split := strings.Split(fv, potSig)
	if len(split) == 2 {
		zero := split[1]
		// 全0
		isMatched, _ := regexp.MatchString(`^0+$`, zero)
		if isMatched {
			return strings.ReplaceAll(fv, "."+zero, "")
		}

		zeroReg := regexp.MustCompile(`0+$`)
		if zeroReg.MatchString(zero) {
			zeroList := zeroReg.FindAllString(zero, -1)
			if len(zeroList) > 0 {
				zero = zero[:len(zero)-len(zeroList[0])]
				return split[0] + "." + zero
			}
		}

	}

	return fv
}
