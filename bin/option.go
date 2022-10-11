package bin

import (
	"reflect"
	"regexp"
	"strings"
)

const OptionTagName = "arg"

// Option the command of options parse.
type Option struct {
	cc *Arg
}

// Unmarshal parse the struct tag name arg <Name type `arg:"i name"`>
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
		if !hasTag {
			tag = strings.ToLower(fld.Name)
		}

		args := c.argParse(tag)

		switch fld.Type.Kind() {
		case reflect.String:
			vv.Field(i).SetString(cc.ArgRaw(args...))
		case reflect.Bool:
			vv.Field(i).SetBool(cc.CheckSetting(args...))
		}
	}
}

// parse arg string.
func (c Option) argParse(tag string) []string {
	pattern := `\s`
	reg, er := regexp.Compile(pattern)
	if er == nil {
		tag = reg.ReplaceAllString(tag, "")
	}
	return strings.Split(tag, ",")
}
