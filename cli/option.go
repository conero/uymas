package cli

import (
	"fmt"
	"gitee.com/conero/uymas/v2/rock"
	"strings"
)

// Option Used for command option parsing document generation, or value validation and retrieval
type Option struct {
	Name     string
	Alias    []string
	Require  bool
	ValidFn  func(ArgsParser) string
	DefValue string
	Help     string
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

// CommandOptional Used for command registration as a parameter option
type CommandOptional struct {
	Help        string
	Alias       []string
	Options     []Option
	SubCommands []CommandOptional
}

// OptionHelpMsg generate an options help document through the options parameters you set
func (c CommandOptional) OptionHelpMsg() string {
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
			help = "为选择参数"
		}
		line := "    " + name + "    " + help
		if len(optList) > 0 {
			line += "，支持别名 " + strings.Join(optList, ",")
		}
		lines = append(lines, line)

	}
	return strings.Join(lines, "\n")
}

// InvalidMsg Determine whether an option is valid by validating the option
func (c CommandOptional) InvalidMsg(args ArgsParser) string {
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

		alias := []string{opt.Name}
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

// Help Used to set help information
func Help(help string, options ...Option) CommandOptional {
	return CommandOptional{Help: help, Options: options}
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
