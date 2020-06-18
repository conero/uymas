package xsql

import "testing"

func TestTable(t *testing.T) {
	var vSql string
	var binds []interface{}

	var adminUser = Table("admin_user")
	vSql, binds = adminUser.ToSQL()
	t.Log(vSql, binds)

	adminUser.Where("name like ? and gender = ?", "%Conero%", 1)
	vSql, binds = adminUser.ToSQL()
	t.Log(vSql, binds)

}
