package bin

import (
	"fmt"
	"testing"
)

func TestOption_Unmarshal(t *testing.T) {
	type base struct {
		Name        string
		DisplayName string `arg:"d, display"`
	}

	cli := NewCLI()
	cli.RegisterEmpty(func(cc *Arg) {
		opt := &Option{cc}
		var bv base
		opt.Unmarshal(&bv)

		fmt.Println(bv)
	})

	cli.Run("-d", "Joshua", "--name", "xyz")
}
