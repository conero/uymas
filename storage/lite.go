package storage

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

//字面变量实际值
type Lite struct {
	variable string
	vType    string // data type
	anyValue Any    // the true map data
}

//new literal variable
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

//literal turn on
func (lite *Lite) GetAny() Any {
	value := lite.variable
	return value
}

//获取数据类型
func (lite *Lite) GetType() string {
	return lite.vType
}

func (lite *Lite) IsNumber() bool {
	return lite.vType == LiteralInt || lite.vType == LiteralFloat
}

//@todo need to do.
// support mathematical expression
func LiteralExpression(expression string) (float64, error) {
	reg := regexp.MustCompile(`^[(+\-*/)\d]+$`)
	if reg.MatchString(expression) {
		regBrackets := regexp.MustCompile(`\(^[()]+\)`)
		for {
			if !regBrackets.MatchString(expression) {
				break
			}
			regBrackets.FindAllString(expression, -1)
		}
	}
	return 0, errors.New("expression not a valid mathematical expression")
}
