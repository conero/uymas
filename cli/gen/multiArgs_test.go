package gen

import (
	"gitee.com/conero/uymas/v2/cli"
	"testing"
)

type tmaStrcut struct {
	Test struct {
		Name string
		Age  int
	}
}

func TestMultiArgs(t *testing.T) {
	args := cli.NewArgs("-test.name", "Joshua", "-test.age", "18")
	var tma tmaStrcut
	err := MultiArgs(args, &tma)
	if err != nil {
		t.Errorf("解析异常，%v", err)
	} else if tma.Test.Age != 18 {
		t.Errorf("test.age 值复制错误")
	} else {
		t.Logf("data: %#v", tma)
	}
}
