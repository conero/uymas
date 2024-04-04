package bin

import (
	"fmt"
	"gitee.com/conero/uymas/util"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	OptionTagName = "arg"
	OptionTagDef  = "argDef"
)

// Option the command of options parse.
type Option struct {
	cc         *Arg
	allow      []string
	exclude    []string // exclude specified options during validation
	excludeReg []string //
}

// Exclude set exclude specified options
func (c *Option) Exclude(excludes ...string) *Option {
	c.exclude = append(c.exclude, excludes...)
	return c
}

// ExcludeReg use regular expressions to exclude specified options during validation
func (c *Option) ExcludeReg(regs ...string) *Option {
	c.excludeReg = append(c.excludeReg, regs...)
	return c
}

// Unmarshal parse the struct tag name arg <Name type `arg:"i,name"`>
//
// <Name type `argDef:"0"`> set default values
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
		defValue, hasDef := fld.Tag.Lookup(OptionTagDef)

		var args []string
		if !hasTag {
			tag = strings.ToLower(fld.Name)
			args = append(args, tag)
		} else {
			args = c.argParse(tag)
		}

		getFnI64 := func() int64 {
			vI64 := int64(cc.ArgInt(args...))
			if vI64 == 0 && hasDef && defValue != "" {
				vI64, _ = strconv.ParseInt(defValue, 10, 10)
			}
			return vI64
		}
		getFnUi64 := func() uint64 {
			vI64 := uint64(cc.ArgInt(args...))
			if vI64 == 0 && hasDef && defValue != "" {
				vI64, _ = strconv.ParseUint(defValue, 10, 10)
			}
			return vI64
		}
		getFnF64 := func() float64 {
			vf64 := cc.ArgFloat64(args...)
			if vf64 == 0 && hasDef && defValue != "" {
				vf64, _ = strconv.ParseFloat(defValue, 10)
			}
			return vf64
		}
		c.allow = append(c.allow, args...)
		switch fld.Type.Kind() {
		case reflect.String:
			vv.Field(i).SetString(cc.DefString(defValue, args...))
		case reflect.Bool:
			vv.Field(i).SetBool(cc.CheckSetting(args...) || hasDef)
		case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
			vv.Field(i).SetInt(getFnI64())
		case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
			vv.Field(i).SetUint(getFnUi64())
		case reflect.Float64, reflect.Float32:
			vv.Field(i).SetFloat(getFnF64())

		}
	}
}

func (c *Option) NotAllow() []string {
	var unAllow []string
	for _, set := range c.cc.Setting {
		if util.ListIndex(c.allow, set) == -1 {
			if util.ListIndex(c.exclude, set) > -1 {
				continue
			}
			unAllow = append(unAllow, set)
		}
	}
	return unAllow
}

func (c *Option) CheckAllow() error {
	for _, set := range c.cc.Setting {
		if util.ListIndex(c.allow, set) == -1 {
			if util.ListIndex(c.exclude, set) > -1 {
				continue
			}
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
