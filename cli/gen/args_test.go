package gen

import (
	"encoding/json"
	"fmt"
	"gitee.com/conero/uymas/v2/cli"
	"testing"
)

type dressData struct {
	Off        bool   `cmd:"off,O,close required help:请求数据关闭"`
	Name       string `json:"name,N"`
	Score      float32
	Age        int
	SupportExt []string
	Rates      []float32 `json:"rates,rate,R"`
	Data       []string  `cmd:"data,d required help:输入数组列表，支持列表"`
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
	}
}
