package bin

import (
	"fmt"
	"testing"
)

func ExampleNewCLI_func() {
	cli := NewCLI()
	cli.RegisterAny(func() {
		fmt.Println("Hello world, cli.")
	})
	cli.Run()

	// Output:
	// Hello world, cli.
}

type defaultAppHelloWorld struct {
	CliApp
}

func (w *defaultAppHelloWorld) DefaultIndex() {
	fmt.Println("Hello world, struct cli.")
}

func ExampleNewCLI_struct() {
	cli := NewCLI()
	// add the new struct instance.
	// type defaultAppHelloWorld struct {
	//	bin.CliApp
	// }
	//
	// func (w *defaultAppHelloWorld) DefaultIndex() {
	//	fmt.Println("Hello world, struct cli.")
	// }
	cli.RegisterAny(new(defaultAppHelloWorld))
	cli.Run()

	// Output:
	// Hello world, struct cli.
}

// to run Example func tpl
func TestNewCLI(t *testing.T) {
	//ExampleXXX
	//ExampleNewCLI_func()
	ExampleNewCLI_struct()
}
