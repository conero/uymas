package bin

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	AppTagName = "cmd"
)

type appRegisterData struct {
	name     string
	alias    []string
	register interface{}
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
	Option   []interface{}
	Register interface{} // register name
}

// AppOption parse struct tag syntax
// `cmd: "tagName;"`
// data valid: required, optional(default)
// name: "name,n,N"
// example help: `help,h; optional`
type AppOption interface {
}
