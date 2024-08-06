package cli

import "fmt"

func ExampleNewCli() {
	cl := NewCli()
	cl.Index(func(parser ...ArgsParser) {
		fmt.Println("Hello World, Uymas")
	})
	err := cl.Run()
	if err != nil {
		fmt.Printf("command action run error. ref: %v", err)
	}

	// Output:
	// Hello World, Uymas
}
