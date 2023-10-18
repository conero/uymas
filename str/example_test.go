package str

import (
	"fmt"
)

/**
 * @DATE        2019/6/3
 * @NAME        Joshua Conero
 * @DESCRIPIT   str 例子
 **/

func ExampleUrl_AbsHref() {
	var u Url
	// "joshua/conero" 与 "./joshua/conero" 效果相同
	fmt.Println(u.AbsHref("joshua/conero", "https://www.about.me/url/example/test"))

	fmt.Println(u.AbsHref("/joshua/conero", "https://www.about.me/url/example/test"))
	// "//" 等符合可被清除
	fmt.Println(u.AbsHref("//joshua/conero", "https://www.about.me/url/example/test"))
	fmt.Println(u.AbsHref("../../joshua/conero", "https://www.about.me/url/example/test"))
	// Output:
	// https://www.about.me/url/example/test/joshua/conero
	// https://www.about.me/joshua/conero
	// https://www.about.me/joshua/conero
	// https://www.about.me/joshua/conero
}

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
