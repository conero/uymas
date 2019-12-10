package bin

import (
	"fmt"
	"github.com/conero/uymas/str"
	"reflect"
)

// @Date：   2018/10/30 0030 13:41
// @Author:  Joshua Conero
// @Name:    名称描述

//项目引用
type ActionInterface interface {
}

// 二级命令别名
type SubCmdAlias struct {
	Alias   map[string][]string
	Matched bool
	Self    interface{}
}

// 命令结构体
type Command struct {
	App  *App
	SCA  *SubCmdAlias
	Util *CmdUitl
}

// 引用初始化接口
func (c *Command) Init() {
	c.App = GetApp()
	c.Util = &CmdUitl{c}
}

// 入口/内部分发(Entrance)
func (c *Command) InnerDistribute() {
	if c.SCA != nil {
		sca := c.SCA
		for sub, alias := range sca.Alias {
			for _, a := range alias {
				if c.App.HasSetting(a) {
					method := reflect.ValueOf(sca.Self).MethodByName(str.Ucfirst(sub))
					method.Call(nil)
					sca.Matched = true
					break
				}
			}
			if sca.Matched {
				break
			}
		}
	}
}

// 运行应用
func (c *Command) Run() {
	fmt.Println("	命令以及初始化成功，请实现项目(Command.Run).")
}

// 二级命令应用
func (c *Command) SubCommandUnfind(subCmd string) {
	fmt.Println("  二级命令不存在：" + subCmd)
}

// 帮助说明
func (c *Command) Help() {
	fmt.Println("  项目帮助说明，外部通过： $ help [访问] 来查看对应的命令帮助")
}

// Command 协助方法
// 通过 cmdInst 与 命令程序解析
type CmdUitl struct {
	cmdInst *Command
}

// 二级命令别名
func (cu *CmdUitl) BaseSubCAlias(inst interface{}, alias map[string][]string) *SubCmdAlias {
	if alias == nil {
		alias = map[string][]string{}
	}
	// 帮助程序
	alias["help"] = []string{"h"}

	// 获取二级对象
	csa := &SubCmdAlias{
		Alias: alias,
		Self:  inst,
	}
	return csa
}
