package bin

import (
	"fmt"
	"gitee.com/conero/uymas"
	"gitee.com/conero/uymas/bin/butil"
	"gitee.com/conero/uymas/bin/parser"
	"gitee.com/conero/uymas/str"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const (
	actionRunConstruct = "Construct"
	actionRunIndex     = "DefaultIndex"     // the AppCli empty Call
	actionRunUnmatched = "DefaultUnmatched" // the AppCli unmatched router Call
	actionRunHelp      = "DefaultHelp"      // the AppCli unmatched router Call
	actionRunHelpName  = "Help"             // the alis of `help`, will to call `DefaultHelp` method.
	appCliFieldCliCmd  = "Cc"               // the AppCli field of Cc.
)

const (
	scriptFileOption = "file,f"   // --file,-f <script-file>
	scriptOption     = "script,s" // --script,-s <script>
)

// CLI the cli application
type CLI struct {
	cmds map[string]any // the register of commands.

	// the map the cmd, support many alias cmd about the cmd
	// the data struct like: {cmd => alias} or {cmd => [alias1, alias2, alias3...]}
	cmdMap              map[string]any
	actionEmptyRegister any // the register callback by empty action.
	actionAnyRegister   any // the register callback by command not handler
	commands            map[string]Cmd
	tempLastCommand     string         // command Cache
	injectionData       map[string]any //reject data from outside like chan control
	registerCmdList     []string       // register name list

	//external fields
	UnLoadDataSyntax   bool   //not support load data syntax, like json/url.
	UnLoadScriptSyntax bool   // disable allow load script like shell syntax.
	ScriptOption       string // default: --script,-s
	ScriptFileOption   string // default: --file,-f
}

// CliCmd the command of the cli application.
type CliCmd struct {
	Data            map[string]any    // the data from the `DataRaw` by parse for type
	DataRaw         map[string]string // the cli application apply the data
	Command         string            // the current command
	SubCommand      string            // the sub command
	Setting         []string          // the setting of command
	Raw             []string          // the raw args
	context         CLI
	cmdType         int                 //the command type enumeration
	commandAlias    map[string][]string // the alias of command, using for App-style and runtime state
	subCommandAlias map[string][]string // the alias of command, using for App-style and runtime state
}

// CliApp the cli app.
type CliApp struct {
	Cc *CliCmd
}

// CliAppInterface the interface of CliApp
type CliAppInterface interface {
	Construct()
}

// CliAppCompleteInterface the complete CliApp show hand method
// should have a field name like `Cc *CliCmd`
// the method call order by `construct > command > help > index > unmatched`
type CliAppCompleteInterface interface {
	CliAppInterface
	DefaultHelp()
	DefaultIndex()
	DefaultUnmatched()
}

type CmdOptions struct {
	Option   string
	Alias    any
	Describe string
}

// Cmd define the struct command
type Cmd struct {
	Command  string
	Alias    any                   //string, []string. the alias of the command
	Describe string                //describe the command
	Handler  func(cc *CliCmd)      //when command call then handler the request
	Options  map[string]CmdOptions // the command option
}

// NewCLI the construct of `CLI`
func NewCLI() *CLI {
	cli := &CLI{
		cmds:     map[string]any{},
		cmdMap:   map[string]any{},
		commands: map[string]Cmd{},
	}
	return cli
}

// NewCliCmd the construct of `CliCmd`, args default set os.Args if no function arguments
func NewCliCmd(args ...string) *CliCmd {
	if args == nil {
		// if the args is empty then use the `os.Args`
		osArgs := os.Args
		if len(osArgs) > 1 {
			args = osArgs[1:]
		}
	}
	c := &CliCmd{
		Raw:     args,
		Setting: []string{},
		DataRaw: map[string]string{},
		Data:    map[string]any{},
	}
	// parse the args
	c.parseArgs()
	return c
}

// NewCliCmdByString construction of `CliCmd` by string
//
//	@todo notice: `--test-string="Joshua 存在空格的字符串 Conero"` 解析失败
func NewCliCmdByString(ss string) *CliCmd {
	return NewCliCmd(butil.StringToArgs(ss)...)
}

// GetCmdList get the list cmd of application
func (cli *CLI) GetCmdList() []string {
	var list = cli.registerCmdList
	if cli.cmds != nil {
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

// RegisterFunc register functional command, the format like
//
//	`RegisterFunc(todo func(cc *CliCmd), cmd string)` or `RegisterFunc(todo func(), cmd, alias string)`
func (cli *CLI) RegisterFunc(todo func(*CliCmd), cmds ...string) *CLI {
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

// register command by function or struct
func (cli *CLI) register(rgst any, cmds ...string) {
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
		// register feedback
		cli.cmds[cmd] = rgst
	}
}

func (cli *CLI) registerFunc(todo func(*CliCmd), cmds ...string) {
	cli.tempLastCommand = ""
	if len(cmds) > 0 {
		cmd := cmds[0]
		cli.tempLastCommand = cmd
		cli.cmds[cmd] = todo
		if len(cmds) > 1 {
			cli.cmdMap[cmd] = cmds[1:]
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

// GetDescribe support the `cmd, alias` param.
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

// RegisterCommand register by command struct data
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

// RegisterApp register the struct app, the format same as RegisterFunc. cmds any be `cmd string` or `cmd, alias string`
func (cli *CLI) RegisterApp(ap any, cmds ...string) *CLI {
	if cmds != nil && len(cmds) > 0 {
		cmd := cmds[0]
		cli.cmds[cmd] = ap
		if len(cmds) > 1 {
			cli.cmdMap[cmd] = cmds[1]
		}
	}
	return cli
}

// RegisterApps Register many apps once.
func (cli *CLI) RegisterApps(aps map[string]any) *CLI {
	for c, ap := range aps {
		cli.RegisterApp(ap, c)
	}
	return cli
}

// RegisterEmpty when the cmd is empty then callback the function, action only be
//  1. function `func(cc *CliCmd)`/`func()` or struct.
func (cli *CLI) RegisterEmpty(action any) *CLI {
	cli.actionEmptyRegister = action
	return cli
}

// RegisterAny when command input not handler will callback the register, the format like:
//  1. function `func(cmd string, cc *CliCmd)`/`func(cmd string)`/`func(cc *CliCmd)`/CliApp/Base Struct
func (cli *CLI) RegisterAny(action any) *CLI {
	cli.actionAnyRegister = action
	// check if cmd dist
	rv := reflect.ValueOf(action)
	if rv.Kind() == reflect.Ptr && rv.Elem().Kind() == reflect.Struct {
		rt := reflect.TypeOf(action)
		defMth := []string{actionRunConstruct, actionRunHelp, actionRunIndex, actionRunUnmatched}
		for i := 0; i < rv.NumMethod(); i++ {
			vMth := rt.Method(i)
			name := vMth.Name
			if str.InQue(name, defMth) > -1 {
				continue
			}
			cli.registerCmdList = append(cli.registerCmdList, strings.ToLower(name))
		}
	}
	return cli
}

// RegisterUnmatched old method `Unfind` of alias compatibly use `RegisterAny`
func (cli *CLI) RegisterUnmatched(callback func(string, *CliCmd)) *CLI {
	cli.RegisterAny(callback)
	return cli
}

// Run the application
func (cli *CLI) Run(args ...string) {
	// construct of `CliCmd`
	cmd := NewCliCmd(args...)
	// start router by register.
	cli.router(cmd)
}

// RunDefault Run cli app using user defined `args` when os.args is empty,
// it'll be useful to debug or default define.
func (cli *CLI) RunDefault(args ...string) {
	osArgs := os.Args
	if len(osArgs) > 1 {
		args = osArgs[1:]
	}
	cli.Run(args...)
}

// CallCmd call the application cmd
func (cli *CLI) CallCmd(cmd string) {
	cm := NewCliCmd(cmd)
	cli.router(cm)
}

// CmdExist test cmd exist in application
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
	// call the `before-call-hook`
	cli.hookBeforeCall(cc)

	// router command is not empty, include func or App.
	isRouterMk := false
	if cc.Command != "" {
		isRouterMk = cli.routerCommand(cc)
	} else { // router command is empty.
		isRouterMk = cli.routerEmpty(cc)
	}

	// router command is default.
	if !isRouterMk {
		if cli.actionAnyRegister != nil {
			isRouterMk = cli.routerAny(cc)
		}

		if !isRouterMk {
			if cc.Command != "" {
				fmt.Printf(" Fail: the command `%v` not find.\n", cc.Command)
				fmt.Println()
			}
			fmt.Printf("   Power by framework %v, Version: %v/%v.\n", uymas.PkgName,
				uymas.Version, uymas.Release)
			fmt.Printf("   You call register `RegisterAny` handler for it.\n")
			fmt.Printf("   Welcome to learn more doc from link. https://pkg.go.dev/gitee.com/conero/uymas \n")
			fmt.Println()
			isRouterMk = false
		}
	}
}

// router when `command` is not empty.
func (cli *CLI) routerCommand(cc *CliCmd) bool {
	routerValidMk := false
	value := cli.findRegisterValueByCommand(cc.Command)
	if value != nil {
		switch value.(type) {
		// call the FuncCmd
		case func():
			value.(func())()
			routerValidMk = true
		case func(cmd CliCmd):
			value.(func(CliCmd))(*cc)
			routerValidMk = true
		case func(*CliCmd):
			value.(func(*CliCmd))(cc)
			routerValidMk = true
		default:
			// call the AppCmd
			v := reflect.ValueOf(value)
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			// set `Cc` that is struct of field.
			if cCField := v.FieldByName(appCliFieldCliCmd); cCField.IsValid() {
				cc.cmdType = int(CmdApp)
				switch cCField.Interface().(type) {
				case *CliCmd:
					v.FieldByName(appCliFieldCliCmd).Set(reflect.ValueOf(cc))
				}
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

			// many Call-func
			callFunc := func(vMethod string) {
				if vMth := v.MethodByName(vMethod); vMth.IsValid() {
					switch vMth.Interface().(type) {
					case func(*CliCmd):
						routerValidMk = true
						vMth.Call([]reflect.Value{reflect.ValueOf(cc)})
					case func():
						routerValidMk = true
						vMth.Call(nil)
					}
				}
			}

			//the subCommand string
			subCmdStr := cc.SubCommand
			if subCmdStr != "" {
				subCmdStr = cc.getAlias(cc.subCommandAlias, subCmdStr)
				subCmdStr = Cmd2StringMap(subCmdStr)
				callFunc(subCmdStr)
			}

			// actionRunHelp, support the command/option like:
			//	command		=> help,?
			//	option		=> --help,-h,-?
			if !routerValidMk && (subCmdStr == actionRunHelpName || subCmdStr == "?" || (subCmdStr == "" && cc.
				CheckSetting("help", "h", "?"))) {
				callFunc(actionRunHelp)
			}

			// actionRunIndex
			if !routerValidMk && subCmdStr == "" {
				callFunc(actionRunIndex)
			}

			// actionRunUnmatched
			if !routerValidMk {
				callFunc(actionRunUnmatched)
			}

			if !routerValidMk && subCmdStr != "" {
				panic(fmt.Sprintf("the method `%s` do not have a handler as `%v`.", cc.SubCommand, subCmdStr))
			}

			routerValidMk = true
		}
	}
	return routerValidMk
}

// router when `command` is empty.
func (cli *CLI) routerEmpty(cc *CliCmd) bool {
	routerValidMk := false
	runFunc := func(vFunc any) {
		switch vFunc.(type) {
		case func(*CliCmd):
			vFunc.(func(*CliCmd))(cc)
			routerValidMk = true
		case func():
			vFunc.(func())()
			routerValidMk = true
		}
	}
	if cli.actionEmptyRegister != nil {
		runFunc(cli.actionEmptyRegister)
	} else if cli.actionAnyRegister != nil {
		runFunc(cli.actionAnyRegister)
	}
	return routerValidMk
}

// router for any call
func (cli *CLI) routerAny(cc *CliCmd) bool {
	isRouterMk := false
	aur := cli.actionAnyRegister
	switch aur.(type) {
	case func(string, *CliCmd):
		aur.(func(string, *CliCmd))(cc.Command, cc)
		isRouterMk = true
	case func(string):
		aur.(func(string))(cc.Command)
		isRouterMk = true
	case func(*CliCmd):
		aur.(func(*CliCmd))(cc)
		isRouterMk = true
	default:
		// actionAnyRegister support the like `CliApp` any struct
		rv := reflect.ValueOf(aur)
		if rv.Kind() == reflect.Ptr && rv.Elem().Kind() == reflect.Struct {
			rvEl := rv.Elem()
			// CliCmd field
			if cCField := rvEl.FieldByName(appCliFieldCliCmd); cCField.IsValid() {
				switch cCField.Interface().(type) {
				case *CliCmd:
					rvEl.FieldByName(appCliFieldCliCmd).Set(reflect.ValueOf(cc))
				case CliCmd:
					rvEl.FieldByName(appCliFieldCliCmd).Set(reflect.ValueOf(cc).Elem())
				}
			}
			// init-method
			if initMth := rv.MethodByName(actionRunConstruct); initMth.IsValid() {
				initMth.Call(nil)
			}
			// many Call-func
			callFunc := func(vMethod string) {
				if vMth := rv.MethodByName(vMethod); vMth.IsValid() {
					switch vMth.Interface().(type) {
					case func(*CliCmd):
						isRouterMk = true
						vMth.Call([]reflect.Value{reflect.ValueOf(cc)})
					case func():
						isRouterMk = true
						vMth.Call(nil)
					}
				}
			}
			var cmdTitle string
			// try to find `command`
			if cc.Command != "" {
				cmdTitle = cc.getAlias(cc.commandAlias, cc.Command)
				cmdTitle = Cmd2StringMap(cmdTitle)
				// check `Construct` repeat call(2 times)
				if cmdTitle != actionRunConstruct {
					// call method
					callFunc(cmdTitle)
				}
			}

			// actionRunHelp, support the command/option like:
			//	command		=> help,?
			//	option		=> --help,-h,-?
			if !isRouterMk && (cmdTitle == actionRunHelpName || cmdTitle == "?" || (cmdTitle == "" && cc.
				CheckSetting("help", "h", "?"))) {
				callFunc(actionRunHelp)
			}

			//default empty call be a index action.
			if !isRouterMk && cmdTitle == "" {
				callFunc(actionRunIndex)
			}

			// actionRunUnmatched
			if !isRouterMk {
				callFunc(actionRunUnmatched)
			}

			// finally not match any method will print the tips.
			if !isRouterMk {
				log.Printf("[WARNING] the method `RegisterUnfind` of param is valid, please reference the doc.")
			}
		}

	}
	return isRouterMk
}

// the hook before call the func
func (cli *CLI) hookBeforeCall(cc *CliCmd) {
	cli.loadDataSyntax(cc)
	cli.loadScriptSyntax(cc)
}

// let program exit by unconventionally
func (cli *CLI) hookInterruptExit() {
	os.Exit(0)
}

// to do load data by setting syntax
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

// to checkout if load script file, script mutil lines
func (cli *CLI) loadScriptSyntax(cc *CliCmd) {
	if cli.UnLoadScriptSyntax {
		return
	}
	// script file
	fileOpt := str.GetNotEmpty(cli.ScriptFileOption, scriptFileOption)
	scriptFile := cc.ArgRaw(strings.Split(fileOpt, ",")...)
	if scriptFile != "" && cc.Command == "" {
		lines := parser.NewScriptFile(scriptFile)
		if len(lines) > 0 {
			for _, line := range lines {
				for _, command := range parser.ParseLine(line) {
					if len(command) > 0 {
						cli.Run(command...)
					}
				}
			}
		} else {
			panic("script is empty or load fail.")
		}
		cli.hookInterruptExit()
	}

	// script multi line string
	scriptOpt := str.GetNotEmpty(cli.ScriptOption, scriptOption)
	script := cc.ArgRaw(strings.Split(scriptOpt, ",")...)
	if script != "" && cc.Command == "" {
		for _, command := range parser.ParseLine(script) {
			if len(command) > 0 {
				cli.Run(command...)
			}
		}
		cli.hookInterruptExit()
	}
}

// find the register value by command.
func (cli *CLI) findRegisterValueByCommand(c string) any {
	var value any = nil
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

// Inject inject for data from outside.
func (cli *CLI) Inject(key string, value any) *CLI {
	if cli.injectionData == nil {
		cli.injectionData = map[string]any{}
	}
	cli.injectionData[key] = value
	return cli
}

// GetInjection get Injection data
func (cli *CLI) GetInjection(key string) any {
	if cli.injectionData == nil {
		return nil
	}
	value, has := cli.injectionData[key]
	if has {
		return value
	}
	return nil
}

// CheckSetting to checkout if the set exist in `CliCmd` sets and support multi.
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

// CheckMustKey check the data key must in the sets and support multi
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

// Cwd get the application current word dir.
func (app *CliCmd) Cwd() string {
	return butil.GetBasedir()
}

// QueueNext get next key from order left to right
func (app *CliCmd) QueueNext(key string) string {
	idx := -1
	qLen := len(app.Raw)
	var value string
	for i := 0; i < qLen; i++ {
		if idx == i {
			value = app.Raw[i]
			break
		}
		if key == app.Raw[i] {
			idx = i + 1
		}
	}
	return value
}

// Next Get key values from multiple key values
func (app *CliCmd) Next(keys ...string) string {
	var value string
	var vLen = len(keys)
	//when keys is empty default use the current Next value that next of `app.Command` or `queue index-2`
	if vLen == 0 {
		if app.Command != "" {
			return app.Next(app.Command)
		} else if len(app.Raw) > 1 {
			return app.Raw[1]
		}
		return ""
	}
	for _, k := range keys {
		value = app.QueueNext(k)
		if value != "" {
			break
		}
	}
	return value
}

// ArgRaw get raw args data, because some args has alias list.
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

// ArgInt get args data see as int
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

// ArgStringSlice get string-slice param args
func (app *CliCmd) ArgStringSlice(keys ...string) []string {
	value := app.Arg(keys...)
	if value != nil {
		switch value.(type) {
		case []string:
			return value.([]string)
		case string:
			return []string{value.(string)}
		default:
			var vSlice []string
			vr := reflect.ValueOf(value)
			if vr.Kind() == reflect.Array || vr.Kind() == reflect.Slice {
				for i := 0; i < vr.Len(); i++ {
					vSlice = append(vSlice, fmt.Sprintf("%v", vr.Index(i).Interface()))
				}
			} else {
				vSlice = append(vSlice, fmt.Sprintf("%v", keys))
			}
			return vSlice
		}
	}
	return nil
}

// ArgRawDefault get raw arg has default
func (app *CliCmd) ArgRawDefault(key, def string) string {
	var value = def
	if v, b := app.DataRaw[key]; b {
		value = v
	}
	return value
}

// Arg get arg after parsed the raw data
func (app *CliCmd) Arg(keys ...string) any {
	var value any = nil
	for _, key := range keys {
		if v, b := app.Data[key]; b {
			value = v
			break
		}
	}
	return value
}

// ArgDefault can default value to get the arg
func (app *CliCmd) ArgDefault(key string, def any) any {
	var value = def
	if v, b := app.Data[key]; b {
		value = v
	}
	return value
}

// ArgRawLine get the raw line input.
func (app *CliCmd) ArgRawLine() string {
	return strings.Join(app.Raw, " ")
}

// CallCmd call cmd
func (app *CliCmd) CallCmd(cmd string) {
	context := app.context
	if context.CmdExist(cmd) && app.Command != cmd {
		app.Command = cmd
		context.router(app)
	}
}

// Context get the context of `CLI`, in case `AppCmd` not `FunctionCmd`
func (app *CliCmd) Context() CLI {
	return app.context
}

func (app *CliCmd) CmdType() int {
	return app.cmdType
}

// AppendData append the Data
func (app *CliCmd) AppendData(vMap map[string]any) *CliCmd {
	if len(vMap) > 0 {
		if app.Data == nil {
			app.Data = map[string]any{}
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
//  1. `$ [command] [option]`
//  2. `$ [command] [sub_command]`
//  3. `$ [option]`
//
// the option format example:
//
//	`[command] -xyz` same like `[command] -x -y -z`
//	`[command] --version --name 'Joshua Conero'`
//	`[command] --list A B C D -L A B C D`
//	`[command] --name='Joshua Conero'`
//
// @todo app.Data --> 类型解析太简陋；支持类型与 Readme.md 不统一
func (app *CliCmd) parseArgs() {
	if app.Raw != nil {
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
							k, v := arg[0:idx], arg[idx+1:]
							app.saveOptionDick(k, v)
							if str.InQue(k, app.Setting) == -1 {
								app.Setting = append(app.Setting, k)
							}
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
							tmpArrLen := len(tmpArr)
							if tmpArrLen > 0 {
								optKey = tmpArr[tmpArrLen-1]
							}
						}
						markKeySuccess = true
					}
				}

				if !markKeySuccess && optKey != "" {
					app.saveOptionDick(optKey, CleanoutString(arg))
				}
			}
		}
	}
}

// merge data when parse options, Synchronous write Data and RawData.
func (app *CliCmd) saveOptionDick(key string, value string) {
	vRaw := value
	if oV, hasOv := app.DataRaw[key]; hasOv {
		vRaw = fmt.Sprintf("%v %v", oV, value)
		if cData, hasData := app.Data[key]; hasData {
			switch cData.(type) {
			case string:
				oldSs := app.Data[key].(string)
				app.Data[key] = []string{oldSs, value}
			case []string:
				oldVar := app.Data[key].([]string)
				oldVar = append(oldVar, value)
				app.Data[key] = oldVar
			}
		}
	} else {
		app.Data[key] = value
	}
	app.DataRaw[key] = vRaw
}

func (app *CliCmd) addAlias(value map[string][]string, key string, alias ...string) map[string][]string {
	if value == nil {
		value = map[string][]string{}
	}
	que, hasKey := value[key]
	if hasKey {
		que = append(que, alias...)
	} else {
		que = alias
	}

	value[key] = que
	return value
}

func (app *CliCmd) addAliasAll(value map[string][]string, alias map[string][]string) map[string][]string {
	if value == nil {
		value = alias
	} else {
		for key, vm := range alias {
			oAlias, has := value[key]
			if has {
				oAlias = append(oAlias, vm...)
			} else {
				oAlias = vm
			}
			value[key] = oAlias
		}
	}
	return value
}

func (app *CliCmd) getAlias(value map[string][]string, c string) string {
	for key, alias := range value {
		if key == c {
			return key
		}
		for _, a := range alias {
			if a == c {
				return key
			}
		}
	}

	return c
}

// CommandAlias Tip: in the future will merge method like CommandAlias And CommandAliasAll, chose one from twos.
func (app *CliCmd) CommandAlias(key string, alias ...string) *CliCmd {
	app.commandAlias = app.addAlias(app.commandAlias, key, alias...)
	return app
}

func (app *CliCmd) CommandAliasAll(alias map[string][]string) *CliCmd {
	app.commandAlias = app.addAliasAll(app.commandAlias, alias)
	return app
}

func (app *CliCmd) SubCommandAlias(key string, alias ...string) *CliCmd {
	app.subCommandAlias = app.addAlias(app.subCommandAlias, key, alias...)
	return app
}

func (app *CliCmd) SubCommandAliasAll(alias map[string][]string) *CliCmd {
	app.subCommandAlias = app.addAliasAll(app.subCommandAlias, alias)
	return app
}

// check the cmd of validation
func isVaildCmd(c string) bool {
	if len(c) == 0 || c[0:1] == "-" {
		return false
	}
	return true
}

// CleanoutString clear out the raw input string like:
//
//	`"string"`		=> `string`
//	`"'string'"`	=> `'string'`
//	`'string'`		=> `string`
//	`'"string"'`	=> `"string"`
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

// ParseValueByStr parse the command value to really type by format.
func ParseValueByStr(ss string) any {
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
