package cli

import "testing"

func TestNewArgs(t *testing.T) {
	arg := NewArgs("app", "--help")
	if arg.Command() != "app" {
		t.Errorf("app 解析命令错误")
	}
	if !arg.Switch("help") {
		t.Errorf("app 解析选项错误")
	}
}
