package bin

import "github.com/conero/uymas/util"

// @Date：   2018/10/30 0030 12:40
// @Author:  Joshua Conero
// @Name:    应用管理

type App struct {
	Data    map[string]string
	Options []string
	Router  *Router
	cwd	string			// 项目当前所在目录
	Command string		// 当前的命令
}

/**
检测属性是否存在
*/
func (app App) HasOptions(opt string) bool {
	has := false
	if idx := util.InStrQue(opt, app.Options); idx > -1 {
		has = true
	}
	return has
}

/**
	获取当的工作目录
 */
func (app App) Cwd () string{
	return app.cwd
}