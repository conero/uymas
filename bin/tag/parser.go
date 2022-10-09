package tag

import (
	"fmt"
	"gitee.com/conero/uymas/bin"
	"gitee.com/conero/uymas/str"
	"reflect"
	"strings"
)

type Runnable interface {
	Exec(*bin.CliCmd)
}

type Parser struct {
	app      any
	appName  Name
	Tags     []Tag
	commands []any // 会议列表
	cli      *bin.CLI
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
			c.Tags = append(c.Tags, *tg)
			setNameFieldFn(fv, field, sf, tg)
		} else {
			setNameFieldFn(fv, field, sf, nil)
		}
	}

	return true
}

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
			rMth, matched := v.(func(cmd *bin.CliCmd))
			if matched {
				tg.runnable = rMth
				break
			}
		}
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
				cli.RegisterFunc(tg.runnable, cmds...)
			} else {
				cli.RegisterFunc(func(cmd *bin.CliCmd) {
					panic("Command doesn't realization exec(*bin.CliCmd) function!")
				}, cmds...)
			}
		}
	}

	//index
	cli.RegisterEmpty(func() {
		fmt.Printf("欢饮使用命令行程序，命令格式如下: \n\n$ %v [command] [option]\n", c.appName.Name)
		if len(cmdsDoc) > 0 {
			fmt.Printf("\n命令列表:\n%v\n", bin.FormatKv(cmdsDoc))
		}
	})
	// unfind
	cli.RegisterUnmatched(func(s string, cmd *bin.CliCmd) {
		if s != "" {
			s += " "
		}
		fmt.Printf("Error) %v命令不存在，请键入 help 查看帮助信息！\n\n", s)
	})

	c.cli = cli
}

func (c *Parser) Run() {
	c.cli.Run()
}
