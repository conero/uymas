package bin

import (
	"github.com/conero/uymas/util"
	"github.com/conero/uymas/util/str"
	"reflect"
	"strings"
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
	cmdIdx := -1
	for i, arg := range getArgs() {
		if i == 0 {
			continue
		}
		argLen := len(arg)
		if 0 == argLen {
			continue
		}
		// 命令解析
		if app.Command == "" {
			app.Command = arg
			app.queueAppend(arg)
			cmdIdx = i + 1
			continue
		}
		// 参数处理
		if "-" == arg[0:1] {
			if argLen > 2 && "--" == arg[0:2] {
				arg = arg[2:]
			} else {
				arg = arg[1:]
			}
			equalIdx := strings.Index(arg, "=")
			if equalIdx > -1 {
				k := arg[0:equalIdx]
				v := arg[equalIdx:]
				app.Data[k] = v
			} else {
				app.Data[arg] = true
				app.Setting = append(app.Setting, arg)
			}
			app.queueAppend(arg)
		} else {
			app.queueAppend(arg)
			// 二级命令
			if app.SubCommand == "" && i == cmdIdx{
				app.SubCommand = arg
				continue
			}
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
			if subCommandAble && app.SubCommand != "" {
				if v.MethodByName(str.Ucfirst(app.SubCommand)).IsValid() {
					v.MethodByName(str.Ucfirst(app.SubCommand)).Call(nil)
				} else {
					v.MethodByName(AppMethodNoSubC).
						Call([]reflect.Value{reflect.ValueOf(app.SubCommand)})
				}
			} else {
				v.MethodByName(AppMethodRun).Call(nil)
			}
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
