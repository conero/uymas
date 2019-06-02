package bin

import (
	"github.com/conero/uymas/str"
	"reflect"
	"strings"
)

// @Date：   2018/10/30 0030 15:11
// @Author:  Joshua Conero
// @Name:    启动路由

// 是否为有效的命令
func isVaildCmd(c string) bool {
	if len(c) == 0 || c[0:1] == "-" {
		return false
	}
	return true
}

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
		if app.Command == "" && isVaildCmd(arg) {
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
				v := arg[equalIdx+1:]
				app.Data[k] = v
				app.DataRaw[k] = v
			} else {
				app.Data[arg] = true
				app.Setting = append(app.Setting, arg)
			}
			app.queueAppend(arg)
		} else {
			app.queueAppend(arg)
			// 二级命令
			if app.SubCommand == "" && i == cmdIdx {
				app.SubCommand = arg
				continue
			}
		}
	}
	// 选项监听
	runOptLsnFn := func() bool {
		if app.Router != nil && app.Router.OptionListener != nil {
			for _, set := range app.Setting {
				if app.Router.OptionListener(set, app) {
					return true
				}
			}
		}
		return false
	}
	// 路由匹配
	if app.Command == "" {
		// 空命令
		if !runOptLsnFn() {
			if app.Router != nil && app.Router.EmptyAction != nil {
				app.Router.EmptyAction()
			} else if _funcStyleEmptyTodo != nil {
				_funcStyleEmptyTodo()
			} else {
				defaultRouter.EmptyAction()
			}
		}
	} else {
		command := getCommandByAlias(app.Command)
		if cmd, has := routerCmdApp[command]; has {
			v := reflect.ValueOf(cmd)
			v.MethodByName(AppMethodInit).Call(nil)
			if subCommandAble && app.SubCommand != "" {
				subC := str.Ucfirst(AmendSubC(app.SubCommand))
				if v.MethodByName(subC).IsValid() {
					v.MethodByName(subC).Call(nil)
				} else {
					v.MethodByName(AppMethodNoSubC).
						Call([]reflect.Value{reflect.ValueOf(app.SubCommand)})
				}
			} else {
				v.MethodByName(AppMethodRun).Call(nil)
			}
		} else {
			if !runOptLsnFn() {
				// 函数式简单路由
				hasRouter := app.Router != nil
				funcRouterMk := false
				if hasRouter && app.Router.FuncAction != nil {
					funcRouterMk = app.Router.FuncAction(command, app)
				}

				// 自定义函数路由
				if !funcRouterMk {
					for cs, cfunc := range appFuncRouterMap {
						if cs == command {
							cfunc()
							funcRouterMk = true
						}
					}
				}

				// 自定义函数式路由标记
				if !funcRouterMk {
					// 未知命令
					if hasRouter && app.Router.UnfindAction != nil {
						app.Router.UnfindAction(command)
					} else if _funcStyleUnfindTo != nil {
						_funcStyleUnfindTo(command)
					} else {
						defaultRouter.UnfindAction(command)
					}
				}
			}
		}
	}
}

/*
支持别名
*/
func getCommandByAlias(command string) string {
	for nCmd, queStr := range routerAliasApp {
		if idx := str.InQue(command, queStr); idx > -1 {
			command = nCmd
			break
		}
	}
	return command
}

// 修正文件命令
func AmendSubC(subC string) string {
	if strings.Index(subC, "-") > -1 {
		subC = strings.Replace(subC, "-", " ", -1)
		subC = str.Ucfirst(subC)
	}
	return subC
}
