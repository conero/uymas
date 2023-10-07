package str

import (
	"fmt"
	"gitee.com/conero/uymas/util/rock"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// Calc String equality operator, which calculates the result of input string equality
// support: `**,^,*,/,+,-,%`
type Calc struct {
	equality  string
	handlerEq string
	result    float64
	// `()`
	regBracket   *regexp.Regexp
	regBracketSg *regexp.Regexp
	clearReg     *regexp.Regexp // `space clear`
	// `*/%`
	mulDivReg   *regexp.Regexp
	mulDivRegSg *regexp.Regexp
	// `+-`
	addSubReg   *regexp.Regexp
	addSubRegSg *regexp.Regexp
	// x**y 或 x^y
	powReg      *regexp.Regexp
	powRegSg    *regexp.Regexp
	Accuracy    int8
	accuracyStr string
	simpleReg   *regexp.Regexp
	simpleRegSg *regexp.Regexp
	expReg      *regexp.Regexp
}

func NewCalc(equality string) *Calc {
	return &Calc{
		equality: equality,
		Accuracy: 7,
	}
}

// Bracket decomposition, clear `()`
func (c *Calc) deBracket(eq string) string {
	if c.regBracket == nil {
		c.regBracket = regexp.MustCompile(`\([^()]+\)`)
		//@todo Verification and comparison required
		//c.regBracket = regexp.MustCompile(`[[:^alpha:]]*\([^()]+\)`)
	}

	bracket := c.regBracket.FindAllString(eq, -1)
	for _, brk := range bracket {
		rslt := c.operNonBrk(brk)
		eq = strings.ReplaceAll(eq, brk, rslt)
	}

	if c.regBracket.MatchString(eq) {
		eq = c.deBracket(eq)
	}
	return eq
}

// Operation without parentheses
func (c *Calc) operNonBrk(eq string) string {
	if c.regBracketSg == nil {
		c.regBracketSg = regexp.MustCompile(`^\(.*\)$`)
	}

	// clear bracket
	if c.regBracketSg.MatchString(eq) {
		eq = eq[1:]
		eq = eq[:len(eq)-1]
	}

	// x**y or x^y
	eq = c.pow(eq)
	// add, subtract, multiply and divide => +-*/
	eq = c.mulDiv(eq)
	eq = c.addSub(eq)

	return eq
}

// multiply and divide
func (c *Calc) mulDiv(eq string) string {
	if c.mulDivReg == nil {
		c.mulDivReg = regexp.MustCompile(`(\d+(\.\d+)?)[*/%](\d+(\.\d+)?)`)
	}
	if c.mulDivRegSg == nil {
		c.mulDivRegSg = regexp.MustCompile(`[*/%]`)
	}

	mulDiv := c.mulDivReg.FindAllString(eq, -1)
	for _, md := range mulDiv {
		splitLs := c.mulDivRegSg.Split(md, -1)
		beg := StringAsFloat(splitLs[0])
		end := StringAsFloat(splitLs[1])

		sgList := c.mulDivRegSg.FindAllString(md, -1)
		var cul float64
		eqSg := sgList[0]
		if eqSg == "*" {
			cul = beg * end
		} else if eqSg == "%" {
			cul = math.Mod(beg, end)
		} else {
			cul = beg / end
		}

		eq = strings.ReplaceAll(eq, md, FloatSimple(fmt.Sprintf(c.accuracyStr, cul)))
	}

	if c.mulDivReg.MatchString(eq) {
		eq = c.mulDiv(eq)
	}

	return eq
}

// add, subtract
func (c *Calc) addSub(eq string) string {
	if c.addSubReg == nil {
		c.addSubReg = regexp.MustCompile(`(\d+(\.\d+)?)[+\-](\d+(\.\d+)?)`)
	}
	if c.addSubRegSg == nil {
		c.addSubRegSg = regexp.MustCompile(`[+\-]`)
	}

	addSub := c.addSubReg.FindAllString(eq, -1)
	for _, as := range addSub {
		splitLs := c.addSubRegSg.Split(as, -1)
		beg := StringAsFloat(splitLs[0])
		end := StringAsFloat(splitLs[1])

		sgList := c.addSubRegSg.FindAllString(as, -1)
		var cul float64
		if sgList[0] == "+" {
			cul = beg + end
		} else {
			cul = beg - end
		}

		eq = strings.ReplaceAll(eq, as, FloatSimple(fmt.Sprintf(c.accuracyStr, cul)))
	}

	// Iterative addition and subtraction
	if c.addSubReg.MatchString(eq) {
		eq = c.addSub(eq)
	}
	return eq
}

func (c *Calc) pow(eq string) string {
	if c.powReg == nil {
		c.powReg = regexp.MustCompile(`(\d+(\.\d+)?)(\*{2}|\^)(\d+(\.\d+)?)`)
	}
	if c.powRegSg == nil {
		c.powRegSg = regexp.MustCompile(`\*{2}|\^`)
	}

	if c.powReg.MatchString(eq) {
		powLs := c.powReg.FindAllString(eq, -1)
		for _, pw := range powLs {
			split := c.powRegSg.Split(pw, -1)
			if len(split) != 2 {
				continue
			}
			count := math.Pow(StringAsFloat(split[0]), StringAsFloat(split[1]))
			eq = strings.ReplaceAll(eq, pw, fmt.Sprintf(c.accuracyStr, count))
		}
	}

	if c.powReg.MatchString(eq) {
		eq = c.pow(eq)
	}

	return eq
}

// Exp support functional expression e.g.
// sprt, log, sin, cos, tan.
//
// Notice: Attempt to expose it to external interfaces
func (c *Calc) Exp(eq string) string {
	if c.expReg == nil {
		// 存在嵌套检测的问题，即函数嵌套时提取有误
		c.expReg = regexp.MustCompile(`(?i)(?:sqrt|log|sin|cos|tan)\(.*\)`)
	}

	expSg := "("
	expLs := c.expReg.FindAllString(eq, -1)
	for _, exp := range expLs {
		idx := strings.Index(exp, expSg)
		if idx == -1 {
			continue
		}

		name := strings.ToLower(exp[:idx])
		subExp := exp[idx+1:]
		subExp = subExp[:len(subExp)-1]
		// Internally nested representation processing
		if c.expReg.MatchString(subExp) {
			subExp = c.Exp(subExp)
		}
		subValue := StringAsFloat(c.operNonBrk(subExp))

		switch name {
		case "sqrt":
			subValue = math.Sqrt(subValue)
		case "log":
			subValue = math.Log(subValue)
		case "sin":
			subValue = math.Asin(subValue)
		case "cos":
			subValue = math.Acos(subValue)
		case "tan":
			subValue = math.Atan(subValue)
		}

		eq = strings.ReplaceAll(eq, exp, FloatSimple(fmt.Sprintf(c.accuracyStr, subValue)))
	}

	if c.expReg.MatchString(eq) {
		eq = c.Exp(eq)
	}

	return eq
}

// Clear interfering characters
func (c *Calc) clearEq(eq string) string {
	// Clear interfering characters
	if c.clearReg == nil {
		c.clearReg = regexp.MustCompile(`\s+`)
	}
	eq = c.clearReg.ReplaceAllString(eq, "")

	// clear `100_00` or `100,000`
	if c.simpleReg == nil {
		c.simpleReg = regexp.MustCompile(`\d+[_,]\d+`)
	}
	// clear `100_00` or `100,000`
	if c.simpleRegSg == nil {
		c.simpleRegSg = regexp.MustCompile(`[_,]`)
	}

	smplLs := c.simpleReg.FindAllString(eq, -1)
	for _, smp := range smplLs {
		smpNew := c.simpleRegSg.ReplaceAllString(smp, "")
		eq = strings.ReplaceAll(eq, smp, smpNew)
	}

	return eq
}

func (c *Calc) Count(args ...string) float64 {
	c.accuracyStr = "%." + fmt.Sprintf("%d", c.Accuracy) + "f"
	equality := rock.ExtractParam(c.equality, args...)
	eq := strings.TrimSpace(equality)

	// Clear interfering characters
	eq = c.clearEq(eq)

	c.handlerEq = eq
	c.equality = equality

	c.handlerEq = c.Exp(c.handlerEq)
	c.handlerEq = c.deBracket(c.handlerEq)
	c.handlerEq = c.operNonBrk(c.handlerEq)
	c.result = StringAsFloat(c.handlerEq)
	return c.result
}

func (c *Calc) String() string {
	return FloatSimple(fmt.Sprintf(c.accuracyStr, c.result))
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

func StringAsFloat(s string) float64 {
	f64, _ := strconv.ParseFloat(s, 10)
	return f64
}
