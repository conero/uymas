package tag

import (
	"fmt"
	"gitee.com/conero/uymas/bin"
	"gitee.com/conero/uymas/str"
	"gitee.com/conero/uymas/util"
	"reflect"
	"strings"
)

type Runnable interface {
	Exec(arg *bin.Arg)
}

// RunnableCommand runnable comand struct template
type RunnableCommand struct {
	Arg *bin.Arg
}

func (c *RunnableCommand) Exec(arg *bin.Arg) {}

type Parser struct {
	app             any
	appName         Name
	Tags            []Tag
	commandsTagDick map[string][]Tag
	commands        []any // 会议列表
	cli             *bin.CLI
}

// parse the tag of field by struct
func (c *Parser) parse() bool {
	rt := reflect.TypeOf(c.app)
	isPtr := false
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
		isPtr = true
	}

	if rt.Kind() != reflect.Struct {
		return false
	}
	var rv reflect.Value
	if isPtr {
		rv = reflect.ValueOf(c.app).Elem()
	} else {
		rv = reflect.ValueOf(c.app)
	}

	// parse `tag.Name`
	// User configuration is preferred for parameter configuration, that is, existing configuration is not overwritten
	setNameMark := false
	setNameFieldFn := func(fv any, field reflect.Value, sf reflect.StructField, tg *Tag) {
		if setNameMark {
			return
		}
		if tg == nil {
			// setName
			if fv != nil && field.Kind() == reflect.Struct {
				if name, matched := fv.(Name); matched {
					if name.Name == "" {
						name.Name = str.Lcfirst(sf.Name)
						field.Set(reflect.ValueOf(name))
					}
					c.appName = name
					setNameMark = true
				}
			}
		} else {
			// setName
			if fv != nil && field.Kind() == reflect.Struct {
				if name, matched := fv.(Name); matched {
					if name.Name == "" {
						var vName string
						if vs := tg.ValueString(TyAppName); vs != "" {
							vName = vs
						} else {
							str.Lcfirst(sf.Name)
						}
						name.Name = vName
					}

					if name.Alias == "" {
						name.Alias = tg.ValueString(OptAlias)
					}
					if name.Help == "" {
						name.Help = tg.ValueString(OptHelp)
					}
					field.Set(reflect.ValueOf(name))
					c.appName = name
					setNameMark = true
				}
			}
		}

	}

	for i := 0; i < rt.NumField(); i++ {
		sf := rt.Field(i)
		field := rv.Field(i)
		fv := field.Interface()
		cmdStr, exist := sf.Tag.Lookup(CmdTagName)
		if !exist {
			// setName
			setNameFieldFn(fv, field, sf, nil)
			continue
		}

		tg := ParseTag(cmdStr)
		if tg != nil {
			tg.Name = str.Lcfirst(sf.Name)
			if tg.Type == CmdCommand {
				c.parseRunnable(tg, field)
				//fmt.Printf("%v:%#v\n", tg.Name, *tg)
			}
			tg.carrier = rv
			tg.carrierKey = sf.Name
			c.Tags = append(c.Tags, *tg)
			setNameFieldFn(fv, field, sf, tg)
		} else {
			setNameFieldFn(fv, field, sf, nil)
		}
	}

	return true
}

// parse struct that let command to be runnable
func (c *Parser) parseRunnable(tg *Tag, field reflect.Value) {
	if tg == nil {
		return
	}
	kind := field.Kind()
	if kind == reflect.Ptr {
		kind = field.Elem().Kind()
	}

	if kind != reflect.Struct {
		return
	}
	for i := 0; i < field.NumMethod(); i++ {
		mth := field.Method(i)
		if mth.IsValid() {
			v := mth.Interface()
			rMth, matched := v.(func(cmd *bin.Arg))
			if matched {
				tg.runnable = rMth
				break
			}
		}
	}

	c.parseCommandTags(tg, field)
}

// parse commands of app option tags
func (c *Parser) parseCommandTags(tg *Tag, field reflect.Value) {
	if c.commandsTagDick == nil {
		c.commandsTagDick = map[string][]Tag{}
	}
	tDick, _ := c.commandsTagDick[tg.Name]
	targetFld := field
	if field.Kind() == reflect.Ptr {
		targetFld = targetFld.Elem()
	}
	cRt := field.Type().Elem()
	isUpdMk := false
	for i := 0; i < targetFld.NumField(); i++ {
		//cFld := field.Field(i)
		cSf := cRt.Field(i)

		tag, exist := cSf.Tag.Lookup(CmdTagName)
		if !exist {
			continue
		}
		cTg := ParseTag(tag)
		if cTg != nil {
			if cTg.Name == "" {
				name := cTg.ValueString(TyOptionName)
				if name == "" {
					name = str.Lcfirst(cSf.Name)
				}
				cTg.Name = name
				if cTg.CheckOption(TyOptionName) {
					cTg.Type = CmdOption
				}
			}
			cTg.carrierKey = cSf.Name
			cTg.carrier = field
			tDick = append(tDick, *cTg)
			isUpdMk = true
		}
	}

	if isUpdMk {
		c.commandsTagDick[tg.Name] = tDick
	}
}

// NewParser app should be struct app or prt
func NewParser(value any) *Parser {
	ps := &Parser{
		app: value,
	}
	ps.parse()
	ps.genCli()
	return ps
}

func (c *Parser) genCli() {
	cli := bin.NewCLI()

	var cmdsDoc = map[string]string{}

	// 命令注册
	for _, tg := range c.Tags {
		// to handler command call
		if tg.Type == CmdCommand {
			cmds := []string{tg.Name}

			help := tg.ValueString(OptHelp)
			if help == "" {
				help = "命令"
			}

			alias := tg.Value(OptAlias)
			if alias != nil {
				cmds = append(cmds, alias...)
				help += "，别名 " + strings.Join(alias, ",")
			}

			cmdsDoc[tg.Name] = help
			if tg.runnable != nil {
				cli.RegisterFunc(func(cmd *bin.Arg) {
					if !c.validCommand(cmd, tg) {
						return
					}
					tg.runnable(cmd)
				}, cmds...)
			} else {
				cli.RegisterFunc(func(cmd *bin.Arg) {
					if !c.validCommand(cmd, tg) {
						return
					}
					panic("Command doesn't realization exec(*bin.Arg) function!")
				}, cmds...)
			}
		}
	}

	helpFn := func() {
		fmt.Printf("欢饮使用命令行程序，命令格式如下: \n\n$ %v [command] [option]\n", c.appName.Name)
		if len(cmdsDoc) > 0 {
			fmt.Printf("\n命令列表:\n%v\n", bin.FormatKv(cmdsDoc))
		}
	}

	//index
	cli.RegisterEmpty(helpFn)
	// unfind
	cli.RegisterUnmatched(func(s string, cmd *bin.Arg) {
		if s != "" {
			s += " "
		}
		fmt.Printf("Error) %v命令不存在，请键入 help 查看帮助信息！\n\n", s)
	})

	// the command help
	cli.RegisterFunc(func(cmd *bin.Arg) {
		subCmd := cmd.SubCommand
		if subCmd == "" {
			helpFn()
			return
		}

		tag, dick := c.CommandTag(subCmd)
		if tag == nil {
			fmt.Printf("%v 命令不存在", subCmd)
			return
		}

		help := tag.ValueString(OptHelp)
		if help == "" {
			help = "命令"
		}

		alias := tag.Value(OptAlias)
		if alias != nil {
			help += "，别名 " + strings.Join(alias, ",")
		}
		fmt.Printf("%v      %v\n", subCmd, help)

		if dick != nil {
			helpDick := map[string]string{}
			for _, ct := range dick {
				name := strings.Join(ct.Names(), ", ")
				help = "选项参数"
				if vh := ct.ValueString(OptHelp); vh != "" {
					help += vh
				}
				helpDick[name] = help
			}
			if len(helpDick) > 0 {
				fmt.Printf("选项列表如下：\n%v", bin.FormatKv(helpDick, "  "))
			}
		}

	}, "help", "?")

	c.cli = cli
}

// valid command like option
func (c *Parser) validCommand(cc *bin.Arg, tag Tag) bool {
	if c.commandsTagDick == nil || len(c.commandsTagDick) == 0 {
		return true
	}
	cTag, exist := c.commandsTagDick[tag.Name]
	if !exist {
		return true
	}
	var optionList []string
	for _, ct := range cTag {
		if ct.Type != CmdOption {
			continue
		}
		sets := ct.Names()
		if ct.IsRequired() && !cc.CheckSetting(sets...) {
			fmt.Printf("%v: 选项不可为空", strings.Join(sets, ","))
			return false
		}
		optionList = append(optionList, sets...)
		ct.setCarrier(sets, cc)
	}

	// check not allow setting by register
	for _, set := range cc.Setting {
		if util.ListIndex(optionList, set) == -1 {
			fmt.Printf("%v: 选项非法", set)
			return false
		}
	}

	return true
}

// CommandTag get tag sub command
func (c *Parser) CommandTag(name string) (tag *Tag, option []Tag) {
	for _, vtg := range c.Tags {
		if vtg.Own(name) {
			tag = &vtg
			break
		}
	}

	if tag != nil && c.commandsTagDick != nil {
		option, _ = c.commandsTagDick[tag.Name]
	}
	return
}

func (c *Parser) Run() {
	c.cli.Run()
}
