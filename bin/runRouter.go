package bin

import (
	"github.com/conero/uymas/util"
	"reflect"
)

// @Date：   2018/10/30 0030 15:11
// @Author:  Joshua Conero
// @Name:    启动路由

// 启动 app 路由器
func runAppRouter() {
	// 路由配置
	router := app.Router
	if router == nil {
		router = defaultRouter
	}
	app.Command = ""
	for i, arg := range getArgs() {
		if i == 0 {
			continue
		}
		if app.Command == "" {
			app.Command = arg
		}
	}

	// 路由匹配
	if app.Command == "" {
		if app.Router != nil && app.Router.EmptyAction != nil {
			app.Router.EmptyAction()
		} else {
			defaultRouter.EmptyAction()
		}
	} else {
		command := getCommandByAlias(app.Command)
		if cmd, has := routerCmdApp[command]; has {
			v := reflect.ValueOf(cmd)
			v.MethodByName(AppMethodInit).Call(nil)
			v.MethodByName(AppMethodRun).Call(nil)
		} else {
			if app.Router != nil && app.Router.UnfindAction != nil {
				app.Router.UnfindAction(command)
			} else {
				defaultRouter.UnfindAction(command)
			}
		}
	}
}

/**
支持别名
*/
func getCommandByAlias(command string) string {
	for nCmd, queStr := range routerAliasApp {
		if idx := util.InStrQue(command, queStr); idx > -1 {
			command = nCmd
			break
		}
	}
	return command
}
