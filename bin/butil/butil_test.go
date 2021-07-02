package butil

import (
	"os"
	"reflect"
	"testing"
)

func TestRootDir(t *testing.T) {
	t.Log(GetBasedir())
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
