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
	Next     int                     `json:"next"`
	// When set, indicates the input data of the option but the command
	IsData      bool   `json:"isData"`
	Mark        string `json:"mark"`  // option input name mark for help
	Owner       string `json:"owner"` // single struct app for child command option map key, or naming rule `Opt[Name]`
	List        []string
	FieldName   string // remember fieldName when gen by struct reflect
	StructGen   bool   // parse the struct into documents and values
	StructItems []Option
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
	IsEntry    bool
	OffValid   bool // Turn off validation
	dataOption *Option
	config     *Config
}

func (c CommandOptional) SetConfig(cfg Config) CommandOptional {
	c.config = &cfg
	return c
}

func (c CommandOptional) GetConfig() Config {
	if c.config == nil {
		return DefaultConfig
	}
	return *c.config
}

// OptionHelpMsg generate an options help document through the options parameters you set
func (c CommandOptional) OptionHelpMsg(levels ...int) string {
	level := rock.Param(0, levels...)
	pref := ""
	if level > 0 {
		pref = fmt.Sprintf("%-"+fmt.Sprintf("%d", level*4)+"s", " ")
	}
	var pairList [][2]string
	maxLen := 0
	for _, opt := range c.Options {
		if opt.IsData {
			var vDataOpt = opt
			c.dataOption = &vDataOpt
			continue
		}
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
		line := help
		if len(optList) > 0 {
			line += "，其他别名 " + strings.Join(optList, ",")
		}
		if opt.DefValue != "" {
			line += "，默认值“" + opt.DefValue + "”"
		}
		if opt.Mark != "" {
			name += " " + opt.Mark
		}
		if nameLen := len(name); nameLen > maxLen {
			maxLen = nameLen
		}
		if len(opt.StructItems) > 0 {
			var siPairList [][2]string
			var siMaxLen = 0
			cfg := c.GetConfig()
			for _, item := range opt.StructItems {
				if item.Help == "" {
					continue
				}
				siName := opt.GetName() + cfg.StructGenSep + item.GetName()
				siLen := len(siName)
				if siLen > siMaxLen {
					siMaxLen = siLen
				}
				siHelp := item.Help
				if item.DefValue != "" {
					siHelp += "，默认值“" + item.DefValue + "”"
				}
				siPairList = append(siPairList, [2]string{
					siName, siHelp,
				})
			}

			for _, sPair := range siPairList {
				line += "\n" + pref + "    " +
					"    -" + fmt.Sprintf("%-"+fmt.Sprintf("%d", siMaxLen)+"s", sPair[0]) + "   " + sPair[1]
			}
		}
		pairList = append(pairList, [2]string{name, line})
	}

	var lines []string
	for _, pair := range pairList {
		line := pref + "    " + fmt.Sprintf("%-"+fmt.Sprintf("%d", maxLen+4)+"s", pair[0]) + pair[1]
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}

func (c CommandOptional) SubCommandHelpMsg(levels ...int) string {
	level := rock.Param(0, levels...)
	pref := ""
	if level > 0 {
		pref = fmt.Sprintf("%-"+fmt.Sprintf("%d", level*2)+"s", " ")
	}

	var pairList [][2]string
	var maxLen = 0
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
			alias = "，其他别名 " + strings.Join(keys, ",")
		}
		line := help + alias
		optHelp := sub.OptionHelpMsg(level)
		if optHelp != "" {
			line += "\n" + optHelp
		}
		if vLen := len(name); vLen > maxLen {
			maxLen = vLen
		}
		pairList = append(pairList, [2]string{name, line})
	}

	var lines []string
	for _, pair := range pairList {
		lines = append(lines, pref+fmt.Sprintf("%-"+fmt.Sprintf("%d", maxLen+2)+"s", pair[0])+pair[1])
	}
	return strings.Join(lines, "\n")
}

// InvalidMsg Determine whether an option is valid by validating the option
func (c CommandOptional) InvalidMsg(args ArgsParser) string {
	subCommand := args.SubCommand()
	if subCommand != "" && len(c.SubCommands) > 0 {
		for _, sco := range c.SubCommands {
			if rock.InList(sco.Keys, subCommand) {
				if sco.OffValid {
					return ""
				}
				return sco.InvalidMsg(args)
			}
		}
	}

	var allowOptions []string
	for _, opt := range c.Options {
		if opt.ValidFn != nil {
			invalidMsg := opt.ValidFn(args)
			if invalidMsg != "" {
				return invalidMsg
			}
		}
		allowOptions = append(allowOptions, opt.GetKeys()...)
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

	if allowOptions != nil {
		cfg := c.GetConfig()
		for _, iptOpt := range args.Option() {
			verifyKey := iptOpt
			keyMutilClass := strings.Split(iptOpt, cfg.StructGenSep)
			if len(keyMutilClass) > 0 {
				verifyKey = keyMutilClass[0]
			}
			if !rock.InList(allowOptions, verifyKey) {
				return iptOpt + ": 选项不支持，请参考帮助文档"
			}
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

func (c CommandOptional) NoValid() CommandOptional {
	c.OffValid = true
	return c
}

func (c CommandOptional) DataOption() *Option {
	if c.dataOption != nil {
		return c.dataOption
	}
	var vOpt Option
	for _, opt := range c.Options {
		if opt.IsData {
			vOpt = opt
			return &vOpt
		}
	}
	return nil
}

// Help Used to set help information
func Help(help string, options ...Option) CommandOptional {
	return CommandOptional{Help: help, Options: options}
}

// HelpSub Used to set help information for the subcommand for top command
func HelpSub(help string, subs ...CommandOptional) CommandOptional {
	return CommandOptional{Help: help, SubCommands: subs}
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
