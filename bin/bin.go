package bin

import (
	"fmt"
	"github.com/conero/uymas"
	"os"
)

// @Date：   2018/10/30 0030 13:20
// @Author:  Joshua Conero
// @Name:    库入口文件

var app *App = nil
var args []string = nil
var defaultRouter *Router

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

// 加载路由器为适配器
func Adapter(router *Router) {
	app.Router = router
}

// 系统运行
func Run() App {
	runAppRouter()
	return *app
}

// 引用初始化
func init() {
	app = &App{}

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
	if cwd, err := os.Getwd(); err == nil{
		app.cwd = cwd
	}
}
