package bin

import (
	"fmt"
	"gitee.com/conero/uymas/logger/lgr"
	"testing"
)

func TestNewCliCmd(t *testing.T) {
	command, subcommand := "git", "clone"
	var args []string
	var testKey []string
	var cc *Arg

	cc = NewCliCmd(command, subcommand, "https://gitee.com/conero/uymas.git")
	// case: `git clone https://gitee.com/conero/uymas.git`
	if cc.Command != "git" || cc.SubCommand != "clone" {
		t.Fatalf("the command parse fail. command: %v VS %v, subcommand: %v VS %v",
			cc.Command, command, cc.SubCommand, subcommand)
	}
	// case: `git -u gitee.com/conero/uymas`
	command = "get"
	cc = NewCliCmd(command, "-u", "gitee.com/conero/uymas")
	if cc.Command != command || cc.SubCommand != "" {
		t.Fatalf("the command parse fail. command: %v VS %v, subcommand: %v VS %v",
			cc.Command, command, cc.SubCommand, "")
	}

	//case
	//类型测试
	//uymas --is-bool True --is-string="Joshua Conero Test" --is-int64=202005 --is-float64 3.1415926535898
	command = "uymas"
	args = []string{command, "--is-bool", "True", "--is-string=\"Joshua Conero Test\"", "--is-int64=202005",
		"--is-float64", "3.1415926535898"}

	cc = NewCliCmd(args...)
	//类型检测
	t.Log(cc.Raw)
	t.Log(cc.DataRaw)
	t.Log(cc.Data)
	//dataRaw 兼键值检测
	testKey = []string{"is-bool", "is-string", "is-int64", "is-float64"}
	for _, tK := range testKey {
		//DataRaw
		if _, hasTk := cc.DataRaw[tK]; !hasTk {
			t.Fatalf("键值`%v`解析失败", tK)
		}

		// Data
		if _, hasTk := cc.Data[tK]; !hasTk {
			t.Fatalf("键值`%v`解析失败", tK)
		}
	}

}

func TestCLI_RegisterEnd(t *testing.T) {
	cli := NewCLI()
	cli.RegisterEnd(func() {
		fmt.Println("  ->Running event end :~~~~~")
		fmt.Println()
	})
	cli.RegisterFunc(func(arg *Arg) {
		fmt.Println(">>default")
		fmt.Println("Default empty.")
	})
	cli.RegisterFunc(func(arg *Arg) {
		fmt.Println(">>test")
		fmt.Println("test The testy command.")
	}, "test")

	// case
	cli.Run()
	// case
	cli.Run("test")
}

func TestCLI_RegisterFunc_Repeat(t *testing.T) {
	app := NewCLI()
	app.RegisterFunc(func(arg *Arg) {
		lgr.Info("it‘s C command")
	}, "repeat")
	app.RegisterFunc(func(arg *Arg) {
		lgr.Info("it‘s B command")
	}, "repeat")
}
