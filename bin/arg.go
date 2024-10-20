package bin

import (
	"fmt"
	"gitee.com/conero/uymas/bin/butil"
	"gitee.com/conero/uymas/bin/parser"
	"gitee.com/conero/uymas/str"
	"gitee.com/conero/uymas/util/rock"
	"os"
	"reflect"
	"strings"
)

type ArgConfig struct {
	// If it's support `--long` that is true, otherwise as `-long`.
	//
	// And `-long` same as `-l -o -n -g`, otherwise `--long`.
	LongOption bool
	// support `:` as equity
	EqualColon bool
}

var DefaultArgConfig = ArgConfig{
	LongOption: true,
	EqualColon: false,
}

// Arg the command of the cli application.
type Arg struct {
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
	isPlgCmd        bool                // is plugin command type
	config          *ArgConfig
}

// CheckSetting checkout if the set exist in `Arg` sets and support multi.
func (app *Arg) CheckSetting(sets ...string) bool {
	has := false
	for _, set := range sets {
		if idx := rock.ListIndex(app.Setting, set); idx > -1 {
			has = true
			break
		}
	}
	return has
}

// CheckMustKey check the data key must in the sets and support multi
func (app *Arg) CheckMustKey(keys ...string) bool {
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
func (app *Arg) Cwd() string {
	return butil.Basedir()
}

// QueueNext get next key from order left to right
func (app *Arg) QueueNext(key string) string {
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
func (app *Arg) Next(keys ...string) string {
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

// NextList get the next list value exclude option.
func (app *Arg) NextList(keys ...string) []string {
	var list []string

	keyMatch := false
	isEmptyKey := len(keys) == 0
	for i, arg := range app.Raw {
		if i == 0 {
			continue
		}

		if strings.Index(arg, "-") == 0 {
			break
		}

		if isEmptyKey {
			keyMatch = true
		} else if !keyMatch && rock.ListIndex(keys, arg) > -1 {
			keyMatch = true
			continue
		}

		if keyMatch {
			list = append(list, arg)
		}

	}
	return list
}

// ArgRaw get raw args data, because some args has alias list.
func (app *Arg) ArgRaw(keys ...string) string {
	var value string
	for _, key := range keys {
		if v, b := app.DataRaw[key]; b {
			value = v
			break
		}
	}

	return value
}

// ArgInt get args data identified as int
func (app *Arg) ArgInt(keys ...string) int {
	value := app.ArgRaw(keys...)
	return parser.ConvInt(value)
}

// ArgBool get args data identified as bool
func (app *Arg) ArgBool(keys ...string) bool {
	value := app.ArgRaw(keys...)
	return parser.ConvBool(value)
}

// ArgFloat64 get args data identified as float64
func (app *Arg) ArgFloat64(keys ...string) float64 {
	value := app.ArgRaw(keys...)
	return parser.ConvF64(value)
}

// ArgStringSlice get string-slice param args
func (app *Arg) ArgStringSlice(keys ...string) []string {
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
				vSlice = append(vSlice, fmt.Sprintf("%v", value))
			}
			return vSlice
		}
	}
	return nil
}

func (app *Arg) ArgIntSlice(keys ...string) []int {
	value := app.Arg(keys...)
	if value != nil {
		switch value.(type) {
		case []int:
			return value.([]int)
		case int:
			return []int{value.(int)}
		default:
			var vSlice []int
			vr := reflect.ValueOf(value)
			if vr.Kind() == reflect.Array || vr.Kind() == reflect.Slice {
				for i := 0; i < vr.Len(); i++ {
					vSlice = append(vSlice, str.StringAsInt(vr.Index(i).String()))
				}
			} else {
				vSlice = append(vSlice, str.StringAsInt(fmt.Sprintf("%v", value)))
			}
			return vSlice
		}
	}
	return nil
}

// ArgRawDefault get raw arg has default
func (app *Arg) ArgRawDefault(key, def string) string {
	var value = def
	if v, b := app.DataRaw[key]; b {
		value = v
	}
	return value
}

// Arg get arg after parsed the raw data
func (app *Arg) Arg(keys ...string) any {
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
func (app *Arg) ArgDefault(key string, def any) any {
	var value = def
	if v, b := app.Data[key]; b {
		value = v
	}
	return value
}

// ArgRawLine get the raw line input.
func (app *Arg) ArgRawLine() string {
	return strings.Join(app.Raw, " ")
}

// CallCmd call cmd
func (app *Arg) CallCmd(cmd string) {
	context := app.context
	if context.CmdExist(cmd) && app.Command != cmd {
		app.Command = cmd
		context.router(app)
	}
}

// Context get the context of `CLI`, in case `AppCmd` not `FunctionCmd`
func (app *Arg) Context() CLI {
	return app.context
}

func (app *Arg) CmdType() int {
	return app.cmdType
}

// AppendData append the Data
func (app *Arg) AppendData(vMap map[string]any) *Arg {
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
func (app *Arg) parseArgs() {
	if app.Raw == nil {
		return
	}
	if app.config == nil {
		app.config = &DefaultArgConfig
	}
	config := app.config
	optKey := ""

	toSplitFn := func(s string) {
		idx := strings.Index(s, "=")
		if idx == -1 && config.EqualColon {
			idx = strings.Index(s, ":")
		}
		if idx > -1 { // --key=value
			optKey = ""
			k, v := s[0:idx], s[idx+1:]
			app.saveOptionDick(k, v)
			if rock.ListIndex(app.Setting, k) == -1 {
				app.Setting = append(app.Setting, k)
			}
		} else {
			optKey = s
			app.Setting = append(app.Setting, s)
		}
	}

	for i, arg := range app.Raw {
		if i == 0 && isValidCmd(arg) {
			app.Command = arg
		} else if i == 1 && app.Command != "" && isValidCmd(arg) {
			app.SubCommand = arg
		} else {
			markKeySuccess := false
			if len(arg) > 1 && "-" == arg[0:1] {
				if "--" == arg[0:2] { // --option
					if config.LongOption { // support `-option`
						arg = arg[2:]
					} else { // support `option`
						arg = arg[1:]
					}
					toSplitFn(arg)
				} else { // -option
					arg = arg[1:]
					if config.LongOption {
						tmpArr := strings.Split(arg, "")
						optKey = ""
						tmpArrLen := len(tmpArr)
						if tmpArrLen > 0 {
							toSplitFn(tmpArr[tmpArrLen-1])
							tmpArr = tmpArr[:tmpArrLen-1]
						}
						app.Setting = append(app.Setting, tmpArr...)

					} else {
						toSplitFn(arg)
					}
				}
				markKeySuccess = true
			}

			if !markKeySuccess && optKey != "" {
				app.saveOptionDick(optKey, CleanoutString(arg))
			}
		}
	}
}

// merge data when parse options, Synchronous write Data and RawData.
func (app *Arg) saveOptionDick(key string, value string) {
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

func (app *Arg) addAlias(value map[string][]string, key string, alias ...string) map[string][]string {
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

func (app *Arg) addAliasAll(value map[string][]string, alias map[string][]string) map[string][]string {
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

func (app *Arg) getAlias(value map[string][]string, c string) string {
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
func (app *Arg) CommandAlias(key string, alias ...string) *Arg {
	app.commandAlias = app.addAlias(app.commandAlias, key, alias...)
	return app
}

func (app *Arg) CommandAliasAll(alias map[string][]string) *Arg {
	app.commandAlias = app.addAliasAll(app.commandAlias, alias)
	return app
}

func (app *Arg) SubCommandAlias(key string, alias ...string) *Arg {
	app.subCommandAlias = app.addAlias(app.subCommandAlias, key, alias...)
	return app
}

func (app *Arg) SubCommandAliasAll(alias map[string][]string) *Arg {
	app.subCommandAlias = app.addAliasAll(app.subCommandAlias, alias)
	return app
}

func (app *Arg) IsPlgCmd() bool {
	return app.isPlgCmd
}

// DefString @todo In the near future, generics will be used instead of template functions
func (app *Arg) DefString(def string, args ...string) string {
	optValue := app.ArgRaw(args...)
	if optValue == "" {
		return def
	}
	return optValue
}

func (app *Arg) DefInt(def int, args ...string) int {
	optValue := app.ArgInt(args...)
	if optValue == 0 {
		return def
	}
	return optValue
}

func (app *Arg) DefF64(def float64, args ...string) float64 {
	optValue := app.ArgFloat64(args...)
	if optValue == 0 {
		return def
	}
	return optValue
}

// ParseOption parse the option object using parameters
func (app *Arg) ParseOption(v any) *Option {
	opt := &Option{cc: app}
	opt.Unmarshal(v)
	return opt
}

func (app *Arg) Config() ArgConfig {
	if app.config == nil {
		return DefaultArgConfig
	}
	return *app.config
}

// NewCliCmd the construct of `Arg`, args default set os.Args if no function arguments
func NewCliCmd(args ...string) *Arg {
	return NewArgWith(&DefaultArgConfig, args...)
}

// NewArgWith Args is instantiated in tape configuration mode
func NewArgWith(cfg *ArgConfig, args ...string) *Arg {
	if cfg == nil {
		cfg = &DefaultArgConfig
	}

	if args == nil {
		// if the args is empty then use the `os.Args`
		osArgs := os.Args
		if len(osArgs) > 1 {
			args = osArgs[1:]
		}
	}
	c := &Arg{
		Raw:     args,
		Setting: []string{},
		DataRaw: map[string]string{},
		Data:    map[string]any{},
		config:  cfg,
	}
	// parse the args
	c.parseArgs()
	return c

}

// NewCliCmdByString construction of `Arg` by string
func NewCliCmdByString(ss string) *Arg {
	return NewCliCmd(butil.StringToArgs(ss)...)
}
