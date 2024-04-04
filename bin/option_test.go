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
		opt.ExcludeReg(`^config.*`)

		if cc.CheckSetting("display") && cc.ArgRaw("display") != "Joshua" {
			t.Errorf("display 选项解析失败")
		}
		if cc.CheckSetting("version", "x", "exclude", "config") {
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
	cli.Run("--config.level", "Joshua", "--config.log", "m", "--config.dev", "true", "--config", "--exclude-reg")
}
