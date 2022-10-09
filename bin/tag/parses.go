package tag

import (
	"gitee.com/conero/uymas/str"
	"reflect"
)

type Parser struct {
	value any
	Tags  []Tag
}

// parse the tag of field by struct
func (c *Parser) parse() bool {
	rt := reflect.TypeOf(c.value)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	if rt.Kind() != reflect.Struct {
		return false
	}

	for i := 0; i < rt.NumField(); i++ {
		sf := rt.Field(i)
		cmdStr, exist := sf.Tag.Lookup(CmdTagName)
		if !exist {
			continue
		}

		tg := ParseTag(cmdStr)
		if tg != nil {
			tg.Name = str.Lcfirst(sf.Name)
		}
	}

	return true
}

// NewParser value should be struct value or prt
func NewParser(value any) *Parser {
	ps := &Parser{
		value: value,
	}
	ps.parse()
	return ps
}
