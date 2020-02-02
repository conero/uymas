package bin

import (
	"github.com/conero/uymas/bin/butil"
	"github.com/conero/uymas/str"
	"os"
)

// the cli application
type CLI struct {
	cmds   map[string]interface{} // the register of commands.
	cmdMap map[string]interface{} // the map the cmd, support many alias cmd about the cmd
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
	cli := &CLI{}
	return cli
}

// the construct of `CliCmd`
func NewCliCmd(args ...string) *CliCmd {
	c := &CliCmd{
		Raw: args,
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
//  `RegisterFunc(todo func(cmd *CliCmd), cmd string)` or `RegisterFunc(todo func(), cmd, alias string)`
func (cli *CLI) RegisterFunc(todo func(cmd *CliCmd), cmds ...string) *CLI {
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
	if cc.Command != "" {
	}
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
func (app *CliCmd) parseArgs() {
	if app.Raw != nil {
		for i, arg := range app.Raw {
			if i == 0 && isVaildCmd(arg) {
				app.Command = arg
			} else if i == 1 && isVaildCmd(arg) {
				app.SubCommand = arg
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
