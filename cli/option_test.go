package cli

import "testing"

// 全局选项命令获取
func TestCommandOptional_GenOptionHelpMsg_Global(t *testing.T) {
	co := CommandOptional{
		Options: []Option{
			Option{
				Name:     "verbose",
				Help:     "详细/冗余输出",
				IsGlobal: true,
			},
		},
	}

	helpMsg := co.GenOptionHelpMsg(true)
	if helpMsg == "" {
		t.Error("GenOptionHelpMsg() error")
	} else {
		t.Log(helpMsg)
	}
}
