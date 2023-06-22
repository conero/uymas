package bin

import (
	"strings"
	"testing"
)

func TestArg_NextList(t *testing.T) {
	// case 1
	cmdStr := "$ find *.php *.js *.css *.go -r --verbose 1 3 4 5"
	arg := NewCliCmdByString(cmdStr)

	want := "find *.php *.js *.css *.go"
	affect := strings.Join(arg.NextList(), " ")
	if affect != want {
		t.Errorf("%v 解析错误 \n  => %v", cmdStr, affect)
	}

	// case2
	want = "*.css *.go"
	affect = strings.Join(arg.NextList("*.js"), " ")
	if affect != want {
		t.Errorf("%v 解析错误 \n  => %v", cmdStr, affect)
	}

	// case3
	want = "*.php *.js *.css *.go"
	affect = strings.Join(arg.NextList("find"), " ")
	if affect != want {
		t.Errorf("%v 解析错误 \n  => %v", cmdStr, affect)
	}

	// case3
	want = ""
	affect = strings.Join(arg.NextList("x-find"), " ")
	if affect != want {
		t.Errorf("%v 解析错误 \n  => %v", cmdStr, affect)
	}
}
