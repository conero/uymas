package bin

import "github.com/conero/uymas/util"

// @Date：   2018/10/30 0030 12:40
// @Author:  Joshua Conero
// @Name:    应用管理

type App struct {
	Data       map[string]interface{}
	Router     *Router
	cwd        string   // 项目当前所在目录
	Command    string   // 当前的命令
	SubCommand string   // 二级命令
	Setting    []string // 项目设置
}

/**
检测属性是否存在
*/
func (app App) HasSetting(set string) bool {
	has := false
	if idx := util.InStrQue(set, app.Setting); idx > -1 {
		has = true
	}
	return has
}

/**
获取当的工作目录
*/
func (app App) Cwd() string {
	return app.cwd
}
