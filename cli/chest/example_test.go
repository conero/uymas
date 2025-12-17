package chest

import "fmt"

func ExampleInputOption() {
	name := InputOption("请输入用户名", "")
	fmt.Println("用户名：", name)

	// Output:
	// Hello World, Uymas
}
