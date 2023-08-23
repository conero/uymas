package digit

import (
	"fmt"
	"math"
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

// NumberCoverChnDigit Arabic numerals to Chinese numerals, supporting uppercase and lowercase
func NumberCoverChnDigit(latest float64, isUpperDef ...bool) string {
	isUpper := true
	if len(isUpperDef) > 0 {
		isUpper = isUpperDef[0]
	}
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
		if cvUnit > UnitSValue && (cvUnit/10)-latest > 0 {
			numbers = append(numbers, vMap[0])
		}
	}

	// Final remaining quantity
	if latest < UnitSValue && latest > 0 {
		numbers = append(numbers, vMap[int8(latest)])
	}
	return strings.Join(numbers, "")
}
