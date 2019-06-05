package bin

import (
	"github.com/conero/uymas/str"
	"reflect"
	"regexp"
	"strconv"
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

// app.Data 数据格式化
// 正则检查并被处理为对应的数据格式
// --key='字符串参数    可含空格'/--key="'字符串参数    可含空格'"     去除标点，并且认为是字符串
// --key=1,2,45,87,96,52,37 整形数组
// 解析的格式有: bool, int64, float64, string, []int, []string,[]float64, string
func StrParseData(v string) interface{} {
	v = strings.TrimSpace(v)
	var newV interface{} = nil
	if v != "" {
		isParseMk := false
		// bool 布尔类型检测
		if str.InQue(strings.ToLower(v), []string{"true", "false"}) > -1 {
			if s, err := strconv.ParseBool(v); err == nil {
				newV = s
				isParseMk = true
			}
		}

		// 解析处理
		if isParseMk { // 前期解析空操作
		} else if matched, err := regexp.MatchString(`^['"]+.*['"]+$`, v); err == nil && matched {
			// --key='s1 s2 s3'
			// --key="'s1 s2 s3'"
			newV = v[1 : len(v)-1]
		} else {
			if s, err := strconv.ParseInt(v, 10, 64); err == nil { // int64
				newV = s
			} else if s, err := strconv.ParseFloat(v, 64); err == nil { // float64
				newV = s
			} else {
				if strings.Index(v, ",") > -1 {
					if strings.Index(v, ".") > -1 {
						if matched, err := regexp.MatchString(`^[1-9.]+[0-9,.]+[0-9.]+$`, v); err == nil && matched {
							// []float64
							// --key=1,3.14,45,87,96,52,37		合法
							// --key=1,2,45,87,96,52,37,ddd		非法
							// --key=1,2,3,4,5					非法
							f64s := []float64{}
							for _, i := range strings.Split(v, ",") {
								if s, err := strconv.ParseFloat(i, 64); err == nil {
									f64s = append(f64s, s)
								}
							}
							isParseMk = true
							newV = f64s
						}
					} else if matched, err := regexp.MatchString(`^[1-9]+[0-9,]+[0-9]+$`, v); err == nil && matched {
						// []int
						// --key=1,2,45,87,96,52,37			合法
						// --key=1,2,45,87,96,52,37,ddd		非法
						ints := []int{}
						for _, i := range strings.Split(v, ",") {
							if iv, err := strconv.Atoi(i); err == nil {
								ints = append(ints, iv)
							}
						}
						isParseMk = true
						newV = ints
					} else {
						// --key=teee,hhhh\,,fgfg
						tmpS := "-._._.-"
						vs := strings.Replace(v, "\\,", tmpS, -1)
						if strings.Index(vs, ",") > -1 {
							newSs := []string{}
							for _, st := range strings.Split(vs, ",") {
								st = strings.Replace(st, tmpS, ",", -1)
								newSs = append(newSs, st)
							}
							newV = newSs
							isParseMk = true
						}
					}
				}

				if !isParseMk {
					newV = v
				}
			}
		}
		//fmt.Printf("%T, %v \r\n", newV, newV)
	}

	return newV
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
				app.Data[k] = StrParseData(v)
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
