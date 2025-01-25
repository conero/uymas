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
