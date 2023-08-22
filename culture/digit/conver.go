package digit

import (
	"fmt"
	"strings"
)

type Cover float64

// ToChnUpper Convert to uppercase Chinese numerals
func (c Cover) ToChnUpper() {
	//@todo
	fmt.Printf("=> %v\n", c)
}

func (c Cover) ToChnRoundUpper() string {
	var numbers []string
	//cv := c
	// @todo

	return strings.Join(numbers, "")
}
