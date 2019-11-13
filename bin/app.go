package bin

import (
	"github.com/conero/uymas/str"
	"path"
)

// @Date：   2018/10/30 0030 12:40
// @Author:  Joshua Conero
// @Name:    应用管理

// 全局命令行实例，可通过 GetApp 获取
type App struct {
	Data       map[string]interface{} // 格式化数据
	DataRaw    map[string]string      // 原始数据，命令行的所有数据解析时都为字符串
	Router     *Router
	cwd        string   // 项目当前所在目录
	prjName    string   // 所在目录项目名称
	Command    string   // 当前的命令
	SubCommand string   // 二级命令
	Setting    []string // 项目设置
	Queue      []string // 命令队列
}

// 检测属性是否存在
func (app *App) HasSetting(set string) bool {
	has := false
	if idx := str.InQue(set, app.Setting); idx > -1 {
		has = true
	}
	return has
}

// 检测设置是否存在，支持多个
func (app *App) CheckSetting(sets ...string) bool {
	has := false
	for _, set := range sets {
		if idx := str.InQue(set, app.Setting); idx > -1 {
			has = true
			break
		}
	}
	return has
}

// 获取当的工作目录
func (app *App) Cwd() string {
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

// 清空队列
func (app *App) resetQueue() {
	app.Queue = []string{}
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

// 多简直获取键值
func (app *App) Next(keys ...string) string {
	var value string
	for _, k := range keys {
		value = app.QueueNext(k)
		if value != "" {
			break
		}
	}
	return value
}

// get raw arg data
func (app *App) ArgRaw(key string) string {
	var value string
	if v, b := app.DataRaw[key]; b {
		value = v
	}
	return value
}

// get raw arg has default
func (app *App) ArgRawDefault(key, def string) string {
	var value = def
	if v, b := app.DataRaw[key]; b {
		value = v
	}
	return value
}

// get arg after parsed the raw data
func (app *App) Arg(key string) interface{} {
	var value interface{} = nil
	if v, b := app.Data[key]; b {
		value = v
	}
	return value
}

// can default value to get the arg
func (app *App) ArgDefault(key string, def interface{}) interface{} {
	var value = def
	if v, b := app.Data[key]; b {
		value = v
	}
	return value
}
