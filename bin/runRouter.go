package bin

import (
	"github.com/conero/uymas/util/str"
	"reflect"
)

// @Date：   2018/10/30 0030 15:11
// @Author:  Joshua Conero
// @Name:    启动路由


// 启动 app 路由器
func runAppRouter()  {
	// 路由配置
	router := app.Router
	if router == nil{
		router = defaultRouter
	}
	app.Command = ""
	for i, arg := range getArgs(){
		if i == 0{
			continue
		}
		if app.Command == ""{
			app.Command = arg
		}
	}


	// 路由匹配
	if app.Command == ""{
		if router != nil && router.EmptyAction != nil{
			app.Router.EmptyAction()
		}else{
			defaultRouter.EmptyAction()
		}
	}else{
		command := str.Ucfirst(app.Command)
		v := reflect.ValueOf(command).Elem()
		v.MethodByName("Init").Call(nil)
	}
}
