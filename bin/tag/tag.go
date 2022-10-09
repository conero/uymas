// Package tag reflect struct tag for bin
package tag

import (
	"gitee.com/conero/uymas/bin"
	"gitee.com/conero/uymas/str"
	"strings"
)

const (
	CmdType uint8 = iota
	CmdApp
	CmdCommand
	CmdOption
)

const (
	TyAppName     = "app"
	TyCommandName = "command"
	TyOptionName  = "option"
)

const (
	CmdTypeKey = "CTY"
	// CmdCommandField the command sets of field
	CmdCommandField = "Commands"
	// CmdExecFn let command to be runnable
	CmdExecFn         = "Exec"
	CmdTagName        = "cmd"
	tagSplitLimiter   = " "
	tagKvEqualLimiter = ":"
	tagMultiLimiter   = ","
)

const (
	OptRequire = "require"
	OptShort   = "short"
	OptAlias   = "alias"
	OptHelp    = "help"
)

type Tag struct {
	// the name is field of struct
	Name     string
	Type     uint8
	Raw      string
	Values   map[string][]string
	runnable func(*bin.CliCmd)
}

// CheckOption check if option exist
func (c *Tag) CheckOption(args ...string) bool {
	for _, arg := range args {
		_, exist := c.Values[arg]
		if exist {
			return true
		}
	}
	return false
}

func (c *Tag) Value(args ...string) []string {
	for _, arg := range args {
		value, exist := c.Values[arg]
		if exist {
			return value
		}
	}
	return nil
}

// ValueString get string app
func (c *Tag) ValueString(args ...string) string {
	values := c.Value(args...)
	if values != nil {
		return strings.Join(values, tagMultiLimiter)
	}
	return ""
}

func (c *Tag) IsRequired() bool {
	return c.CheckOption(OptRequire)
}

func (c *Tag) nameForShort() string {
	if c.Name != "" {
		return strings.ToLower(c.Name[:1])
	}
	return ""
}

func (c *Tag) Short() []string {
	exist := c.CheckOption(OptShort)
	if exist {
		values := c.Value(OptShort)
		if values != nil {
			return values
		}

		if defShort := c.nameForShort(); defShort != "" {
			values = append(values, defShort)
		}
	}
	return nil
}

func ParseTag(vTag string) *Tag {
	tg := &Tag{
		Type:   CmdApp,
		Raw:    vTag,
		Values: map[string][]string{},
	}

	setNameFn := func(as []string) {
		var name string
		for _, a := range as {
			if a != "" {
				name = a
				break
			}
		}

		// defined name by tag
		if name != "" {
			tg.Name = name
		}
	}

	for _, s := range strings.Split(vTag, tagSplitLimiter) {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}

		idx := strings.Index(s, tagKvEqualLimiter)
		var key string
		var value string
		values, _ := tg.Values[key]
		if idx > -1 {
			key = strings.TrimSpace(s[:idx])
			value = strings.TrimSpace(s[idx+1:])
			idx = strings.Index(value, tagMultiLimiter)
			if idx > -1 {
				value = strings.ReplaceAll(value, tagSplitLimiter, "")
				values = append(values, strings.Split(value, tagMultiLimiter)...)
			} else {
				values = append(values, value)
			}
		} else {
			key = s
		}

		// default command type
		switch key {
		case TyAppName:
			tg.Type = CmdApp
			setNameFn(values)
		case TyCommandName:
			tg.Type = CmdCommand
			setNameFn(values)
		case TyOptionName:
			tg.Type = CmdOption
			setNameFn(values)
		}

		tg.Values[key] = values
	}
	return tg
}

// Own check name is own be tag by name or alias
func (c *Tag) Own(name string) bool {
	if c.Name == name {
		return true
	}
	alias := c.Value(OptAlias)
	if str.InQue(name, alias) > -1 {
		return true
	}
	return false
}

func (c *Tag) Names() []string {
	var names []string
	if c.Name != "" {
		names = append(names, c.Name)
	}
	alias := c.Value(OptAlias)
	if len(alias) > 0 {
		names = append(names, alias...)
	}
	return names
}
