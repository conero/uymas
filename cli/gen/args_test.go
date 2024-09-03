package gen

import (
	"encoding/json"
	"fmt"
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/rock"
	"gitee.com/conero/uymas/v2/str"
	"testing"
)

type dressData struct {
	Off        bool   `cmd:"off,O,close required help:请求数据关闭"`
	Name       string `cmd:"name,N"`
	Score      float32
	Age        int
	SupportExt []string
	Rates      []float32 `cmd:"rates,rate,R"`
	Data       []string  `cmd:"data,d required help:输入数组列表，\\s\\s支持列表"`
}

func TestArgsDress(t *testing.T) {
	vStr := "Joshua Conero"
	var vF32 float32 = 75

	cmdMock := []string{"uymas", "test", "-N", vStr, "--score", fmt.Sprintf("%.2f", vF32), "-O"}
	args := cli.NewArgs(cmdMock...)
	//t.Logf("command: %v\nargs:%v", cmdMock, args)
	var data dressData
	err := ArgsDress(args, &data)
	if err != nil {
		t.Errorf("解析值异常，%v", err)
	} else if data.Name != vStr {
		t.Errorf("字符串赋值失败，%s ≠ %s", data.Name, vStr)
	} else if data.Score != vF32 {
		t.Errorf("float32 赋值失败，%v ≠ %v", data.Score, vF32)
	} else if !data.Off {
		t.Errorf("bool 赋值失败，%v ≠ %v", data.Score, vF32)
	}

}

func TestArgsDecompose(t *testing.T) {
	optionsList, err := ArgsDecompose(dressData{})
	if err != nil {
		t.Errorf("Args 解析失败错误，%v", err)
	} else if optionsList == nil {
		t.Errorf("Args 解析数据为空")
	} else {
		bys, _ := json.Marshal(optionsList)
		t.Logf("解析后的数据：\n%s", string(bys))
		t.Logf("%#v", optionsList[len(optionsList)-1])
	}
}

func TestOptionTagParse(t *testing.T) {
	help := `输入用户指定命令`
	name := `Joshua\sConero`
	refDefault := str.Str(name).Unescape()
	vTag := `name,n required help:` + help + ` default:` + name

	// case 1
	opt := OptionTagParse(vTag)
	if opt == nil {
		t.Errorf("tag 解析失败")
	} else {
		if opt.Help != help {
			t.Errorf("help 解析失败")
		}
		if opt.DefValue != refDefault {
			t.Errorf("name 解析失败，%#v", opt.DefValue)
		}
		if !rock.ListEq(opt.Alias, []string{"name", "n"}) {
			t.Errorf("name 命令解析失败，%#v", opt.Alias)
		}
	}
}
