package tag

func ExampleParser_Run() {
	type test struct {
		Name string `cmd:"option:name require help:输入姓名"`
		Test string `cmd:"option:test help:输入test 表达式"`
	}

	// the app main struct
	type app struct {
		CTY  Name  `cmd:"app:yang"`
		Test *test `cmd:"command:test alias:tst,t help:测试命令工具 valuable"`
	}

	//func main(){
	myapp := &app{
		Test: &test{},
	}
	parser := NewParser(myapp)
	parser.Run()
	//}

	// Output:
	//欢饮使用命令行程序，命令格式如下:
	//
	//$ yang [command] [option]
	//
	//命令列表:
	//test   测试命令工具，别名 tst,t
}
