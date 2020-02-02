package bin

import "testing"

func TestNewCliCmd(t *testing.T) {
	command, subcommand := "git", "clone"
	var cc *CliCmd

	cc = NewCliCmd(command, subcommand, "https://github.com/conero/uymas.git")
	// case
	if cc.Command != "git" || cc.SubCommand != "clone" {
		t.Fatalf("the command parse fail. command: %v VS %v, subcommand: %v VS %v",
			cc.Command, command, cc.SubCommand, subcommand)
	}
	// case
	command = "get"
	cc = NewCliCmd(command, "-u", "github.com/conero/uymas")
	if cc.Command != command || cc.SubCommand != "" {
		t.Fatalf("the command parse fail. command: %v VS %v, subcommand: %v VS %v",
			cc.Command, command, cc.SubCommand, "")
	}
}
