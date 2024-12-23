// Package calc enter string equations to implement numeric arithmetic
package calc

import (
	"fmt"
	"gitee.com/conero/uymas/v2/number"
	"gitee.com/conero/uymas/v2/rock"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const CalcAccuracy int8 = 7

var (
	cacheCalc *Calc
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
	facReg      *regexp.Regexp
	constReg    *regexp.Regexp
}

// NewCalc String Equation Calculation
// support `f8,exp` to set .Accuracy
func NewCalc(equality string) *Calc {
	idx := strings.Index(equality, ",")
	accVal := CalcAccuracy
	if idx > -1 {
		accStr := strings.ToLower(equality[:idx])
		if strings.Index(accStr, "f") == 0 {
			accStr = accStr[1:]
			v, err := strconv.ParseInt(accStr, 10, 8)
			if err == nil && v > 0 {
				accVal = int8(v)
			}
		}
		equality = equality[idx+1:]
	}

	return &Calc{
		equality: equality,
		Accuracy: accVal,
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

	// 支持 e/pi
	eq = c.constSupt(eq)
	// `n!`
	eq = c.factorial(eq)
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

func (c *Calc) factorial(eq string) string {
	if c.facReg == nil {
		c.facReg = regexp.MustCompile(`\d+!`)
	}

	if !c.facReg.MatchString(eq) {
		return eq
	}

	for _, fd := range c.facReg.FindAllString(eq, -1) {
		fdVal := fd[:len(fd)-1]
		val := number.Factorial(uint64(StringAsI64(fdVal)))
		eq = strings.ReplaceAll(eq, fd, fmt.Sprintf("%d", val))
	}

	if c.facReg.MatchString(eq) {
		eq = c.factorial(eq)
	}

	return eq
}

// support const like, pi/e
func (c *Calc) constSupt(eq string) string {
	if c.constReg == nil {
		c.constReg = regexp.MustCompile(`(?i)\d*(e|pi)\d*`)
	}

	if !c.constReg.MatchString(eq) {
		return eq
	}
	for _, fd := range c.constReg.FindAllString(eq, -1) {
		fdEq := strings.ToLower(fd)
		if fdEq == "e" || fdEq == "pi" {
			var value float64
			if fdEq == "pi" {
				value = math.Pi
			} else {
				value = math.E
			}

			eq = strings.ReplaceAll(eq, fd, fmt.Sprintf(c.accuracyStr, value))
			continue
		}

		// e
		idx := strings.Index(fdEq, "e")
		if idx > -1 {
			if idx == 0 { // 开头
				fdEq = strings.ReplaceAll(fdEq, "e", fmt.Sprintf(c.accuracyStr, math.E)+"*")
			} else if idx == len(fdEq)-1 { // 结尾
				fdEq = strings.ReplaceAll(fdEq, "e", "*"+fmt.Sprintf(c.accuracyStr, math.E))
			} else {
				fdEq = strings.ReplaceAll(fdEq, "e", "*"+fmt.Sprintf(c.accuracyStr, math.E)+"*")
			}
			eq = strings.ReplaceAll(eq, fd, fdEq)
			continue
		}

		// pi
		idx = strings.Index(fdEq, "pi")
		if idx > -1 {
			if idx == 0 { // 开头
				fdEq = strings.ReplaceAll(fdEq, "pi", fmt.Sprintf(c.accuracyStr, math.Pi)+"*")
			} else if idx == len(fdEq)-2 { // 结尾
				fdEq = strings.ReplaceAll(fdEq, "pi", "*"+fmt.Sprintf(c.accuracyStr, math.Pi))
			} else {
				fdEq = strings.ReplaceAll(fdEq, "pi", "*"+fmt.Sprintf(c.accuracyStr, math.Pi)+"*")
			}
			eq = strings.ReplaceAll(eq, fd, fdEq)
			continue
		}
	}

	if c.constReg.MatchString(eq) {
		eq = c.constSupt(eq)
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
		// sin 与 sinh，带h的为双曲线三角函数
		c.expReg = regexp.MustCompile(`(?i)(sqrt|log|log2|log10|sin|cos|tan|sinh|cosh|tanh|asin|acos|atan|asinh|acosh|atanh|abs)\(`)
	}

	// 嵌套环境下：由于正则表达式无法实现取反，遂采用其他方法进行改造
	if !c.expReg.MatchString(eq) {
		return eq
	}

	index := c.expReg.FindAllStringIndex(eq, -1)
	size := len(index)
	for x, idx := range index {
		var sLen int
		if size > x+1 {
			sLen = index[x+1][0]
		} else {
			sLen = strings.LastIndex(eq, ")") + 1
		}
		name := eq[idx[0] : idx[1]-1]

		// 获取原始表达式
		rawEq := eq[idx[0]:sLen]
		braB := strings.Count(rawEq, "(")
		braE := strings.Count(rawEq, ")")

		// 闭合的括号＞开始括号，进行字符串切割
		if braE > braB {
			scanCtt := 0
			for bIdx, b := range []byte(rawEq) {
				scanChar := string(b)
				if scanChar == ")" {
					scanCtt += 1
				}
				if scanCtt == braB {
					rawEq = rawEq[:bIdx+1]
					braE = scanCtt
					break
				}
			}
		} else {
			scanCtt := 0
			for bIdx, b := range []byte(rawEq) {
				scanChar := string(b)
				if scanChar == ")" {
					scanCtt += 1
				}
				if scanCtt == braB {
					rawEq = rawEq[:bIdx+1]
					braE = scanCtt
					break
				}
			}
		}

		if braB != braE {
			continue
		}

		subValue := c.expCalc(name, rawEq)
		eq = strings.ReplaceAll(eq, rawEq, FloatSimple(fmt.Sprintf(c.accuracyStr, subValue)))
		break
	}

	if c.expReg.MatchString(eq) {
		eq = c.Exp(eq)
	}

	return eq
}

// 表达式实际执行
func (c *Calc) expCalc(name, eq string) float64 {
	eqClear := eq[len(name)+1 : len(eq)-1]

	child := CalcEq(eqClear)
	subValue := StringAsFloat(child.String())

	switch name {
	case "sqrt":
		subValue = math.Sqrt(subValue)
	case "log":
		subValue = math.Log(subValue)
	case "log2":
		subValue = math.Log2(subValue)
	case "log10":
		subValue = math.Log10(subValue)
	case "sin":
		subValue = math.Sin(subValue)
	case "cos":
		subValue = math.Cos(subValue)
	case "tan":
		subValue = math.Tan(subValue)
	case "sinh":
		subValue = math.Sinh(subValue)
	case "cosh":
		subValue = math.Cosh(subValue)
	case "tanh":
		subValue = math.Tanh(subValue)
	case "asin":
		subValue = math.Asin(subValue)
	case "acos":
		subValue = math.Acos(subValue)
	case "atan":
		subValue = math.Atan(subValue)
	case "asinh":
		subValue = math.Asinh(subValue)
	case "acosh":
		subValue = math.Acosh(subValue)
	case "atanh":
		subValue = math.Atanh(subValue)
	case "abs": // absolute-value
		subValue = math.Abs(subValue)
	}
	return subValue
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
	equality := rock.Param(c.equality, args...)
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
	f64, _ := strconv.ParseFloat(s, 64)
	return f64
}

func StringAsI64(s string) int64 {
	v, _ := strconv.ParseInt(s, 10, 60)
	return v
}

func CalcEq(eq string) Calc {
	if cacheCalc == nil {
		cacheCalc = NewCalc("")
	}
	cacheCalc.Count(eq)
	return *cacheCalc
}

// NumberSplitFormat numeric value segmentation and beautification
func NumberSplitFormat(n float64, bits ...int) string {
	bit := rock.Param(3, bits...)
	s := fmt.Sprintf("%f", n)
	queue := strings.Split(s, ".")
	vInt := queue[0]
	vLen := len(vInt)
	if vLen <= bit {
		return FloatSimple(s)
	}
	var subQueue []string
	var last = 0
	for i := 0; i < vLen; i++ {
		c := i + 1
		if c == vLen || c%bit != 0 {
			continue
		}
		start := vLen - c
		subQueue = append([]string{vInt[start : start+3]}, subQueue...)
		last = start
	}

	if last < vLen {
		subQueue = append([]string{vInt[:last]}, subQueue...)
	}

	queue[0] = strings.Join(subQueue, ",")
	if len(queue) > 1 {
		// 全零
		if queue[1] == strings.Repeat("0", len(queue[1])) {
			queue = append([]string{}, queue[0])
		}
	}
	return strings.Join(queue, ".")
}

// NumberClear number string clear like '_' or ','
func NumberClear(s string) string {
	matched, _ := regexp.MatchString(`\d+([_,]\d)+\d*(\.\d)*`, s)
	if !matched {
		return s
	}
	clrReg := regexp.MustCompile(`[_,]`)
	return clrReg.ReplaceAllString(s, "")
}
