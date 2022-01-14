package number

import "fmt"

func ExampleUnit_Unit() {
	n := Unit(5329742)
	fmt.Printf("5329742: %v\n", n)

	// Output:
	// 5329742: 5.3297 M
}
