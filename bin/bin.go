// 轻量级命令行框架
package bin

import (
	"fmt"
	"github.com/conero/uymas"
	"os"
	"path"
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
var routerAliasApp map[string][]string = nil          // 项目别名匹配
var subCommandAble bool = true                        // 二级命令有效
var appRuningWorkDir string                           // 应用运行目录
var appFuncRouterMap map[string]func(a *App) = nil    // 函数式路由地址字典
var _funcStyleEmptyTodo func(a *App) = nil            // 空函数命令使用
var _funcStyleUnfindTo func(a *App, cmd string) = nil // 命令未知
var _funNewFRdataDick map[string]FRdata = nil         // 新的函数命名

const (
	AppMethodInit     = "Init"
	AppMethodRun      = "Run"
	AppMethodNoSubC   = "SubCommandUnfind"
	AppMethodHelp     = "Help"
	FuncRegisterEmpty = "_inner_empty_func"
)

//函数式注册数据
type FRdata struct {
	Cmd     string
	Alias   []string // 函数别名
	Todo    func(a *App)
	OptDick OptionDick
	Opts    []Option
}

// 建议使用 InjectArgs 代替，后期可能进行优化
// 命令程序初始化入口，用于开发时非直接编译测试
// @todo 0.6 删除
// Deprecated: use InjectArgs instead, will delete in 0.6
func Init(param []string) {
	args = param
}

// 此方法用于非 os.Args 测试
// 注入参数
func InjectArgs(params ...string) {
	args = Args()
	newArgs := []string{""}
	if len(args) == 0 {
		newArgs[0] = args[0]
	}

	// 替换系统测试
	newArgs = append(newArgs, params...)
	args = newArgs
}

// 获取输入的参数
func Args() []string {
	if args == nil {
		args = os.Args
	}
	return args
}

// 命令别名集(单个)
func Alias(cmd string, alias ...string) {
	routerAliasApp[cmd] = alias
}

// 命令别名集(多个)
func AliasMany(alias map[string][]string) {
	for cmd, als := range alias {
		Alias(cmd, als...)
	}
}

// 项目注册(单个)
func Register(name string, cmd interface{}) {
	routerCmdApp[name] = cmd
}

// 注册多个项目
func RegisterApps(data map[string]interface{}) {
	for n, c := range data {
		Register(n, c)
	}
}

// 自定义函数式注册
func RegisterFunc(cmd string, todo func(a *App)) {
	appFuncRouterMap[cmd] = todo
}

// 空函数命令注册
func EmptyFunc(todo func(a *App)) {
	_funcStyleEmptyTodo = todo
}

// 路由失败时的函数
func UnfindFunc(todo func(a *App, cmd string)) {
	_funcStyleUnfindTo = todo
}

//新的函数式注册
func NewFRdata(dd []FRdata) {
	if dd != nil {
		if _funNewFRdataDick == nil {
			_funNewFRdataDick = map[string]FRdata{}
		}
		for _, fr := range dd {
			if fr.Cmd == FuncRegisterEmpty {
				EmptyFunc(fr.Todo)
			} else {
				// 函数注册
				RegisterFunc(fr.Cmd, fr.Todo)
			}
			_funNewFRdataDick[fr.Cmd] = fr
		}
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
func Run() *App {
	runAppRouter()
	return app
}

// 获取命令行 App
func GetApp() *App {
	return app
}

// 引用初始化
func init() {
	routerCmdApp = map[string]interface{}{}
	routerAliasApp = map[string][]string{}
	appFuncRouterMap = map[string]func(a *App){}

	app = &App{
		Data:    map[string]interface{}{},
		DataRaw: map[string]string{},
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
		FuncAction: nil,
	}
	// 生成当前的工作目录
	if cwd, err := os.Getwd(); err == nil {
		app.cwd = cwd
	}

}

// 获取命令正在运行的代码
func Rwd() string {
	if appRuningWorkDir == "" {
		rwd := os.Args[0]
		rwd = strings.Replace(rwd, "\\", "/", -1)
		rwd = path.Dir(rwd)
		appRuningWorkDir = rwd
	}
	return appRuningWorkDir
}

// 空命令检测
func IsEmptyCmd() bool {
	args = Args()
	return 1 == len(args)
}
