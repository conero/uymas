package bin

// @Date：   2018/10/30 0030 13:41
// @Author:  Joshua Conero
// @Name:    名称描述

//项目引用
type ActionInterface interface {
}

type Command struct {
	App App
}

// 引用初始化接口
func (command *Command) Init() {
	command.App = *app
}
