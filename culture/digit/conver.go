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
	var numbers []string
	var unitList = []int{UnitYValue, UnitWValue, UnitQValue, UnitBValue, UnitSValue}

	var latest = c
	for _, unit := range unitList {
		cvUnit := Cover(unit)
		if latest < cvUnit {
			if latest < UnitSValue && latest > 0 {
				numbers = append(numbers, vUpperMap[int8(latest)])
				latest = 0
				break
			}
			continue
		}
		value := int(math.Floor(float64(latest / cvUnit)))
		if value > 10 {
			var pValue = Cover(value)
			numbers = append(numbers, pValue.ToChnRoundUpper())
		} else {
			numbers = append(numbers, vUpperMap[int8(value)])
		}

		latest = latest - Cover(value)*cvUnit
		switch cvUnit {
		case UnitYValue:
			numbers = append(numbers, UnitUpperY)
		case UnitWValue:
			numbers = append(numbers, UnitUpperW)
		case UnitQValue:
			numbers = append(numbers, UnitUpperQ)
		case UnitBValue:
			numbers = append(numbers, UnitUpperB)
		case UnitSValue:
			numbers = append(numbers, UnitUpperS)
		case 0:
			numbers = append(numbers, vUpperMap[int8(value)])
		}

		// zero fill
		if cvUnit > UnitSValue && cvUnit-latest > 10 {
			numbers = append(numbers, vUpperMap[0])
		}
	}

	// Final remaining quantity
	if latest < UnitSValue && latest > 0 {
		numbers = append(numbers, vUpperMap[int8(latest)])
	}
	return strings.Join(numbers, "")
}
