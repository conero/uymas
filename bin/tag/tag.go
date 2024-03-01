// Package tag reflect struct tag for bin
package tag

import (
	"gitee.com/conero/uymas/v2/bin"
	"gitee.com/conero/uymas/v2/bin/parser"
	"gitee.com/conero/uymas/v2/util"
	"reflect"
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
	CmdFieldArg       = "Arg"
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
	Name       string
	Type       uint8
	Raw        string
	Values     map[string][]string
	runnable   func(*bin.Arg)
	carrier    reflect.Value
	carrierKey string
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

	// to set `Tag.Name` by function
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

// Own check name its own be tag by name or alias
func (c *Tag) Own(name string) bool {
	if c.Name == name {
		return true
	}
	alias := c.Value(OptAlias)
	if util.ListIndex(alias, name) > -1 {
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

func (c *Tag) setArgsCarrier(cc *bin.Arg) {
	carrier := c.carrier
	if carrier.IsNil() || carrier.IsZero() || !carrier.IsValid() {
		return
	}
	if carrier.Kind() == reflect.Ptr {
		carrier = carrier.Elem()
	}
	value := carrier.FieldByName(CmdFieldArg)
	if value.IsValid() {
		if value.IsNil() || value.IsZero() {
			switch value.Interface().(type) {
			case *bin.Arg:
				value.Set(reflect.ValueOf(cc))
			}
		}
	}
}

// set the value of carrier.
func (c *Tag) setCarrier(names []string, cc *bin.Arg) bool {
	if c.carrierKey == "" {
		return false
	}
	if c.carrier.IsNil() || c.carrier.IsZero() || !c.carrier.IsValid() {
		return false
	}
	// set the field named `Arg`
	c.setArgsCarrier(cc)
	valid := c.carrier
	if valid.Kind() == reflect.Ptr {
		valid = valid.Elem()
	}

	if valid.Kind() != reflect.Struct {
		return false
	}

	field := valid.FieldByName(c.carrierKey)
	if !field.IsValid() {
		return false
	}

	//@todo Currently, string types are supported temporarily. Later, generic types are used to support more types
	// set value
	isSet := false
	rawValue := cc.ArgRaw(names...)
	switch field.Kind() {
	case reflect.String:
		val := reflect.ValueOf(rawValue)
		field.Set(val)
		isSet = true
	case reflect.Bool:
		vBool := false
		if rawValue != "" {
			vBool = parser.ConvBool(rawValue)
		} else {
			vBool = cc.CheckSetting(names...)
		}
		val := reflect.ValueOf(vBool)
		field.Set(val)
		isSet = true
	case reflect.Int:
		val := reflect.ValueOf(parser.ConvInt(rawValue))
		field.Set(val)
		isSet = true
	case reflect.Int64:
		val := reflect.ValueOf(parser.ConvI64(rawValue))
		field.Set(val)
		isSet = true
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint, reflect.Uint8,
		reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val := reflect.ValueOf(parser.ConvInt(rawValue))
		field.Set(val.Convert(field.Type()))
		isSet = true
	case reflect.Float64, reflect.Float32:
		val := reflect.ValueOf(parser.ConvF64(rawValue))
		field.Set(val.Convert(field.Type()))
		isSet = true
	}
	return isSet
}
