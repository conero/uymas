package bin

import (
	"fmt"
	"github.com/conero/uymas"
	"github.com/conero/uymas/bin/butil"
	"github.com/conero/uymas/bin/parser"
	"github.com/conero/uymas/str"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const (
	actionRunConstruct = "Construct"
)

// the cli application
type CLI struct {
	cmds map[string]interface{} // the register of commands.

	// the map the cmd, support many alias cmd about the cmd
	// the data struct like: {cmd => alias} or {cmd => [alias1, alias2, alias3...]}
	cmdMap               map[string]interface{}
	actionEmptyRegister  interface{} // the register callback by empty action.
	actionUnfindRegister interface{} // the register callback by command not handler
	commands             map[string]Cmd
	tempLastCommand      string                 // command Cache
	injectionData        map[string]interface{} //reject data from outside like chan control

	//external fields
	UnLoadDataSyntax bool //not support load data syntax, like json/url.
}

// the command of the cli application.
type CliCmd struct {
	Data       map[string]interface{} // the data from the `DataRaw` by parse for type
	DataRaw    map[string]string      // the cli application apply the data
	Command    string                 // the current command
	SubCommand string                 // the sub command
	Setting    []string               // the setting of command
	Raw        []string               // the raw args
	context    CLI
	cmdType    int //the command type enumeration
}

// the cli app.
type CliApp struct {
	Cc *CliCmd
}

// the interface of CliApp
type CliAppInterface interface {
	Construct()
}

type CmdOptions struct {
	Option   string
	Alias    interface{}
	Describe string
}

//define the struct command
type Cmd struct {
	Command  string
	Alias    interface{}           //string, []string. the alias of the command
	Describe string                //describe the command
	Handler  func(cc *CliCmd)      //when command call then handler the request
	Options  map[string]CmdOptions // the command option
}

// the construct of `CLI`
func NewCLI() *CLI {
	cli := &CLI{
		cmds:     map[string]interface{}{},
		cmdMap:   map[string]interface{}{},
		commands: map[string]Cmd{},
	}
	return cli
}

// the construct of `CliCmd`
func NewCliCmd(args ...string) *CliCmd {
	c := &CliCmd{
		Raw:     args,
		Setting: []string{},
		DataRaw: map[string]string{},
		Data:    map[string]interface{}{},
	}
	// parse the args
	c.parseArgs()
	return c
}

//@todo notice: `--test-string="Joshua 存在空格的字符串 Conero"` 解析失败
// construction of `CliCmd` by string
func NewCliCmdByString(ss string) *CliCmd {
	return NewCliCmd(butil.StringToArgs(ss)...)
}

//get the list cmd of application
func (cli *CLI) GetCmdList() []string {
	var list []string
	if cli.cmds != nil {
		list = []string{}
		for cmd, _ := range cli.cmds {
			if alias, hasMk := cli.cmdMap[cmd]; hasMk {
				var cmdList = []string{cmd}
				if v, isStr := alias.(string); isStr {
					cmdList = append(cmdList, v)
					cmd = strings.Join(cmdList, ", ")
				} else if v, isStrQue := alias.([]string); isStrQue {
					cmdList = append(cmdList, v...)
					cmd = strings.Join(cmdList, ", ")
				}
			}

			list = append(list, cmd)
		}
	}
	return list
}

//register functional command, the format like
//  `RegisterFunc(todo func(cc *CliCmd), cmd string)` or `RegisterFunc(todo func(), cmd, alias string)`
func (cli *CLI) RegisterFunc(todo func(cc *CliCmd), cmds ...string) *CLI {
	if len(cmds) > 0 {
		cmd := cmds[0]
		if len(cmds) > 1 {
			cli.cmdMap[cmd] = cmds[1]
		}
		// make the function map th struct
		cli.commands[cmd] = Cmd{
			Command: cmd,
			Alias:   cli.cmdMap[cmd],
		}
	}
	cli.registerFunc(todo, cmds...)
	return cli
}

func (cli *CLI) registerFunc(todo func(cc *CliCmd), cmds ...string) {
	cli.tempLastCommand = ""
	if len(cmds) > 0 {
		cmd := cmds[0]
		cli.tempLastCommand = cmd
		cli.cmds[cmd] = todo
		if len(cmds) > 1 {
			cli.cmdMap[cmd] = cmds[1]
		}
	} else {
		// if `cmds` is empty, then set `CLI.RegisterEmpty`
		if cli.actionEmptyRegister == nil {
			cli.actionEmptyRegister = todo
		} else {
			panic("CLI.RegisterFunc param `cmds` is empty will call `RegisterEmpty` that never be call before if" +
				" else fail register. ")
		}
	}
}

func (cli *CLI) Describe(desc string) bool {
	if cli.tempLastCommand != "" {
		if c, has := cli.commands[cli.tempLastCommand]; has {
			c.Describe = desc
			cli.commands[cli.tempLastCommand] = c
			return true
		}
	}
	return false
}

// support the `cmd, alias` param.
func (cli *CLI) GetDescribe(cmd string) string {
	if strings.Index(cmd, ",") > -1 {
		que := strings.Split(cmd, ",")
		cmd = strings.TrimSpace(que[0])
	}

	if v, has := cli.commands[cmd]; has {
		return v.Describe
	}

	return fmt.Sprintf("the command %v", cmd)
}

//register by command struct data
func (cli *CLI) RegisterCommand(c Cmd) *CLI {
	if c.Command != "" {
		cli.commands[c.Command] = c
		if c.Handler != nil {
			cmds := []string{c.Command}
			alias := c.Alias
			if alias != nil {
				if v, isStr := alias.(string); isStr {
					cmds = append(cmds, v)
				} else if v, isStrQue := alias.([]string); isStrQue {
					cmds = append(cmds, v...)
				}
			}
			cli.registerFunc(c.Handler, cmds...)
		}
	}
	return cli
}

//register the struct app, the format same as RegisterFunc. cmds any be `cmd string` or `cmd, alias string`
func (cli *CLI) RegisterApp(ap interface{}, cmds ...string) *CLI {
	if cmds != nil && len(cmds) > 0 {
		cmd := cmds[0]
		cli.cmds[cmd] = ap
		if len(cmds) > 1 {
			cli.cmdMap[cmd] = cmds[1]
		}
	}
	return cli
}

//Register many apps once.
func (cli *CLI) RegisterApps(aps map[string]interface{}) *CLI {
	for c, ap := range aps {
		cli.RegisterApp(ap, c)
	}
	return cli
}

// when the cmd is empty then callback the function, action only be
//   1. function `func(cc *CliCmd)`/`func()` or struct.
func (cli *CLI) RegisterEmpty(action interface{}) *CLI {
	cli.actionEmptyRegister = action
	return cli
}

// when command input not handler will callback the register, the format like:
//   1. function `func(cmd string, cc *CliCmd)`/`func(cmd string)`/`func(cc *CliCmd)`
func (cli *CLI) RegisterUnfind(action interface{}) *CLI {
	cli.actionUnfindRegister = action
	return cli
}

//the run the application
func (cli *CLI) Run(args ...string) {
	if args == nil {
		// if the args is empty then use the `os.Args`
		osArgs := os.Args
		if len(osArgs) > 1 {
			args = osArgs[1:]
		}
	}
	// construct of `CliCmd`
	cmd := NewCliCmd(args...)
	// start router by register.
	cli.router(cmd)
}

//call the application cmd
func (cli *CLI) CallCmd(cmd string) {
	cm := NewCliCmd(cmd)
	cli.router(cm)
}

//test cmd exist in application
func (cli *CLI) CmdExist(cmds ...string) bool {
	cmdExist := false
	for _, cmd := range cmds {
		_, exist := cli.cmdMap[cmd]
		if exist {
			cmdExist = true
			break
		}
	}

	if !cmdExist {
		for _, cm := range cli.cmdMap {
			//KV: string->string
			if cmStr, isStr := cm.(string); isStr && str.InQue(cmStr, cmds) > -1 {
				cmdExist = true
				break
			} else if cmStrQue, isStrArray := cm.([]string); isStrArray {
				for _, cStr := range cmds {
					if str.InQue(cStr, cmStrQue) > -1 {
						cmdExist = true
						break
					}
				}
			}
		}
	}
	return cmdExist
}

// @todo never stop to optimize the method.
// to star `router`
func (cli *CLI) router(cc *CliCmd) {
	//set the last `*CLI` as context of `CliCmd`.
	cc.context = *cli
	cc.cmdType = int(CmdFunc)
	routerValidMk := false
	if cc.Command != "" {
		value := cli.findRegisterValueByCommand(cc.Command)
		if value != nil {
			switch value.(type) {
			// call the FuncCmd
			case func(c *CliCmd):
				cli.hookBeforeCall(cc)
				value.(func(c *CliCmd))(cc)
				routerValidMk = true
			default:
				// call the AppCmd
				v := reflect.ValueOf(value).Elem()
				ccStr := "Cc"
				// set `Cc` that is struct of field.
				if v.FieldByName(ccStr).IsValid() {
					cc.cmdType = int(CmdApp)
					v.FieldByName(ccStr).Set(reflect.ValueOf(cc))
				} else {
					panic(fmt.Sprintf("%v:the command field of `Cc` is not valid filed.", cc.Command))
				}

				// any construct to call method need this before.
				v = reflect.ValueOf(value)
				// to call the construct action.
				if v.MethodByName(actionRunConstruct).IsValid() {
					v.MethodByName(actionRunConstruct).Call(nil)
				} else {
					panic(fmt.Sprintf("%v: the command is not vaild.", cc.Command))
				}

				//the subCommand string
				subCmdStr := cc.SubCommand
				if subCmdStr != "" {
					subCmdStr = strings.Title(subCmdStr)
					if v.MethodByName(subCmdStr).IsValid() {
						cli.hookBeforeCall(cc)
						v.MethodByName(subCmdStr).Call(nil)
					} else {
						panic(fmt.Sprintf("the method `%s` do not have a handler as `%v`.", cc.SubCommand, subCmdStr))
					}
				}

				routerValidMk = true
			}
		}

		// `unfind` handler
		if !routerValidMk {
			if cli.actionUnfindRegister != nil {
				aur := cli.actionUnfindRegister
				switch aur.(type) {
				case func(cmd string, cc *CliCmd):
					cli.hookBeforeCall(cc)
					aur.(func(cmd string, cc *CliCmd))(cc.Command, cc)
					routerValidMk = true
				case func(cmd string):
					cli.hookBeforeCall(cc)
					aur.(func(cmd string))(cc.Command)
					routerValidMk = true
				case func(cc *CliCmd):
					cli.hookBeforeCall(cc)
					aur.(func(cc *CliCmd))(cc)
					routerValidMk = true
				default:
					log.Printf("[WARNING] the method `RegisterUnfind` of param is valid, please reference the doc.")
				}
			}

			if !routerValidMk {
				fmt.Printf(" Fail: the command (%v) not find.\r\n", cc.Command)
				fmt.Printf("   Power from softwore %v, Version: %v/%v.\r\n\r\n", uymas.PkgName,
					uymas.Version, uymas.Release)
				routerValidMk = true
			}
		}
	}

	//empty calls.
	if !routerValidMk {
		if cli.actionEmptyRegister != nil {
			aer := cli.actionEmptyRegister
			switch aer.(type) {
			case func(cc *CliCmd):
				cli.hookBeforeCall(cc)
				aer.(func(cc *CliCmd))(cc)
				routerValidMk = true
			case func():
				cli.hookBeforeCall(cc)
				aer.(func())()
				routerValidMk = true
			}
		}
	}
}

// the hook before call the func
func (cli *CLI) hookBeforeCall(cc *CliCmd) {
	cli.loadDataSyntax(cc)
}

// the do load data by setting syntax
func (cli *CLI) loadDataSyntax(cc *CliCmd) {
	raw := cc.DataRaw
	if !cli.UnLoadDataSyntax && len(raw) > 0 {
		allowLoads := []string{
			"load-json", "LJ", "load-json-file", "LJF", "load-json-url", "LJU",
			"load-url", "LU", "load-url-file", "LUF", "load-url-url", "LUU",
		}
		for _, allow := range allowLoads {
			var (
				loadType    string
				contentType parser.DataReceiverType
			)
			if content, exist := raw[allow]; exist {
				switch allow {
				case "load-json", "LJ":
					loadType = parser.RJson
					contentType = parser.ReceiverContent
				case "load-json-file", "LJF":
					loadType = parser.RJson
					contentType = parser.ReceiverFile
				case "load-json-url", "LJU":
					loadType = parser.RJson
					contentType = parser.ReceiverUrl
				case "load-url", "LU":
					loadType = parser.RUrl
					contentType = parser.ReceiverContent
				case "load-url-file", "LUF":
					loadType = parser.RUrl
					contentType = parser.ReceiverFile
				case "load-url-url", "LUU":
					loadType = parser.RUrl
					contentType = parser.ReceiverUrl
				}

				if loadType != "" {
					dr, _ := parser.NewDataReceiver(loadType)
					dr.Receiver(contentType, content)
					cc.AppendData(dr.GetData())
				}
			}
		}
	}
}

// find the register value by command.
func (cli *CLI) findRegisterValueByCommand(c string) interface{} {
	var value interface{} = nil
	cmds := cli.cmds
	if v, has := cmds[c]; has {
		value = v
	} else if cli.cmdMap != nil {
		for aCmd, mV := range cli.cmdMap {
			isBreak := false
			switch mV.(type) {
			case string:
				if c == mV.(string) {
					if v, has := cmds[aCmd]; has {
						isBreak = true
						value = v
					}
				}
			case []string:
				for _, vs := range mV.([]string) {
					if c == vs {
						if v, has := cmds[aCmd]; has {
							isBreak = true
							value = v
						}
					}
				}
			}

			if isBreak {
				break
			}
		}
	}
	return value
}

//inject for data from outside.
func (cli *CLI) Inject(key string, value interface{}) *CLI {
	if cli.injectionData == nil {
		cli.injectionData = map[string]interface{}{}
	}
	cli.injectionData[key] = value
	return cli
}

//get Injection data
func (cli *CLI) GetInjection(key string) interface{} {
	if cli.injectionData == nil {
		return nil
	}
	value, has := cli.injectionData[key]
	if has {
		return value
	}
	return nil
}

/*****  methods of the `CliCmd` ***/
// 检测属性是否存在
func (app *CliCmd) HasSetting(set string) bool {
	has := false
	if idx := str.InQue(set, app.Setting); idx > -1 {
		has = true
	}
	return has
}

// 检测设置是否存在，支持多个
func (app *CliCmd) CheckSetting(sets ...string) bool {
	has := false
	for _, set := range sets {
		if idx := str.InQue(set, app.Setting); idx > -1 {
			has = true
			break
		}
	}
	return has
}

//检测必须要的参数值
func (app *CliCmd) CheckMustKey(keys ...string) bool {
	check := true
	for _, k := range keys {
		if v, has := app.DataRaw[k]; !has || v == "" {
			check = false
			break
		}
	}
	return check
}

// 获取当的工作目录
func (app *CliCmd) Cwd() string {
	return butil.GetBasedir()
}

// 队列右邻值
func (app *CliCmd) QueueNext(key string) string {
	idx := -1
	qLen := len(app.Raw)
	var vaule string
	for i := 0; i < qLen; i++ {
		if idx == i {
			vaule = app.Raw[i]
			break
		}
		if key == app.Raw[i] {
			idx = i + 1
		}
	}
	return vaule
}

// Get key values from multiple key values
func (app *CliCmd) Next(keys ...string) string {
	var value string
	var vLen = len(keys)
	//when keys is empty default use the current Next value
	if vLen == 0 {
		return app.Next(app.SubCommand)
	}
	for _, k := range keys {
		value = app.QueueNext(k)
		if value != "" {
			break
		}
	}
	return value
}

// get raw args data, because some args has alias list.
func (app *CliCmd) ArgRaw(keys ...string) string {
	var value string
	for _, key := range keys {
		if v, b := app.DataRaw[key]; b {
			value = v
			break
		}
	}

	return value
}

//get args data see as int
func (app *CliCmd) ArgInt(keys ...string) int {
	value := app.ArgRaw(keys...)
	if "" == value {
		return 0
	}
	if num, err := strconv.Atoi(value); err == nil {
		return num
	}
	return 0
}

// get raw arg has default
func (app *CliCmd) ArgRawDefault(key, def string) string {
	var value = def
	if v, b := app.DataRaw[key]; b {
		value = v
	}
	return value
}

// get arg after parsed the raw data
func (app *CliCmd) Arg(keys ...string) interface{} {
	var value interface{} = nil
	for _, key := range keys {
		if v, b := app.Data[key]; b {
			value = v
			break
		}
	}
	return value
}

// can default value to get the arg
func (app *CliCmd) ArgDefault(key string, def interface{}) interface{} {
	var value = def
	if v, b := app.Data[key]; b {
		value = v
	}
	return value
}

//get the raw line input.
func (app *CliCmd) ArgRawLine() string {
	return strings.Join(app.Raw, " ")
}

//call cmd
func (app *CliCmd) CallCmd(cmd string) {
	context := app.context
	if context.CmdExist(cmd) && app.Command != cmd {
		app.Command = cmd
		context.router(app)
	}
}

// get the context of `CLI`, in case `AppCmd` not `FunctionCmd`
func (app *CliCmd) Context() CLI {
	return app.context
}

func (app *CliCmd) CmdType() int {
	return app.cmdType
}

// append the Data
func (app *CliCmd) AppendData(vMap map[string]interface{}) *CliCmd {
	if len(vMap) > 0 {
		if app.Data == nil {
			app.Data = map[string]interface{}{}
		}
		if app.DataRaw == nil {
			app.DataRaw = map[string]string{}
		}
		for k, v := range vMap {
			var value string
			if v != nil {
				value = fmt.Sprintf("%v", v)
			}
			app.Data[k] = v
			app.DataRaw[k] = value
		}
	}
	return app
}

// the application parse raw args inner.
//
// the command format like that:
// 		1. `$ [command] [option]`
// 		2. `$ [command] [sub_command]`
// 		3. `$ [option]`
//
// the option format example:
//		`[command] -xyz` same like `[command] -x -y -z`
//		`[command] --version --name 'Joshua Conero'`
//		`[command] --list A B C D -L A B C D`
//		`[command] --name='Joshua Conero'`
//@todo app.Data --> 类型解析太简陋；支持类型与 Readme.md 不统一
func (app *CliCmd) parseArgs() {
	if app.Raw != nil {
		optKeyList := []string{}
		optKey := ""
		for i, arg := range app.Raw {
			if i == 0 && isVaildCmd(arg) {
				app.Command = arg
			} else if i == 1 && app.Command != "" && isVaildCmd(arg) {
				app.SubCommand = arg
			} else {
				strLen := len(arg)
				markKeySuccess := false
				if strLen > 1 {
					if strLen > 1 && "--" == arg[0:2] {
						arg = arg[2:]
						idx := strings.Index(arg, "=")
						if idx == -1 { // --key
							optKey = arg
							app.Setting = append(app.Setting, arg)
						} else { // --key=value
							optKey = ""
							tmpKey := arg[0:idx]
							tmpValue := arg[idx+1:]
							app.DataRaw[tmpKey] = tmpValue
						}
						markKeySuccess = true
					} else if "-" == arg[0:1] {
						arg = arg[1:]
						// -x
						if len(arg) == 1 {
							optKey = arg
							app.Setting = append(app.Setting, arg)
						} else { // -xyz => -x -y -z
							tmpArr := strings.Split(arg, "")
							optKey = ""
							for _, vs := range tmpArr {
								app.Setting = append(app.Setting, vs)
							}
						}
						markKeySuccess = true
					}
				}

				if !markKeySuccess && optKey != "" {
					arg = CleanoutString(arg)
					if ddVal, ddHas := app.Data[optKey]; ddHas {
						switch ddVal.(type) {
						case string:
							oldSs := app.Data[optKey].(string)
							app.Data[optKey] = []string{oldSs, arg}
						case []string:
							oldVarr := app.Data[optKey].([]string)
							oldVarr = append(oldVarr, arg)
							app.Data[optKey] = oldVarr
						}
					} else {
						app.Data[optKey] = arg
					}
				}
			}

			if optKey != "" && -1 == str.InQue(optKey, optKeyList) {
				optKeyList = append(optKeyList, optKey)
			}
		}

		//`app.Data` => `app.DataRaw`
		for _, k := range optKeyList {
			if dV, kHas := app.Data[k]; kHas {
				if _, rKhas := app.DataRaw[k]; !rKhas {
					switch dV.(type) {
					case []string:
						app.DataRaw[k] = strings.Join(dV.([]string), " ")
					case string:
						app.DataRaw[k] = dV.(string)
					}
				}
			}
		}

		//parse the raw cmd data to data.
		//`app.DataRaw` => `app.Data`
		for rawKey, rawVal := range app.DataRaw {
			if _, hasRawKey := app.Data[rawKey]; !hasRawKey {
				app.Data[rawKey] = ParseValueByStr(rawVal)
			}
		}
	}
}

// check the cmd of validation
func isVaildCmd(c string) bool {
	if len(c) == 0 || c[0:1] == "-" {
		return false
	}
	return true
}

/*
clear out the raw input string like:
	`"string"`		=> `string`
	`"'string'"`	=> `'string'`
	`'string'`		=> `string`
	`'"string"'`	=> `"string"`
*/
func CleanoutString(ss string) string {
	ssLen := len(ss)
	first, last := ss[0:1], ss[ssLen-1:]
	if first == last {
		if first == "'" || last == `"` {
			ss = ss[1 : ssLen-1]
		}
	}

	return ss
}

//parse the command value to really type by format.
func ParseValueByStr(ss string) interface{} {
	ss = strings.TrimSpace(ss)
	ssLow := strings.ToLower(ss)

	//bool
	if ssLow == "true" || ssLow == "false" {
		if ssLow == "true" {
			return true
		} else {
			return false
		}
	}

	//int64
	if i64, er := strconv.ParseInt(ss, 10, 64); er == nil {
		return i64
	}

	//float64
	if f64, er := strconv.ParseFloat(ss, 64); er == nil {
		return f64
	}

	return ss
}
