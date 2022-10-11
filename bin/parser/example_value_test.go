package parser

import (
	"fmt"
	"testing"
)

// func TestExampleConvBool(t *testing.T) {
func ExampleConvBool() {
	var raw = "True"
	var cv bool

	// case
	cv = ConvBool(raw)
	fmt.Printf("ConvBool:%v => %v\n", raw, cv)

	// case
	raw = "TRUE"
	cv = ConvBool(raw)
	fmt.Printf("ConvBool:%v => %v\n", raw, cv)

	// case
	raw = "vaild"
	cv = ConvBool(raw)
	fmt.Printf("ConvBool:%v => %v\n", raw, cv)

	//Output:
	//ConvBool:True => true
	//ConvBool:TRUE => true
	//ConvBool:vaild => false
}

func ExampleConvI64() {
	var raw = "1949"
	var cv int64

	// case
	cv = ConvI64(raw)
	fmt.Printf("ConvI64:%v => %v\n", raw, cv)

	// case
	raw = "+1949"
	cv = ConvI64(raw)
	fmt.Printf("ConvI64:%v => %v\n", raw, cv)

	// case
	raw = "-1949.1001"
	cv = ConvI64(raw)
	fmt.Printf("ConvBool:%v => %v\n", raw, cv)

	// case
	raw = "-2022"
	cv = ConvI64(raw)
	fmt.Printf("ConvI64:%v => %v\n", raw, cv)

	// case
	raw = "vaild"
	cv = ConvI64(raw)
	fmt.Printf("ConvI64:%v => %v\n", raw, cv)

	//Output:
	//ConvI64:1949 => 1949
	//ConvI64:+1949 => 1949
	//ConvI64:-1949.1001 => -1949
	//ConvI64:-2022 => -2022
	//ConvI64:vaild => 0
}

func ExampleConvF64() {
	var raw = "1949"
	var cv float64

	// case
	cv = ConvF64(raw)
	fmt.Printf("ConvF64:%v => %v\n", raw, cv)

	// case
	raw = "+1949"
	cv = ConvF64(raw)
	fmt.Printf("ConvF64:%v => %v\n", raw, cv)

	// case
	raw = "-1949.1001"
	cv = ConvF64(raw)
	fmt.Printf("ConvF64:%v => %v\n", raw, cv)

	// case
	raw = "-3.14159265359"
	cv = ConvF64(raw)
	fmt.Printf("ConvI64:%v => %v\n", raw, cv)

	// case
	raw = "vaild"
	cv = ConvF64(raw)
	fmt.Printf("ConvI64:%v => %v\n", raw, cv)

	//Output:
	//ConvF64:1949 => 1949
	//ConvF64:+1949 => 1949
	//ConvF64:-1949.1001 => -1949.1001
	//ConvI64:-3.14159265359 => -3.14159265359
	//ConvI64:vaild => 0
}

func ExampleConvInt() {
	var raw = "1949"
	var cv int

	// case
	cv = ConvInt(raw)
	fmt.Printf("ConvInt:%v => %v\n", raw, cv)

	// case
	raw = "+1949"
	cv = ConvInt(raw)
	fmt.Printf("ConvInt:%v => %v\n", raw, cv)

	// case
	raw = "-1949.1001"
	cv = ConvInt(raw)
	fmt.Printf("ConvInt:%v => %v\n", raw, cv)

	// case
	raw = "-2022"
	cv = ConvInt(raw)
	fmt.Printf("ConvInt:%v => %v\n", raw, cv)

	// case
	raw = "vaild"
	cv = ConvInt(raw)
	fmt.Printf("ConvInt:%v => %v\n", raw, cv)

	//Output:
	//ConvInt:1949 => 1949
	//ConvInt:+1949 => 1949
	//ConvInt:-1949.1001 => -1949
	//ConvInt:-2022 => -2022
	//ConvInt:vaild => 0
}

// 用于任意测试
func TestConvInt_base(t *testing.T) {
	//ExampleConvI64()
	//ExampleConvF64()
	ExampleConvInt()
}
