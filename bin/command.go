package bin

import (
	"fmt"
	"github.com/conero/uymas/util/str"
	"reflect"
)

// @Date：   2018/10/30 0030 13:41
// @Author:  Joshua Conero
// @Name:    名称描述

//项目引用
type ActionInterface interface {
}
type SubCmdAlias struct {
	Alias   map[string][]string
	Matched bool
	Self    interface{}
}

type Command struct {
	App App
	SCA *SubCmdAlias
}

// 引用初始化接口
func (c *Command) Init() {
	c.App = *app
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
	fmt.Println("	命令以及初始化成功，请实现项目.")
}

// 二级命令应用
func (c *Command) SubCommandUnfind(subCmd string) {
	fmt.Println(" 二级命令不存在：" + subCmd)
}
