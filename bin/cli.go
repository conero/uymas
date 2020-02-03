package bin

import (
	"fmt"
	"github.com/conero/uymas"
	"github.com/conero/uymas/bin/butil"
	"github.com/conero/uymas/str"
	"os"
	"strings"
)

// the cli application
type CLI struct {
	cmds                 map[string]interface{} // the register of commands.
	cmdMap               map[string]interface{} // the map the cmd, support many alias cmd about the cmd
	actionEmptyRegister  interface{}            // the register callback by empty action.
	actionUnfindRegister interface{}            // the register callback by command not handler
}

// the command of the cli application.
type CliCmd struct {
	Data       map[string]interface{} // the data from the `DataRaw` by parse for type
	DataRaw    map[string]string      // the cli application apply the data
	Router     *Router
	Command    string   // the current command
	SubCommand string   // the sub command
	Setting    []string // the setting of command
	Raw        []string // the raw args
}

// the construct of `CLI`
func NewCLI() *CLI {
	cli := &CLI{
		cmds:   map[string]interface{}{},
		cmdMap: map[string]interface{}{},
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
			list = append(list, cmd)
		}
	}
	return list
}

//register functional command, the format like
//  `RegisterFunc(todo func(cc *CliCmd), cmd string)` or `RegisterFunc(todo func(), cmd, alias string)`
func (cli *CLI) RegisterFunc(todo func(cc *CliCmd), cmds ...string) *CLI {
	if cmds != nil && len(cmds) > 0 {
		cmd := cmds[0]
		cli.cmds[cmd] = todo
		if len(cmds) > 1 {
			cli.cmdMap[cmd] = cmds[1]
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

// when the cmd is empty then callback the function, action only be
//   1. function `func(cc *CliCmd)` or struct.
func (cli *CLI) RegisterEmpty(action interface{}) *CLI {
	cli.actionEmptyRegister = action
	return cli
}

// when command input not handler will callback the register, the format like:
//   1. function `func(cmd string, cc *CliCmd)`
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

// @todo need to make it.
// to star `router`
func (cli *CLI) router(cc *CliCmd) {
	routerValidMk := false
	if cc.Command != "" {
		value := cli.findRegisterValueByCommand(cc.Command)
		if value != nil {
			switch value.(type) {
			// call the functional
			case func(c *CliCmd):
				value.(func(c *CliCmd))(cc)
				routerValidMk = true
			}
		}

		// `unfind` handler
		if !routerValidMk {
			if cli.actionUnfindRegister != nil {
				aur := cli.actionUnfindRegister
				switch aur.(type) {
				case func(cmd string, cc *CliCmd):
					aur.(func(cmd string, cc *CliCmd))(cc.Command, cc)
					routerValidMk = true
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
				aer.(func(cc *CliCmd))(cc)
				routerValidMk = true
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

// 多简直获取键值
func (app *CliCmd) Next(keys ...string) string {
	var value string
	for _, k := range keys {
		value = app.QueueNext(k)
		if value != "" {
			break
		}
	}
	return value
}

// get raw arg data
func (app *CliCmd) ArgRaw(key string) string {
	var value string
	if v, b := app.DataRaw[key]; b {
		value = v
	}
	return value
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
func (app *CliCmd) Arg(key string) interface{} {
	var value interface{} = nil
	if v, b := app.Data[key]; b {
		value = v
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
	}
}

// check the cmd of validation
func isVaildCmd(c string) bool {
	if len(c) == 0 || c[0:1] == "-" {
		return false
	}
	return true
}

// cleanout the raw input string like:
//		`"string"`		=> `string`
//		`"'string'"`	=> `'string'`
//		`'string'`		=> `string`
//		`'"string"'`	=> `"string"`
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
