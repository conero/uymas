// Package tag reflect struct tag for bin
package tag

import (
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
	CmdTypeKey        = "CTY"
	CmdTagName        = "cmd"
	tagSplitLimiter   = " "
	tagKvEqualLimiter = ":"
	tagMultiLimiter   = ","
)

const (
	OptRequire = "require"
	OptShort   = "short"
)

type Tag struct {
	// the name is field of struct
	Name   string
	Type   uint8
	Raw    string
	Values map[string][]string
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

// ValueString get string value
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

	for i, s := range strings.Split(vTag, tagSplitLimiter) {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}

		if i == 0 {
			switch s {
			case TyAppName:
				tg.Type = CmdApp
			case TyCommandName:
				tg.Type = CmdCommand
			case TyOptionName:
				tg.Type = CmdOption
			}
		} else {
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

			tg.Values[key] = values
		}

	}
	return tg
}
