package digit

import (
	"errors"
	"fmt"
	"gitee.com/conero/uymas/v2/rock"
	"math"
	"strconv"
	"strings"
)

type Cover float64

// ToChnUpper Convert to uppercase Chinese numerals
//
// Used to convert a number as an integer to a Chinese uppercase number
func (c Cover) ToChnUpper() string {
	return NumberCoverChnDigit(float64(int64(c)))

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

// NumberCover convert numbers to Chinese style with definable dictionary
func NumberCover(latest float64, vMap map[uint32]string) (dgt string, err error) {
	if vMap == nil || len(vMap) == 1 {
		err = errors.New("vMap is empty")
		return
	}
	var unitList = []uint32{UnitYValue, UnitWValue, UnitQValue, UnitBValue, UnitSValue}
	var numbers []string
	for _, unit := range unitList {
		cvUnit := float64(unit)
		if latest < cvUnit {
			if latest < 10 && latest > 0 {
				numbers = append(numbers, vMap[uint32(latest)])
				latest = 0
				break
			}
			continue
		}
		value := uint32(math.Floor(latest / cvUnit))
		if value > 10 {
			childStr, childEr := NumberCover(float64(value), vMap)
			if childEr != nil {
				err = childEr
				return
			}
			numbers = append(numbers, childStr)
		} else {
			numbers = append(numbers, vMap[value])
		}

		latest = latest - float64(value)*cvUnit
		unitStr, exist := vMap[unit]
		if !exist {
			err = fmt.Errorf("%d: as key not exist in vMap", unit)
			return
		}
		numbers = append(numbers, unitStr)

		// zero fill
		if cvUnit > 10 && latest > 0 && (cvUnit/10)-latest > 0 {
			numbers = append(numbers, vMap[0])
		}
	}

	// Final remaining quantity
	if latest < 10 && latest > 0 {
		numbers = append(numbers, vMap[uint32(latest)])
	}

	dgt = strings.Join(numbers, "")
	return
}

// NumberCoverChnDigit Arabic numerals to Chinese numerals, supporting uppercase and lowercase
func NumberCoverChnDigit(latest float64, isUpperDef ...bool) string {
	isUpper := rock.Param(true, isUpperDef...)
	var vMap = vUpperMap
	if !isUpper {
		vMap = vLowerMap
	}

	dgt, err := NumberCover(latest, vMap)
	if err != nil {
		return ""
	}
	return dgt
}

// BUG(who): NumberCoverRmb 6.01 -> math.Modf frac is inaccurate.link: https://github.com/golang/go/issues/62232

// NumberCoverRmb Transforming Numbers into People's Digital Writing
func NumberCoverRmb(amount float64, isUpperDef ...bool) string {
	isUpper := rock.Param(true, isUpperDef...)
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
