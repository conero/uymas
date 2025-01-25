package calc

import (
	"fmt"
)

func ExampleNewCalc() {
	cl := NewCalc("3!+2pi")
	cl.Count()
	fmt.Printf("%v\n", cl.String())

	// 等式中指定精度
	cl = NewCalc("f17, 3!+2pi")
	cl.Count()
	fmt.Printf("%v\n", cl.String())

	cl.Count("3-(2^2-pi)")
	fmt.Printf("%v\n", cl.String())

	// Output:
	// 12.2831854
	// 12.28318530717958623
	// 2.14159265358979312
}
