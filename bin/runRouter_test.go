package bin

import (
	"github.com/conero/uymas/unit"
	"testing"
)

// @Date：   2018/12/18 0018 11:14
// @Author:  Joshua Conero
// @Name:    runRouter 测试

// 修正二级命令
// all-key => AllKey
func TestAmendSubC(t *testing.T) {
	if "AllKey" != AmendSubC("all-key") {

	}
	value := unit.StrSingLine([][]string{
		[]string{"AllKey", "all-key", AmendSubC("all-key")},
	})
	if _, isStr := value.(string); isStr {
		t.Fatal(value)
		return
	}
	if !value.(bool) {
		t.Fail()
	}

}

// runAppRouter 测试
func TestCommand_Run(t *testing.T) {
	// 自定义查询
	InjectArgs("test", "action")
	// 运行入口
	Run()
	if app.Command != "test" {
		t.Fatal("runAppRouter 测试错误：test != " + app.Command)
	}
	if app.SubCommand != "action" {
		t.Fatal("runAppRouter 测试错误：action != " + app.SubCommand)
	}
}
