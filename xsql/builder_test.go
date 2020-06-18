package xsql

import "testing"

func TestTable(t *testing.T) {
	var vSql string
	var binds []interface{}

	//select
	var adminUser = Table("admin_user")
	vSql, binds = adminUser.ToSQL()
	t.Log(vSql, binds)

	adminUser.Where("name like ? and gender = ?", "%Conero%", 1)
	vSql, binds = adminUser.ToSQL()
	t.Log(vSql, binds)

	adminUser.Page(1, 15)
	vSql, binds = adminUser.ToSQL()
	t.Log(vSql, binds)

	adminUser.Page(37)
	vSql, binds = adminUser.ToSQL()
	t.Log(vSql, binds)

	//for paginate
	for i := 38; i < 45; i++ {
		adminUser.Page(i)
		vSql, binds = adminUser.ToSQL()
		t.Log(vSql, binds)
	}

	//update
	adminUser.Update(map[string]interface{}{"allow_login": true, "test_mk": "test u value", "ability": 33.5}).
		ResetWhere("id = ?", 1103).Where("gender = ?", 1)
	vSql, binds = adminUser.ToSQL()
	t.Log(vSql, binds)

	//delete
	adminUser.Delete().ResetWhere("id = ?", 1103).Where("gender = ?", 1)
	vSql, binds = adminUser.ToSQL()
	t.Log(vSql, binds)

	//insert
	adminUser.Insert(map[string]interface{}{"allow_login": true, "test_mk": ".test u value.", "ability": 33.5})
	vSql, binds = adminUser.ToSQL()
	t.Log(vSql, binds)
}
