package bin

import (
	"fmt"
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

// 字符解析
func TestStrParseData(t *testing.T) {
	// bool
	b1 := StrParseData("true")
	if b, is := b1.(bool); !is || !b {
		t.Error("bool 解析失败， is, b := ", is, ", ", b)
	}

	// ---------------------------
	// string
	w1 := "I'am the big man."
	s1 := StrParseData(`'I'am the big man.'`)
	if s, is := s1.(string); !is || s != w1 {
		t.Error("string 解析失败， s, is := ", s, ", ", is)
	}
	// string 2
	w1 = "'I'am the big man.'"
	s1 = StrParseData(`''I'am the big man.''`)
	if s, is := s1.(string); !is || s != w1 {
		t.Error("string 解析失败， s, is := ", s, ", ", is)
	}

	// string 3
	w1 = "'I'am the big man.'"
	s1 = StrParseData(`"'I'am the big man.'"`)
	if s, is := s1.(string); !is || s != w1 {
		t.Error("string 解析失败， s, is := ", s, ", ", is)
	}

	// int64
	var w64 int64 = 14422337766
	i64 := StrParseData("14422337766")
	if it64, is := i64.(int64); !is || it64 != w64 {
		t.Error("int64 解析失败， it64, is, type := ", it64, ", ", is, ",", fmt.Sprintf("%T", it64))
	}

	// float64
	var wf64 float64 = 14422337766.124785
	f64 := StrParseData("14422337766.124785")
	if it64, is := f64.(float64); !is || f64 != wf64 {
		t.Error("int64 解析失败， it64, is, type := ", it64, ", ", is, ",", fmt.Sprintf("%T", f64))
	}

	// []int
	var wQi []int = []int{1, 6, 8, 100, 500, 9999}
	qi := StrParseData("1,6,8,100,500,9999")
	if ints, is := qi.([]int); !is || fmt.Sprintf("%v", wQi) != fmt.Sprintf("%v", qi) {
		t.Error("int64 解析失败， ints, is, type := ", ints, ", ", is, ",", fmt.Sprintf("%T", qi))
	}

	// []int
	//var wfi []float64 = []float64{1.0, 6.0, 3.45, 100.0, 500.0, 9999.0}
	var wfi []float64 = []float64{1, 6.3, 3.45, 100.8, 500.4, 9999.1}
	fi := StrParseData("1,6.3,3.45,100.8,500.4,9999.1")
	if f64s, is := fi.([]float64); !is || fmt.Sprintf("%v", wfi) != fmt.Sprintf("%v", fi) {
		t.Error("int64 解析失败， ints, is, type := ", f64s, ", ", is, ",", fmt.Sprintf("%T", fi))
	}

	// []string
	var wss []string = []string{"Don't", "give", "a", "damn", "shit,", "man"}
	ssi := StrParseData("Don't,give,a,damn,shit\\,,man")
	if f1, is := ssi.([]string); !is || fmt.Sprintf("%v", wss) != fmt.Sprintf("%v", f1) {
		t.Error("int64 解析失败， ints, is, type := ", f1, ", ", is, ",", fmt.Sprintf("%T", ssi))
	}
}
