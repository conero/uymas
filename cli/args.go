package cli

import (
	"gitee.com/conero/uymas/v2/data/input"
	"gitee.com/conero/uymas/v2/rock"
	"os"
	"strings"
)

// ArgsParser command line parameter parsing interface
type ArgsParser interface {
	// Values The original data type of the command
	Values() map[string][]string
	// Get command line data by key value
	Get(keys ...string) string
	Int(keys ...string) int
	Int64(keys ...string) int64
	F64(keys ...string) float64
	Uint64(keys ...string) uint64
	// Def get command line data by key value and specify default values
	Def(def string, keys ...string) string
	IntDef(def int, keys ...string) int
	Int64Def(def int64, keys ...string) int64
	Uint64Def(def uint64, keys ...string) uint64
	F64Def(def float64, keys ...string) float64
	// Switch determines whether the option specified by the key value exists
	Switch(keys ...string) bool
	// Command get the command of the command line program
	Command() string
	SubCommand() string
	Option() []string
	CommandList() []string

	// Next find command from `CommandList`
	Next(cmds ...string) string

	// HelpCmd get help command
	HelpCmd(params ...[]string) string
}

type ArgsConfig struct {
	// supports short option resolution, if it's true then `-xyz` will same like `-x`, `-y`, `-z`.
	//
	ShortOption bool
	EqualSigner []string
}

var DefArgsConfig = ArgsConfig{
	EqualSigner: []string{"="},
}

// Args command line program parameters
type Args struct {
	raw         []string
	command     string
	subCommand  string
	commandList []string
	option      []string
	values      map[string][]string
	config      ArgsConfig
	ArgsParser
}

// parse data by args
func (c *Args) parse() {
	if c.values == nil {
		c.values = map[string][]string{}
	}
	config := c.config

	// split option as KV pairs
	optionSplitFn := func(opt string) []string {
		for _, eq := range config.EqualSigner {
			idx := strings.Index(opt, eq)
			if idx > -1 {
				return []string{opt[:idx], opt[idx+1:]}
			}
		}
		return nil
	}

	lastKey := ""
	// remember option
	recordOptionFn := func(opts ...string) {
		if opts == nil {
			return
		}
		vLen := len(opts)
		if vLen == 1 {
			// --xy=222222222222, -xxxx=cvy
			pair := optionSplitFn(opts[0])
			if len(pair) > 0 {
				lastKey = pair[0]
				var values = c.values[lastKey]
				values = append(values, pair[1])
				c.values[lastKey] = values
				c.option = append(c.option, lastKey)
				return
			}
		}
		c.option = append(c.option, opts...)
		lastKey = opts[vLen-1]
	}
	for i, arg := range c.raw {
		var option string
		idx := strings.Index(arg, "-")
		if idx == 0 {
			if strings.Index(arg, "--") == 0 {
				recordOptionFn(arg[2:])
				continue
			}

			option = arg[1:]
			if config.ShortOption {
				optionList := strings.Split(option, "")
				recordOptionFn(optionList...)
				continue
			}

			recordOptionFn(option)
			continue

		}
		if i == 0 {
			c.command = arg
			c.commandList = append(c.commandList, arg)
			continue
		} else if i == 1 && c.command != "" {
			c.subCommand = arg
			c.commandList = append(c.commandList, arg)
			continue
		}
		if lastKey != "" {
			var values = c.values[lastKey]
			values = append(values, arg)
			c.values[lastKey] = values
			continue
		}
		c.commandList = append(c.commandList, arg)
	}
}

func (c *Args) Values() map[string][]string {
	return c.values
}

func (c *Args) GetValues(keys ...string) []string {
	if c.values == nil {
		return nil
	}

	for _, key := range keys {
		list, exist := c.values[key]
		if exist {
			return list
		}
	}

	return nil
}

func (c *Args) GetValueJoin(seq string, keys ...string) string {
	if c.values == nil {
		return ""
	}

	values := c.GetValues(keys...)
	if len(values) > 0 {
		return strings.Join(values, seq)
	}
	return ""
}

func (c *Args) Get(keys ...string) string {
	return c.GetValueJoin(" ", keys...)
}

func (c *Args) Def(def string, keys ...string) string {
	value := c.Get(keys...)
	if value != "" {
		return value
	}
	return def
}

func (c *Args) IntDef(def int, keys ...string) int {
	value := c.Int(keys...)
	if value != 0 {
		return value
	}
	return def
}

func (c *Args) Int64Def(def int64, keys ...string) int64 {
	value := c.Int64(keys...)
	if value != 0 {
		return value
	}
	return def
}

func (c *Args) Uint64Def(def uint64, keys ...string) uint64 {
	value := c.Uint64(keys...)
	if value != 0 {
		return value
	}
	return def
}

func (c *Args) F64Def(def float64, keys ...string) float64 {
	value := c.F64(keys...)
	if value != 0 {
		return value
	}
	return def
}

func (c *Args) Switch(keys ...string) bool {
	for _, key := range keys {
		if rock.InList(c.option, key) {
			return true
		}
	}
	return false
}

func (c *Args) Command() string {
	return c.command
}

func (c *Args) SubCommand() string {
	return c.subCommand
}

func (c *Args) Option() []string {
	return c.option
}

func (c *Args) CommandList() []string {
	return c.commandList
}

func (c *Args) Next(cmds ...string) string {
	vLen := len(c.commandList)
	if vLen == 0 {
		return ""
	}

	for _, cm := range cmds {
		for i, refCmd := range c.commandList {
			if cm == refCmd && i < vLen-1 {
				return c.commandList[i+1]
			}
		}
	}
	return ""
}

func (c *Args) HelpCmd(params ...[]string) string {
	cmds := rock.ParamIndex(1, []string{"help"}, params...)
	command := c.Next(cmds...)
	if command != "" {
		return command
	}

	opts := rock.ParamIndex(2, []string{"help", "h"}, params...)
	return c.Get(opts...)
}

func (c *Args) Int(keys ...string) int {
	value := c.Get(keys...)
	if value != "" {
		return input.Stringer(value).Int()
	}
	return 0
}

func (c *Args) Int64(keys ...string) int64 {
	value := c.Get(keys...)
	if value != "" {
		return input.Stringer(value).Int64()
	}
	return 0
}

func (c *Args) F64(keys ...string) float64 {
	value := c.Get(keys...)
	if value != "" {
		return input.Stringer(value).Float()
	}
	return 0
}

func (c *Args) Uint64(keys ...string) uint64 {
	value := c.Get(keys...)
	if value != "" {
		return input.Stringer(value).Uint64()
	}
	return 0
}

func NewArgs(args ...string) ArgsParser {
	if len(args) == 0 {
		args = os.Args[1:]
	}
	arg := &Args{
		raw:    args,
		config: DefArgsConfig,
	}
	arg.parse()
	return arg
}

func NewArgsWith(config ArgsConfig, args ...string) ArgsParser {
	if len(args) == 0 {
		args = os.Args[1:]
	}
	arg := &Args{
		raw:    args,
		config: config,
	}
	arg.parse()
	return arg
}
