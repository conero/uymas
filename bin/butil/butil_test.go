package butil

import (
	"gitee.com/conero/uymas/v2/util/fs"
	"os"
	"reflect"
	"testing"
)

func TestRootDir(t *testing.T) {
	t.Log(Basedir())
	t.Log(os.Args)
}

func TestStringToArgs(t *testing.T) {
	var (
		line    string
		args    []string
		refArgs []string
	)

	//case1
	line = "git commit    -m 'Is a test comment,\\' good do it.'"
	args = StringToArgs(line)
	refArgs = []string{"git", "commit", "-m", "'Is a test comment,\\' good do it.'"}
	if !reflect.DeepEqual(args, refArgs) {
		t.Fatalf("line <%v> parse args is bad\n   --> %#v", line, args)
	}

	//case2
	line = "uymas --load-json '{\"country\":\"China\",\"New China National Day\":\"1949.10.1\",\"n-year\":1949}'"
	args = StringToArgs(line)
	refArgs = []string{"uymas", "--load-json", `'{"country":"China","New China National Day":"1949.10.1","n-year":1949}'`}
	if !reflect.DeepEqual(args, refArgs) {
		t.Fatalf("line <%v> parse args is bad\n   --> %#v", line, args)
	}
}

func TestDetectPath(t *testing.T) {
	// case 1
	ipt, want := "./logs/230702", "./logs/230702"
	if realStr := DetectPath(ipt); realStr != want {
		t.Errorf("请求错误, %v != %v", realStr, want)
	}

	// case 2
	ipt, want = "logs/230702", fs.RootPath("logs/230702")
	if realStr := DetectPath(ipt); realStr != want {
		t.Errorf("请求错误, %v != %v", realStr, want)
	}

	// case 3
	ipt, want = "", fs.RootPath("")
	if realStr := DetectPath(ipt); realStr != want {
		t.Errorf("请求错误, %v != %v", realStr, want)
	}

	// case 3
	ipt, want = ".gitignore", fs.RootPath(".gitignore")
	if realStr := DetectPath(ipt); realStr != want {
		t.Errorf("请求错误, %v != %v", realStr, want)
	}
}
