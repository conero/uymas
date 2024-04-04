package bin

import (
	"testing"
)

func TestOption_Unmarshal(t *testing.T) {
	type base struct {
		Name        string
		DisplayName string `arg:"d, display"`
	}

	cli := NewCLI()
	cli.RegisterEmpty(func(cc *Arg) {
		opt := &Option{cc: cc}
		var bv base
		opt.Unmarshal(&bv)
		opt.Exclude("exclude")

		if cc.CheckSetting("display") && cc.ArgRaw("display") != "Joshua" {
			t.Errorf("display 选项解析失败")
		}
		if cc.CheckSetting("version", "x", "exclude") {
			if err := opt.CheckAllow(); err != nil {
				t.Logf("%v", err)
			} else {
				t.Error("CheckAllow invalid")
			}
		}
		t.Logf("Option: %#v\n", bv)
		t.Logf("Arg Data: %#v\n", cc.Data)
	})

	cli.Run("-d", "Joshua", "--name", "xyz")
	cli.Run("--display", "Joshua", "--name", "xyz", "--version", "-x")

	//设置排除选项
	cli.Run("--exclude", "Joshua", "ju", "m", "-g", "--grep", "*x")
}
