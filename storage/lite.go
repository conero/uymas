package storage

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Lite literal variable actual value
type Lite struct {
	variable string
	vType    string // data type
	anyValue Any    // the true map data
}

// NewLite new literal variable
func NewLite(variable string) *Lite {
	variable = strings.TrimSpace(variable)
	lite := &Lite{
		variable: variable,
		anyValue: nil,
	}
	lite.judgeType()
	return lite
}

func (lite *Lite) judgeType() {
	value := lite.variable
	if value == "" {
		lite.vType = LiteralNull
	} else {
		isMapMk := false
		// INT
		// regexp judge the data type
		isIntReg := regexp.MustCompile(`^[\d]+$`)
		if isIntReg.MatchString(value) {
			lite.vType = LiteralInt
			lite.anyValue, _ = strconv.Atoi(value)
			isMapMk = true
		}

		//FLOAT
		if !isMapMk {
			isfloatReg := regexp.MustCompile(`^[\d]+\.[\d]+$`)
			if isfloatReg.MatchString(value) {
				lite.vType = LiteralFloat
				lite.anyValue, _ = strconv.ParseFloat(value, 64)
				isMapMk = true
			}
		}

		//bool
		if !isMapMk {
			switch value {
			case "true", "True":
				lite.anyValue = true
				lite.vType = LiteralBool
				isMapMk = true
			case "false", "False":
				lite.anyValue = false
				lite.vType = LiteralBool
				isMapMk = true
			}
		}

		//string
		if !isMapMk {
			lite.anyValue = value
			lite.vType = LiteralString
		}
	}
}

// GetAny literal turn on
func (lite *Lite) GetAny() Any {
	value := lite.variable
	return value
}

// GetType get value type
func (lite *Lite) GetType() string {
	return lite.vType
}

func (lite *Lite) IsNumber() bool {
	return lite.vType == LiteralInt || lite.vType == LiteralFloat
}

// LiteralExpression support mathematical expression
func LiteralExpression(expression string) (float64, error) {
	reg := regexp.MustCompile(`^[(+\-*/)\d\s.]+$`)
	if reg.MatchString(expression) {
		regBrackets := regexp.MustCompile(`\([^()]+\)`)
		bracketRegRpl := regexp.MustCompile(`\(|\)`)
		for {
			if !regBrackets.MatchString(expression) {
				break
			}
			brackets := regBrackets.FindAllString(expression, -1)
			for _, exp := range brackets {
				exp2 := bracketRegRpl.ReplaceAllString(exp, "")
				value := ExpNoBracket(exp2)
				fmt.Println(exp2, value)
				expression = strings.ReplaceAll(expression, exp, value)
			}
		}
	}
	return 0, errors.New("expression not a valid mathematical expression")
}

// ExpNoBracket the expression without bracket
func ExpNoBracket(expression string) string {
	var result string = "0"
	reg := regexp.MustCompile(`^[+\-*/\d\s.]+$`)
	if reg.MatchString(expression) {
		mulDivReg := regexp.MustCompile(`[\d.]+[*/]+[\d.]`)
		mulDivSignReg := regexp.MustCompile(`\*|\/`)
		checkReg := regexp.MustCompile(`^[\d.]+$`)

		for {
			if !mulDivReg.MatchString(expression) {
				break
			}

			if checkReg.MatchString(expression) {
				break
			}

			for _, frag := range mulDivReg.FindAllString(expression, -1) {
				fragQue := mulDivSignReg.Split(frag, -1)
				frag1, frag2 := fragQue[0], fragQue[1]
				fragV1, _ := strconv.ParseFloat(frag1, 64)
				fragV2, _ := strconv.ParseFloat(frag2, 64)

				rpl := ""
				if strings.Contains(frag, "/") {
					rpl = fmt.Sprintf("%v", fragV1/fragV2)
				} else if strings.Contains(frag, "*") {
					rpl = fmt.Sprintf("%v", fragV1*fragV2)
				}

				expression = strings.ReplaceAll(expression, frag, rpl)
			}
		}

	}

	return result
}
