package bin

import (
	"fmt"
	"github.com/conero/uymas"
	"os"
	"reflect"
	"strings"
)

// @Date：   2018/10/30 0030 13:20
// @Author:  Joshua Conero
// @Name:    库入口文件

var app *App = nil
var args []string = nil
var defaultRouter *Router
var routerCmdApp map[string]interface{} = nil
var routerAliasApp map[string][]string = nil // 项目别名匹配
var subCommandAble bool = true               // 二级命令有效

const (
	AppMethodInit   = "Init"
	AppMethodRun    = "Run"
	AppMethodNoSubC = "SubCommandUnfind"
	AppMethodHelp   = "Help"
)

/**
初始化
*/
func Init(param []string) {
	args = param
}

/**
获取输入的参数
*/
func getArgs() []string {
	if args == nil {
		args = os.Args
	}
	return args
}

/**
命令别名集(单个)
*/
func Alias(cmd string, alias ...string) {
	routerAliasApp[cmd] = alias
}

/**
命令别名集(多个)
*/
func AliasMany(alias map[string][]string) {
	for cmd, als := range alias {
		Alias(cmd, als...)
	}
}

/**
项目注册(单个)
*/
func Register(name string, cmd interface{}) {
	routerCmdApp[name] = cmd
}

/**
注册多个项目
*/
func RegisterApps(data map[string]interface{}) {
	for n, c := range data {
		Register(n, c)
	}
}

// 请求命令行帮助
func CallCmdHelp(key string) bool {
	a, has := routerCmdApp[key]
	key = strings.ToLower(key)
	if !has {
		for k, a1 := range routerCmdApp {
			if key == strings.ToLower(k) {
				a = a1
				has = true
				break
			}
		}
	}
	if has {
		ins := reflect.ValueOf(a).MethodByName(AppMethodHelp)
		if ins.IsValid() {
			ins.Call(nil)
			return true
		}
	}
	return false
}

// 加载路由器为适配器
func Adapter(router *Router) {
	app.Router = router
}

// 二级命令配置
func SubCommand(able bool) {
	subCommandAble = able
}

// 系统运行
func Run() App {
	runAppRouter()
	return *app
}

// 获取命令行 App
func GetApp() App {
	return *app
}

// 引用初始化
func init() {
	routerCmdApp = map[string]interface{}{}
	routerAliasApp = map[string][]string{}
	app = &App{
		Data:    map[string]interface{}{},
		Setting: []string{},
		Queue:   []string{},
	}

	// 默认路由，用于设置路由为空时
	defaultRouter = &Router{
		UnfindAction: func(action string) {
			fmt.Println("	欢迎使用 uymas包:" + uymas.Version + "/" + uymas.Release)
			fmt.Println("	" + action + " 命令不存在")
		},
		EmptyAction: func() {
			fmt.Println("	欢迎使用 uymas包:" + uymas.Version + "/" + uymas.Release)
		},
	}
	// 生成当前的工作目录
	if cwd, err := os.Getwd(); err == nil {
		app.cwd = cwd
	}

}
