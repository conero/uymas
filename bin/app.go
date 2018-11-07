package bin

import (
	"github.com/conero/uymas/util/str"
	"path"
)

// @Date：   2018/10/30 0030 12:40
// @Author:  Joshua Conero
// @Name:    应用管理

type App struct {
	Data       map[string]interface{}
	Router     *Router
	cwd        string   // 项目当前所在目录
	prjName    string   // 所在目录项目名称
	Command    string   // 当前的命令
	SubCommand string   // 二级命令
	Setting    []string // 项目设置
	Queue      []string // 命令队列
}

/**
检测属性是否存在
*/
func (app App) HasSetting(set string) bool {
	has := false
	if idx := str.InQue(set, app.Setting); idx > -1 {
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

// 获取项目名称
func (app *App) PrjName() string {
	if app.prjName == "" {
		_, name := path.Split(app.Cwd())
		app.prjName = name
	}
	return app.prjName
}

// 队列参数新增
func (app *App) queueAppend(arg string) *App {
	app.Queue = append(app.Queue, arg)
	return app
}

// 队列右邻值
func (app *App) QueueNext(key string) string {
	idx := -1
	qLen := len(app.Queue)
	var vaule string
	for i := 0; i < qLen; i++ {
		if idx == i {
			vaule = app.Queue[i]
			break
		}
		if key == app.Queue[i] {
			idx = i + 1
		}
	}
	return vaule
}
