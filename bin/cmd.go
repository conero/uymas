package bin

import (
	"errors"
	"fmt"
	"gitee.com/conero/uymas/util"
	"reflect"
	"strings"
)

const (
	AppTagName            = "cmd"
	OptValidationRequire  = "required" // 必须
	OptValidationOptional = "optional" // 可选
)

const (
	tagSplitDelimiter = ";"
)

type appRegisterData struct {
	name     string
	alias    []string
	register any
}

type App struct {
	Title       string // app title
	Description string // app description text
	CmdList     []AppCmd
	cli         *CLI
	register    []appRegisterData // parse register cmd
	doc         string
}

// Append add new cmd
func (c *App) Append(ac ...AppCmd) *App {
	c.CmdList = append(c.CmdList, ac...)
	return c
}

func (c *App) getCli() *CLI {
	if c.cli == nil {
		c.parseDoc()
		c.cli = NewCLI()
		// register <help>
		c.cli.RegisterFunc(func(cmd *CliCmd) {
			fmt.Println(c.GetDoc() + "\n")
		}, "help")

		// register cmd list
		for _, rgst := range c.register {
			cmd := []string{rgst.name}
			if len(rgst.alias) > 0 {
				cmd = append(cmd, rgst.alias...)
			}
			if rgst.register != nil {
				allow := false
				switch rgst.register.(type) {
				case func(), func(cmd *CliCmd), func(cmd CliCmd):
					allow = true
				default:
					rv := reflect.ValueOf(rgst.register)
					if rv.Kind() == reflect.Ptr {
						rv = rv.Elem()
					}
					if rv.Kind() == reflect.Struct {
						allow = true
					}
				}
				if !allow {
					panic(fmt.Sprintf("%v Register not allow type %#v", cmd, c.register))
				}
				c.cli.register(rgst.register, cmd...)
			}
		}
	}
	return c.cli
}

// Run run cmd application
func (c *App) Run(args ...string) {
	c.getCli().Run(args...)
}

// get string value judge default value
func defString(v, def string) string {
	v = strings.TrimSpace(v)
	if v == "" {
		v = strings.TrimSpace(def)
	}
	return v
}

func (c *App) parseDoc() string {
	if c.doc == "" {
		var que []string
		que = append(que, fmt.Sprintf("%v", defString(c.Title, "the app of uymas.")))
		//blank b2
		que = append(que, fmt.Sprintf("  %v\n", defString(c.Description, "power by gitee.com/conero/uymas")))
		//b2
		que = append(que, "Usage:  $ [command] [option]")

		existCmd := len(c.CmdList) > 0
		if existCmd {
			que = append(que, "\n")
		}
		//命令行读取
		for _, cmd := range c.CmdList {
			vCmd := fmt.Sprintf("%v    %v", cmd.Name, cmd.Title)
			que = append(que, vCmd)
			c.register = append(c.register, appRegisterData{
				name:     cmd.Name,
				alias:    cmd.Alias,
				register: cmd.Register,
			})
		}

		c.doc = strings.Join(que, "\n")
	}
	return c.doc
}

func (c *App) GetDoc() string {
	return c.parseDoc()
}

type AppCmd struct {
	Alias    []string // the alias list of cmd
	Name     string   // cmd name default by `AppCmd` struct
	Title    string   // the description of cmd
	Option   []any
	Register any // register name
}

// AppOptionGroup the group of many of AppOption.
type AppOptionGroup struct {
	lastCmd    *CliCmd
	optionDick map[string]*AppOption
}

func (c *AppOptionGroup) ParseEach(v any, each func(*AppOption)) error {
	vf := reflect.ValueOf(v)
	isStruct := false

	vKind := vf.Kind()
	if vKind == reflect.Struct {
		isStruct = true
	} else if vKind == reflect.Ptr {
		isStruct = vf.Elem().Kind() == reflect.Struct
	}

	if !isStruct {
		return errors.New(fmt.Sprintf("the param v is not (ptr) struct type."))
	}

	optDick := map[string]*AppOption{}
	vt := reflect.TypeOf(v)
	if vt.Kind() == reflect.Ptr {
		vt = vt.Elem()
	}
	for i := 0; i < vt.NumField(); i++ {
		field := vt.Field(i)
		tagStr := field.Tag.Get(AppTagName)
		var opt *AppOption
		if tagStr != "" {
			opt = ParseOptionTag(tagStr)
			if each != nil {
				each(opt)
			}
		} else {
			opt = &AppOption{
				Name: strings.ToLower(field.Name),
			}
		}

		if opt != nil {
			optDick[opt.Name] = opt
		}
	}

	c.optionDick = optDick
	return nil
}

func (c *AppOptionGroup) Unmarshal(cmd *Arg, v any) error {
	c.lastCmd = cmd
	err := c.ParseEach(v, func(opt *AppOption) {
	})
	return err
}

func ParseOptionGroup(v any) *AppOptionGroup {
	aog := &AppOptionGroup{}
	err := aog.ParseEach(v, nil)
	if err != nil {
		return nil
	}
	return aog
}

// Option get option by name
func (c *AppOptionGroup) Option(name string) *AppOption {
	if c.optionDick != nil {
		opt, exist := c.optionDick[name]
		if exist {
			return opt
		}
		return c.OptionSearchAlias(name)
	}
	return nil
}

func (c *AppOptionGroup) OptionSearchAlias(name string) *AppOption {
	if c.optionDick != nil {
		for k, opt := range c.optionDick {
			if k == name {
				return opt
			}
			if util.ListIndex(opt.Alias, name) > -1 {
				return opt
			}
		}
	}
	return nil
}

// AppOption parse struct tag syntax
// `cmd: "tagName;"`
// data valid: required, optional(default)
// name: "name,n,N"
// example help: `help,h; optional`
type AppOption struct {
	Validation string // validation item like: required, optional(default)
	Name       string
	Alias      []string
	raw        string // the raw of tag string
	IsSet      bool   // option is set
}

// parse tag string
func (c *AppOption) parse() {
	queue := strings.Split(c.raw, tagSplitDelimiter)
	for i, s := range queue {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		// Parse tag item: `name`
		if i == 0 { // example: "`cmd: name,n,N;`"
			s = strings.TrimSpace(s)
			var alias []string
			for j, vName := range strings.Split(s, ",") {
				if j == 0 {
					c.Name = strings.TrimSpace(vName)
				} else {
					alias = append(alias, strings.TrimSpace(vName))
				}
			}
			c.Alias = alias
		} else if s == OptValidationRequire || s == OptValidationOptional { // parse value of OptValidation.
			c.Validation = s
		}
	}
}

func ParseOptionTag(tag string) *AppOption {
	ap := &AppOption{
		Validation: OptValidationOptional,
		raw:        tag,
	}
	ap.parse()
	return ap
}
