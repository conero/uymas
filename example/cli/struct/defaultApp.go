package main

import (
	"fmt"
	"gitee.com/conero/uymas/v2/cli/evolve"
)

type optDemo struct {
	Number int `cmd:"number,n default:1024 help:整形字符串获取"`
}

type OptTest struct {
	Key string `cmd:"key help:获取测试值索引"`
}

type defaultApp struct {
	evolve.Command
	IsVerbose bool    `cmd:"verbose,vv globalOwner help:详细输出内容"`
	Dir       string  `cmd:"dir,D globalOwner help:设置工作目录"`
	OptDemo   optDemo `cmd:"owner:demo help:命令用例测试"`
	OptTest   OptTest `cmd:"owner:test notValid"`
}

func (c *defaultApp) DefIndex() {
	fmt.Println("Hello World, Struct.")
}

func (c *defaultApp) Demo() {
	opt := c.OptDemo
	fmt.Println("Demo command")
	if c.IsVerbose {
		fmt.Println("  详细输出……")
	}
	fmt.Println()
	fmt.Println("输出选项测试值")
	fmt.Printf("  Number: %d\n", opt.Number)
	fmt.Printf("  Dir(global): %s\n", c.Dir)
	fmt.Printf("  verbose(global): %v\n", c.IsVerbose)
}

func (c *defaultApp) Test() {
	fmt.Println("Test command")
}
