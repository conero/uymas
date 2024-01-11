package digit

import (
	"fmt"
	"gitee.com/conero/uymas/util/rock"
	"math"
	"strconv"
	"strings"
)

type Cover float64

// ToChnUpper Convert to uppercase Chinese numerals
func (c Cover) ToChnUpper() {
	//@todo
	fmt.Printf("%v => %v\n", int64(c), c.ToChnRoundUpper())
}

func (c Cover) ToChnRoundUpper() string {
	return NumberCoverChnDigit(float64(c))
}

func (c Cover) ToChnRoundLower() string {
	return NumberCoverChnDigit(float64(c), false)
}

func (c Cover) ToRmbUpper() string {
	return NumberCoverRmb(float64(c), true)
}

func (c Cover) ToRmbLower() string {
	return NumberCoverRmb(float64(c), false)
}

// NumberCoverChnDigit Arabic numerals to Chinese numerals, supporting uppercase and lowercase
func NumberCoverChnDigit(latest float64, isUpperDef ...bool) string {
	isUpper := rock.ExtractParam(true, isUpperDef...)
	var numbers []string
	var unitList = []int{UnitYValue, UnitWValue, UnitQValue, UnitBValue, UnitSValue}

	var vMap map[int8]string
	if isUpper {
		vMap = vUpperMap
	} else {
		vMap = vLowerMap
	}

	for _, unit := range unitList {
		cvUnit := float64(unit)
		if latest < cvUnit {
			if latest < UnitSValue && latest > 0 {
				numbers = append(numbers, vMap[int8(latest)])
				latest = 0
				break
			}
			continue
		}
		value := int(math.Floor(latest / cvUnit))
		if value > 10 {
			numbers = append(numbers, NumberCoverChnDigit(float64(value), isUpper))
		} else {
			numbers = append(numbers, vMap[int8(value)])
		}

		latest = latest - float64(value)*cvUnit
		switch cvUnit {
		case UnitYValue:
			var unitStr string
			if isUpper {
				unitStr = UnitUpperY
			} else {
				unitStr = UnitLowerY
			}
			numbers = append(numbers, unitStr)
		case UnitWValue:
			var unitStr string
			if isUpper {
				unitStr = UnitUpperW
			} else {
				unitStr = UnitLowerW
			}
			numbers = append(numbers, unitStr)
		case UnitQValue:
			var unitStr string
			if isUpper {
				unitStr = UnitUpperQ
			} else {
				unitStr = UnitLowerQ
			}
			numbers = append(numbers, unitStr)
		case UnitBValue:
			var unitStr string
			if isUpper {
				unitStr = UnitUpperB
			} else {
				unitStr = UnitLowerB
			}
			numbers = append(numbers, unitStr)
		case UnitSValue:
			var unitStr string
			if isUpper {
				unitStr = UnitUpperS
			} else {
				unitStr = UnitLowerS
			}
			numbers = append(numbers, unitStr)
		case 0:
			numbers = append(numbers, vMap[int8(value)])
		}

		// zero fill
		if cvUnit > UnitSValue && latest > 0 && (cvUnit/10)-latest > 0 {
			numbers = append(numbers, vMap[0])
		}
	}

	// Final remaining quantity
	if latest < UnitSValue && latest > 0 {
		numbers = append(numbers, vMap[int8(latest)])
	}
	return strings.Join(numbers, "")
}

// BUG(who): NumberCoverRmb 6.01 -> math.Modf frac is inaccurate.link: https://github.com/golang/go/issues/62232

// NumberCoverRmb Transforming Numbers into People's Digital Writing
func NumberCoverRmb(amount float64, isUpperDef ...bool) string {
	isUpper := rock.ExtractParam(true, isUpperDef...)
	val, _ := math.Modf(amount)

	// Processing decimals
	amountStr := fmt.Sprintf("%#v", amount)
	splitIdx := strings.Index(amountStr, ".")
	if splitIdx > 0 {
		endIndex := splitIdx + 3
		strLen := len(amountStr)
		if endIndex > strLen {
			endIndex = strLen
		}
		amountStr = amountStr[splitIdx+1 : endIndex]
	} else {
		amountStr = "0"
	}

	strLen := len(amountStr)
	// Complete the number of pure angular positions (without quantiles) with 0
	if strLen == 1 {
		amountStr += "0"
	}
	fracInt, _ := strconv.Atoi(amountStr)
	var str string
	if val > 0 {
		str = NumberCoverChnDigit(val, isUpper)
	}
	if str != "" {
		str += "元"
	}
	if fracInt == 0 {
		str += "整"
	} else if fracInt >= 10 {
		jiaoValue := int(float64(fracInt) * 0.1)
		str += NumberCoverChnDigit(float64(jiaoValue), isUpper)
		str += "角"
		fenValue := fracInt % 10
		if fenValue > 0 {
			fen := NumberCoverChnDigit(float64(fenValue), isUpper)
			if fen != "" {
				fen += "分"
				str += fen
			}
		}
	} else if fracInt < 10 {
		fen := NumberCoverChnDigit(float64(fracInt%10), isUpper)
		if fen != "" {
			fen += "分"
			str += fen
		}
	}
	return str
}
