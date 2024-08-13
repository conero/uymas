package ansi

import (
	"fmt"
	"testing"
)

func TestStyle(t *testing.T) {
	styleText := Style(2019, Black, BkgCyan)
	fmt.Println(styleText)
}
