package main

import (
	"fmt"
	"gitee.com/conero/uymas/v2/cli/ansi"
	"gitee.com/conero/uymas/v2/cli/chest"
	"gitee.com/conero/uymas/v2/cli/evolve"
	"gitee.com/conero/uymas/v2/util/fs"
)

type test struct {
	evolve.Command
}

func (c *test) Demo() {
	fmt.Println("test demo, ha!")
	fmt.Println()
	fmt.Println("rootPath: " + fs.RootPath())
	fmt.Println("rootApp: " + fs.AppName())
}

func (c *test) DefHelp() {
	fmt.Println("测试命令，全局选项如：")
	fmt.Println(" --cmd [command]          判断命令是否存在")
	fmt.Println()
	fmt.Println(" demo                     普通测试用例")
}

func (c *test) DefIndex() {
	command := c.X.Args.Get("cmd")
	if command != "" {
		exist, err := chest.CmdExist(command)
		if err != nil {
			fmt.Printf("命令行检测错误，%s\n", err)
		} else if exist {
			fmt.Printf("%s：命令存在\n", command)
		} else {
			fmt.Printf("%s：%s\n", command, ansi.Style("命令不存在", ansi.Red))
		}
		return
	}
	fmt.Println("您在使用 test 命令")
}

func main() {
	evl := evolve.NewEvolve()
	evl.Command(func() {
		fmt.Println("Evolution For Index.")
	}, "index")
	evl.Command(new(test), "test", "t")
	//evl.Run("test", "demo")
	evl.Run()
}
