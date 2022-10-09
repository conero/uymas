package main

import (
	"fmt"
	"gitee.com/conero/uymas/bin"
	"gitee.com/conero/uymas/bin/tag"
)

// Test 命令
type Test struct {
}

func (c *Test) Exec(cc *bin.CliCmd) {
	fmt.Println("test 命令引用入口！")
}

// App
// @todo 实现 struct 到 bin 的映射
type App struct {
	CTY  tag.Name `cmd:"app:yang"`
	Test *Test    `cmd:"command:test alias:tst,t help:测试命令工具"`
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
