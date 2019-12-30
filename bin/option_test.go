package bin

import (
	"fmt"
	"testing"
)

func TestOption_GetDescrip(t *testing.T) {
	opt := Option{
		Key:         "password",
		Description: "",
		Logogram:    "",
	}
	fmt.Println(opt.GetDescrip())
	fmt.Println(opt.Logogram)
}
