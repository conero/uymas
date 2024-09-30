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
	}
}
