package cli

import (
	"fmt"
	"gitee.com/conero/uymas/v2/rock"
	"strings"
)

// Option Used for command option parsing document generation, or value validation and retrieval
type Option struct {
	//@todo Name and Alias will merge into single field: Keys
	Name     string                  `json:"name"`
	Alias    []string                `json:"alias"`
	Require  bool                    `json:"require"`
	ValidFn  func(ArgsParser) string `json:"-"`
	DefValue string                  `json:"defValue"`
	Help     string                  `json:"help"`
}

// GetName Gets option names automatically compatible with aliases or actual names
func (c Option) GetName() string {
	if c.Name != "" {
		return c.Name
	}

	if len(c.Alias) > 0 {
		return c.Alias[0]
	}
	return ""
}

func (c Option) GetKeys() []string {
	var keys []string
	if c.Name != "" {
		keys = append(keys, c.Name)
	}

	if len(c.Alias) > 0 {
		keys = append(keys, c.Alias...)
	}
	return keys
}

func (c Option) IsRequire() Option {
	c.Require = true
	return c
}

func (c Option) Default(s string) Option {
	c.DefValue = s
	return c
}

func OptionHelp(help string, keys ...string) Option {
	return Option{
		Help:  help,
		Alias: keys,
	}
}

// CommandOptional Used for command registration as a parameter option
type CommandOptional struct {
	Help  string
	Name  string
	Alias []string
	// A command list that includes commands and aliases
	Keys        []string
	Options     []Option
	SubCommands []CommandOptional
	// Whether the subcommand is an entry
	IsEntry bool
}

// OptionHelpMsg generate an options help document through the options parameters you set
func (c CommandOptional) OptionHelpMsg(levels ...int) string {
	level := rock.Param(0, levels...)
	pref := ""
	if level > 0 {
		pref = fmt.Sprintf("%-"+fmt.Sprintf("%d", level*4)+"s", " ")
	}
	var lines []string
	for _, opt := range c.Options {
		var optList []string
		if opt.Name != "" {
			optList = append(optList, opt.Name)
		}
		optList = append(optList, opt.Alias...)
		optList = optionRecoverRawList(optList)
		var name string
		var optNum = len(optList)
		if optNum <= 2 {
			name = strings.Join(optList, ",")
			optList = []string{}
		} else if optNum > 2 {
			name = strings.Join(optList[:2], ",")
			optList = optList[2:]
		}

		help := opt.Help
		if help == "" {
			help = "支持参数选项"
		}
		if opt.Require {
			help = "* " + help
		}
		line := "    " + name + "    " + help
		if len(optList) > 0 {
			line += "，支持别名 " + strings.Join(optList, ",")
		}
		if opt.DefValue != "" {
			line += "，默认值“" + opt.DefValue + "”"
		}
		lines = append(lines, pref+line)

	}
	return strings.Join(lines, "\n")
}

func (c CommandOptional) SubCommandHelpMsg(levels ...int) string {
	level := rock.Param(0, levels...)
	pref := ""
	if level > 0 {
		pref = fmt.Sprintf("%-"+fmt.Sprintf("%d", level*4)+"s", " ")
	}
	var lines []string
	for _, sub := range c.SubCommands {
		if sub.IsEntry {
			continue
		}

		keyNum := len(sub.Keys)
		if keyNum == 0 {
			continue
		}
		name := sub.Keys[0]
		keys := sub.Keys[1:]
		help := sub.Help
		if help == "" {
			help = "子命令"
		}

		alias := ""
		if keyNum > 1 {
			alias = "，支持别名 " + strings.Join(keys, ",")
		}
		line := name + "    " + help + alias
		optHelp := sub.OptionHelpMsg(level)
		if optHelp != "" {
			line += "\n" + optHelp
		}

		lines = append(lines, pref+line)
	}
	return strings.Join(lines, "\n")
}

// InvalidMsg Determine whether an option is valid by validating the option
func (c CommandOptional) InvalidMsg(args ArgsParser) string {
	subCommand := args.SubCommand()
	if subCommand != "" && len(c.SubCommands) > 0 {
		for _, sco := range c.SubCommands {
			if rock.InList(sco.Keys, subCommand) {
				return sco.InvalidMsg(args)
			}
		}
	}

	for _, opt := range c.Options {
		if opt.ValidFn != nil {
			invalidMsg := opt.ValidFn(args)
			if invalidMsg != "" {
				return invalidMsg
			}
		}
		if !opt.Require {
			continue
		}

		var alias []string
		if opt.Name != "" {
			alias = append(alias, opt.Name)
		}
		alias = append(alias, opt.Alias...)
		if opt.DefValue != "" || args.Switch(alias...) {
			continue
		}

		value := args.Get(alias...)
		if value == "" {
			return strings.Join(optionRecoverRawList(alias), ",") + "  必须为选项设置值"
		}

	}
	return ""
}

func (c CommandOptional) GetDefault(keys ...string) string {
	for _, opt := range c.Options {
		for _, key := range keys {
			if rock.InList(opt.GetKeys(), key) {
				return opt.DefValue
			}
		}
	}
	return ""
}

// NameAlias Set a command or alias, used for subcommand document registration
func (c CommandOptional) NameAlias(name string, alias ...string) CommandOptional {
	c.Name = name
	c.Alias = append(c.Alias, alias...)
	c.Keys = append(c.Keys, name)
	c.Keys = append(c.Keys, alias...)
	return c
}

// SubEntry Marked as a subcommand entry
func (c CommandOptional) SubEntry() CommandOptional {
	c.IsEntry = true
	c.Name = ""
	return c
}

func (c CommandOptional) SubCommand(subName string) (optional CommandOptional, isFind bool) {
	for _, co := range c.SubCommands {
		if rock.InList(co.Keys, subName) {
			optional = co
			isFind = true
			return
		}
	}
	return
}

// Help Used to set help information
func Help(help string, options ...Option) CommandOptional {
	return CommandOptional{Help: help, Options: options}
}

// HelpSub Used to set help information for the subcommand for top command
func HelpSub(help string, commands ...CommandOptional) CommandOptional {
	return CommandOptional{Help: help, SubCommands: commands}
}

func optionRecoverRaw(option string) string {
	if len(option) == 1 {
		return fmt.Sprintf("-%s", option)
	}
	return fmt.Sprintf("--%s", option)
}

func optionRecoverRawList(options []string) []string {
	for i, opt := range options {
		options[i] = optionRecoverRaw(opt)
	}
	return options
}