package xsql

import (
	"fmt"
	"strings"
)

const (
	builderInsert = "INSERT"
	builderDelete = "DELETE"
	builderUpdate = "UPDATE"
	builderSelect = "SELECT"
)

// Builder need to do.
// the sql builder
type Builder struct {
	table        string
	alias        string
	vSql         string
	binds        []any
	columns      []string
	keyword      string         // like: select, update, delete
	wheres       [][]string     // [][link-word, string]/[][string]
	condLimit    string         // limit condition
	condOrderBys []string       //order-by
	condGroupBys []string       //group-by
	pageSize     int            // default the pageSize to 20
	data         map[string]any //the data to update or insert
	joins        [][]string     //json table list, [table alias, cond, type]
}

func (c *Builder) Join(table, cond, vtype string) *Builder {
	c.joins = append(c.joins, []string{table, cond, vtype})
	return c
}

func (c *Builder) LeftJoin(table, cond string) *Builder {
	return c.Join(table, cond, "LEFT")
}

func (c *Builder) RightJoin(table, cond string) *Builder {
	return c.Join(table, cond, "RIGHT")
}

func (c *Builder) InnerJoin(table, cond string) *Builder {
	return c.Join(table, cond, "INNER")
}

func (c *Builder) GetTableString() string {
	table := c.table
	if c.alias != "" {
		table = fmt.Sprintf("%v %v", table, c.alias)
	}

	var queue []string
	for _, join := range c.joins {
		vLen := len(join)
		if vLen == 3 {
			queue = append(queue, fmt.Sprintf("%v JOIN %v ON %v", join[2], join[0], join[1]))
		} else if vLen == 2 {
			queue = append(queue, fmt.Sprintf("INNER JOIN %v ON %v", join[0], join[1]))
		}
	}
	if len(queue) > 0 {
		table = fmt.Sprintf("%v %v", table, strings.Join(queue, " "))
	}
	return table
}

func (c *Builder) Insert(data map[string]any) *Builder {
	c.keyword = builderInsert
	c.data = data
	return c
}

func (c *Builder) Delete() *Builder {
	c.keyword = builderDelete
	return c
}

func (c *Builder) Update(data map[string]any) *Builder {
	c.keyword = builderUpdate
	c.data = data
	return c
}

func (c *Builder) Select(columns ...string) *Builder {
	c.keyword = builderSelect
	return c.Columns(columns...)
}

func (c *Builder) Columns(columns ...string) *Builder {
	c.columns = append(c.columns, columns...)
	return c
}

func (c *Builder) Where(where string, binds ...any) *Builder {
	where = strings.TrimSpace(where)
	if where != "" {
		c.wheres = append(c.wheres, []string{where})
		c.binds = append(c.binds, binds...)
	}

	return c
}

func (c *Builder) ResetWhere(where string, binds ...any) *Builder {
	c.binds = []any{}
	c.wheres = [][]string{}
	c.Where(where, binds...)
	return c
}

func (c *Builder) ResetBinds() *Builder {
	c.binds = []any{}
	return c
}

func (c *Builder) Page(page int, orPageSize ...int) *Builder {
	if len(orPageSize) > 0 {
		c.pageSize = orPageSize[0]
	}

	//default the pageSize to 20
	if c.pageSize == 0 {
		c.pageSize = 20
	}

	c.condLimit = fmt.Sprintf("LIMIT %v,%v", (page-1)*c.pageSize, c.pageSize)
	return c
}

// GroupBy call once, then more to reset.
func (c *Builder) GroupBy(groupBys ...string) *Builder {
	c.condGroupBys = groupBys
	return c
}

// OrderBy call once, then more to reset.
func (c *Builder) OrderBy(orderBys ...string) *Builder {
	c.condOrderBys = orderBys
	return c
}

func (c *Builder) Limit(offset int, orRowCount ...int) *Builder {
	if len(orRowCount) > 0 {
		c.condLimit = fmt.Sprintf("LIMIT %v,%v", offset, orRowCount[0])
	} else {
		c.condLimit = fmt.Sprintf("LIMIT %v", offset)
	}
	return c
}

func (c *Builder) parseWhere() string {
	var where string
	if len(c.wheres) == 0 {
		return where
	}
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
	return where
}

func (c *Builder) createInsertSql() {
	var columns []string
	var valueRepls []string
	var binds []any
	for k, v := range c.data {
		columns = append(columns, k)
		valueRepls = append(valueRepls, "?")
		binds = append(binds, v)
	}

	c.binds = binds
	c.vSql = fmt.Sprintf("%v INTO %v (%v) VALUES (%v)",
		builderInsert, c.table, strings.Join(columns, ", "), strings.Join(valueRepls, ", "))
}

func (c *Builder) createDeleteSql() {
	where := c.parseWhere()
	c.vSql = fmt.Sprintf("%v FROM %v %v",
		builderDelete, c.table, where)
	c.vSql = strings.TrimSpace(c.vSql)
}

func (c *Builder) createUpdateSql() {
	var oBinds = c.binds
	var columnsKv []string
	var binds []any
	for k, v := range c.data {
		columnsKv = append(columnsKv, fmt.Sprintf("%v = ?", k))
		binds = append(binds, v)
	}

	if oBinds != nil {
		binds = append(binds, oBinds...)
	}
	c.binds = binds
	where := c.parseWhere()
	c.vSql = fmt.Sprintf("%v %v SET %v %v",
		builderUpdate, c.table, strings.Join(columnsKv, ", "), where)
	c.vSql = strings.TrimSpace(c.vSql)

}

func (c *Builder) createSelectSql() {
	table := c.GetTableString()

	columns := "*"
	if len(c.columns) > 0 {
		columns = strings.Join(c.columns, ", ")
	}

	where := c.parseWhere()

	c.vSql = fmt.Sprintf("%v %v FROM %v %v", builderSelect, columns, table, where)
	//group-by
	if len(c.condGroupBys) > 0 {
		c.vSql += fmt.Sprintf(" GROUP BY %v", strings.Join(c.condGroupBys, ", "))
	}

	//order-by
	if len(c.condOrderBys) > 0 {
		c.vSql += fmt.Sprintf(" ORDER BY %v", strings.Join(c.condOrderBys, ", "))
	}

	//limit
	if c.condLimit != "" {
		c.vSql += " " + c.condLimit
	}

	c.vSql = strings.TrimSpace(c.vSql)
}

// ToSQL build SQL every time
func (c *Builder) ToSQL() (string, []any) {
	keyword := c.keyword
	if keyword == "" {
		keyword = builderSelect
	}
	switch keyword {
	case builderInsert:
		c.createInsertSql()
	case builderDelete:
		c.createDeleteSql()
	case builderUpdate:
		c.createUpdateSql()
	case builderSelect:
		c.createSelectSql()

	}
	return c.vSql, c.binds
}

// GetSQL only get SQL where exist sql if not will call SQL builder
func (c *Builder) GetSQL() (string, []any) {
	if c.vSql == "" {
		c.ToSQL()
	}
	return c.vSql, c.binds
}

// Table Table(string/string interface)
// Table([]string{name, alias})
func Table(table any, alias ...string) *Builder {
	if table == nil {
		panic("if params of Table is invalid, `table == nil`")
	}
	c := &Builder{}
	switch vTable := table.(type) {
	case string:
		c.table = vTable
	case []string:
		tables := vTable
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
