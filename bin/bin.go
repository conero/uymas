// 轻量级命令行框架
package bin

import (
	"fmt"
	"github.com/conero/uymas"
	"github.com/conero/uymas/number"
	"github.com/conero/uymas/str"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"
)

// @Date：   2018/10/30 0030 13:20
// @Author:  Joshua Conero
// @Name:    库入口文件

var app *App = nil
var args []string = nil
var defaultRouter *Router
var routerCmdApp map[string]interface{} = nil
var routerAliasApp map[string][]string = nil  // 项目别名匹配
var subCommandAble bool = true                // 二级命令有效
var appRuningWorkDir string                   // 应用运行目录
var appFuncRouterMap map[string]func() = nil  // 函数式路由地址字典
var _funcStyleEmptyTodo func() = nil          // 空函数命令使用
var _funcStyleUnfindTo func(cmd string) = nil // 命令未知

const (
	AppMethodInit   = "Init"
	AppMethodRun    = "Run"
	AppMethodNoSubC = "SubCommandUnfind"
	AppMethodHelp   = "Help"
)

// 建议使用 InjectArgs 代替，后期可能进行优化
// 命令程序初始化入口，用于开发时非直接编译测试
func Init(param []string) {
	args = param
}

// 此方法用于非 os.Args 测试
// 注入参数
func InjectArgs(params ...string) {
	args = getArgs()
	newArgs := []string{""}
	if len(args) == 0 {
		newArgs[0] = args[0]
	}

	// 替换系统测试
	newArgs = append(newArgs, params...)
	args = newArgs
}

// 获取输入的参数
func getArgs() []string {
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
func RegisterFunc(cmd string, todo func()) {
	appFuncRouterMap[cmd] = todo
}

// 空函数命令注册
func EmptyFunc(todo func()) {
	_funcStyleEmptyTodo = todo
}

// 路由失败时的函数
func UnfindFunc(todo func(cmd string)) {
	_funcStyleUnfindTo = todo
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
	appFuncRouterMap = map[string]func(){}

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

// 获取字符串格式化
// [[k,  v]]
func FormatStr(d string, ss ...[][]string) string {
	if d == "" {
		// 4 个空格
		d = "   "
	}
	bit := d[0:1]

	// 或者最大长度
	maxLen := 0
	for _, sg := range ss {
		for _, s := range sg {
			kLen := len(s[0])
			if kLen > maxLen {
				maxLen = kLen
			}
		}
	}

	maxLen += len(d)

	// 格式化
	var contents string
	for _, sg := range ss {
		for _, s := range sg {
			ss1 := s[0] + strings.Repeat(bit, maxLen-len(s[0])) + s[1] + "\n"
			contents += ss1
		}
	}

	return contents
}

// 格式化数组字符
// 用于命令行输出
// prefs 为 "" 时默认以数组索引开头；否则默给定的输出
func FormatQue(que []interface{}, prefs ...string) string {
	pref := ""  // 开头符号
	dter := " " // 空格
	if prefs != nil && len(prefs) > 0 {
		pref = prefs[0]
		if len(prefs) > 1 {
			dter = prefs[1]
		}
	}
	s := ""
	queLen := len(que)
	mdLen := 4 + len(strconv.Itoa(queLen))
	for i, q := range que {
		if pref == "" {
			iStr := strconv.Itoa(i) + "."
			s += iStr + strings.Repeat(dter, mdLen-len(iStr)) + fmt.Sprintf(" %v\n", q)
		} else {
			s += pref + strings.Repeat(dter, mdLen-len(pref)) + fmt.Sprintf(" %v\n", q)
		}
	}
	return s
}

// 表格格式化
// (data, bool) 是否使用 idx
func FormatTable(data [][]interface{}, args ...interface{}) string {
	useIdxMk := true
	if args != nil {
		if v, isBool := args[0].(bool); isBool {
			useIdxMk = v
		}
	}

	// 数据处理
	data2Str := [][]string{}
	maxLenQue := []int{}
	for _, dd := range data {
		ddStr := []string{}
		for i, d := range dd {
			ddStr = append(ddStr, fmt.Sprintf("%v", d))
			ddStrLen := len(ddStr)
			if len(maxLenQue) > i {
				if maxLenQue[i] < ddStrLen {
					maxLenQue[i] = ddStrLen
				}
			} else {
				maxLenQue = append(maxLenQue, ddStrLen)
			}
		}
		data2Str = append(data2Str, ddStr)
	}

	var s string
	dCtt := len(data)
	maxLen := number.SumQInt(maxLenQue) + dCtt*2
	if useIdxMk {
		dCttLen := len(strconv.Itoa(dCtt) + ".")
		maxLen += dCttLen + dCtt*2
		maxLenQue = append([]int{dCttLen}, maxLenQue...)
	} else {
		maxLen += (dCtt - 1) * 2
	}

	for j, sdd := range data2Str {
		line := ""
		tLen := maxLen
		if useIdxMk {
			jStr := strconv.Itoa(j + 1)
			tLen -= tLen
			jStr = str.PadRight(jStr, " ", maxLenQue[0]+2)
			s += jStr
		}
		for i, sd := range sdd {
			maxCol := maxLenQue[i]
			if useIdxMk {
				maxCol = maxLenQue[i+1]
			}
			s += str.PadRight(sd, " ", maxCol+2)
		}
		s += line + "\n"
	}
	return s
}

// 空命令检测
func IsEmptyCmd() bool {
	args = getArgs()
	return 1 == len(args)
}
