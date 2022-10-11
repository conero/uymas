package main

import (
	"fmt"
	"gitee.com/conero/uymas/bin"
	"gitee.com/conero/uymas/bin/tag"
)

// Test 命令
type Test struct {
	Name string `cmd:"option:name require help:输入姓名"`
	Test string `cmd:"option:test help:输入test 表达式"`
}

func (c *Test) Exec(cc *bin.Arg) {
	fmt.Println("test 命令引用入口！")
	fmt.Printf("name: %v\n", c.Name)
}

// App
// @todo 实现 struct 到 bin 的映射
type App struct {
	CTY  tag.Name `cmd:"app:yang"`
	Test *Test    `cmd:"command:test alias:tst,t help:测试命令工具 valuable"`
	//Commands []any
}

func main() {
	app := &App{
		Test: &Test{},
		//Commands: []any{
		//	&Test{},
		//},
	}
	parser := tag.NewParser(app)
	parser.Run()
}
