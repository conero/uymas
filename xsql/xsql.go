// Package xsql extends for the sql package.
// power by <Joshua Conero>(2020)
package xsql

import (
	"database/sql"
	"fmt"
)

type Config struct {
	TablePrefix string
	DbType      string
}

type Xsql struct {
	cfg   Config
	conn  *sql.DB
	query *Query
}

func (x *Xsql) Name(tb string, alias ...string) *Builder {
	return Table(fmt.Sprintf("%s%s", x.cfg.TablePrefix, tb), alias...)
}

func (x *Xsql) SelectFunc(call func(*Builder), tb string, alias ...string) ([]map[string]any, error) {
	table := x.Name(tb, alias...)
	if call != nil {
		call(table)
	}
	bSql, bind := table.ToSQL()
	return x.query.Select(bSql, bind...)
}

func NewXsql(conn *sql.DB, cfgs ...Config) *Xsql {
	x := &Xsql{
		conn:  conn,
		query: NewQuery(conn),
	}
	if len(cfgs) > 0 {
		x.cfg = cfgs[0]
	}
	return x
}
