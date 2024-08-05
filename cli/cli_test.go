package cli

import "testing"

func TestNewCli(t *testing.T) {
	cl := NewCli()
	err := cl.Run("demo", "todo")
	if err != nil {
		t.Errorf("命令行解析错误，%v", err)
	}
}
