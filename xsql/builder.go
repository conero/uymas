package xsql

import (
	"fmt"
	"strings"
)

const (
	builderSelect = "SELECT"
)

//need to do.
//the sql builder
type Builder struct {
	table   string
	alias   string
	vSql    string
	binds   []interface{}
	columns []string
	keyword string     // like: select, update, delete
	wheres  [][]string // [][link-word, string]/[][string]
}

func (c *Builder) Select(columns ...string) *Builder {
	c.keyword = builderSelect
	return c.Columns(columns...)
}

func (c *Builder) Columns(columns ...string) *Builder {
	c.columns = append(c.columns, columns...)
	return c
}

func (c *Builder) Where(where string, binds ...interface{}) *Builder {
	c.wheres = append(c.wheres, []string{where})
	c.binds = append(c.binds, binds...)
	return c
}

//build SQL every time
func (c *Builder) ToSQL() (string, []interface{}) {
	keyword := c.keyword
	if keyword == "" {
		keyword = builderSelect
	}
	switch keyword {
	case builderSelect:
		table := c.table
		if c.alias != "" {
			table = table + " " + c.alias
		}

		columns := "*"
		if len(c.columns) > 0 {
			columns = strings.Join(c.columns, ", ")
		}

		where := ""
		if len(c.wheres) > 0 {
			var whereQueue []string
			for _, whs := range c.wheres {
				if len(whs) == 1 {
					if len(whereQueue) == 0 {
						whereQueue = append(whereQueue, whs[0])
					} else {
						whereQueue = append(whereQueue, fmt.Sprintf("AND (%v)", whs[0]))
					}
				} else if len(whs) == 2 {
					whereQueue = append(whereQueue, fmt.Sprintf("%v (%v)", whs[0], whs[1]))
				}
			}

			if len(whereQueue) > 0 {
				where = "WHERE " + strings.Join(whereQueue, " ")
			}
		}

		c.vSql = fmt.Sprintf("%v %v FROM %v %v", builderSelect, columns, table, where)
		c.vSql = strings.TrimSpace(c.vSql)

	}
	return c.vSql, c.binds
}

//only get SQL where exist sql if not will call SQL builder
func (c *Builder) GetSQL() (string, []interface{}) {
	if c.vSql == "" {
		c.ToSQL()
	}
	return c.vSql, c.binds
}

//Table(string/string interface)
//Table([]string{name, alias})
func Table(table interface{}, alias ...string) *Builder {
	if table == nil {
		panic("if params of Table is invalid, `table == nil`")
	}
	c := &Builder{}
	switch table.(type) {
	case string:
		c.table = table.(string)
	case []string:
		tables := table.([]string)
		if len(tables) != 2 {
			panic("if param of Table.table is []string, format should be `(table, alias string)`")
		}
		c.table = tables[0]
		c.alias = tables[1]

	default:
		c.table = fmt.Sprintf("%v", table)
	}

	if len(alias) > 0 {
		c.alias = alias[0]
	}
	return c
}
