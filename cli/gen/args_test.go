package gen

import (
	"fmt"
	"gitee.com/conero/uymas/v2/cli"
	"testing"
)

func TestArgsDress(t *testing.T) {
	type dressData struct {
		Off        bool   `cmd:"off,O,close"`
		Name       string `json:"name,N"`
		Score      float32
		Age        int
		SupportExt []string
		Rates      []float32 `json:"rates,rate,R"`
	}

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
