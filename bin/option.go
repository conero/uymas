package bin

import (
	"fmt"
	"gitee.com/conero/uymas/util"
	"reflect"
	"regexp"
	"strings"
)

const OptionTagName = "arg"

// Option the command of options parse.
type Option struct {
	cc    *Arg
	allow []string
}

// Unmarshal parse the struct tag name arg <Name type `arg:"i,name"`>
func (c *Option) Unmarshal(v any) {
	cc := c.cc
	vt := reflect.TypeOf(v).Elem()
	vv := reflect.ValueOf(v).Elem()
	i := vt.NumField()
	for {
		i -= 1
		if i < 0 {
			break
		}
		fld := vt.Field(i)
		tag, hasTag := fld.Tag.Lookup(OptionTagName)

		var args []string
		if !hasTag {
			tag = strings.ToLower(fld.Name)
			args = append(args, tag)
		} else {
			args = c.argParse(tag)
		}

		c.allow = append(c.allow, args...)
		switch fld.Type.Kind() {
		case reflect.String:
			vv.Field(i).SetString(cc.ArgRaw(args...))
		case reflect.Bool:
			vv.Field(i).SetBool(cc.CheckSetting(args...))
		}
	}
}

func (c *Option) NotAllow() []string {
	var unAllow []string
	for _, set := range c.cc.Setting {
		if util.ListIndex(c.allow, set) == -1 {
			unAllow = append(unAllow, set)
		}
	}
	return unAllow
}

func (c *Option) CheckAllow() error {
	for _, set := range c.cc.Setting {
		if util.ListIndex(c.allow, set) == -1 {
			return fmt.Errorf(" unexpected argument '%s' found", set)
		}
	}
	return nil
}

// parse arg string.
func (c *Option) argParse(tag string) []string {
	pattern := `\s`
	reg, er := regexp.Compile(pattern)
	if er == nil {
		tag = reg.ReplaceAllString(tag, "")
	}
	return strings.Split(tag, ",")
}
