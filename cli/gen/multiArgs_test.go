package gen

import (
	"gitee.com/conero/uymas/v2/cli"
	"testing"
)

type tmaStrcut struct {
	Test struct {
		Name string
		Age  int
		Keys []string
		Cs   struct {
			Name    string
			Rate    float64
			IsLock  bool
			Channel struct {
				Name    string
				IsAdmin bool
			}
		}
	}
}

func TestMultiArgs(t *testing.T) {
	// case
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

	// case
	args = cli.NewArgs("-test:name", "Elem Sr", "-test:age", "92")
	err = MultiArgs(args, &tma, ":")
	if err != nil {
		t.Errorf("解析异常，%v", err)
	} else if tma.Test.Age != 92 {
		t.Errorf("test.age 值复制错误")
	} else {
		t.Logf("data: %#v", tma)
	}
}

func TestMultiArgsMap(t *testing.T) {
	args := cli.NewArgs("-test.name", "Joshua",
		"-test.age", "18", "-test.child.source", "Guizhou")
	vmap := map[string]any{}

	err := MultiArgsMap(args, vmap)
	if err != nil {
		t.Errorf("map 赋值异常，%s", err.Error())
	} else {
		t.Logf("%#v", vmap)
	}

	// case
	args = cli.NewArgs("-jc:name", "Joshua",
		"-jc:age", "18", "-jc:server:source", "Guizhou", "-jc:db:host", "-jc:db:host", "127.0.0.1", "-jc:db:port", "2024")
	vmap = map[string]any{}

	err = MultiArgsMap(args, vmap, ":")
	if err != nil {
		t.Errorf("map 赋值异常，%s", err.Error())
	} else {
		t.Logf("%#v", vmap)
	}

}
