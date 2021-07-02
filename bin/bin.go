// @Date：   2018/10/30 0030 13:20
// @Author:  Joshua Conero

// Package bin is sample command application lib, provides functional and classic style Apis.
package bin

const (
	AppMethodInit     = "Init"
	AppMethodRun      = "Run"
	AppMethodNoSubC   = "SubCommandUnfind"
	AppMethodHelp     = "Help"
	FuncRegisterEmpty = "_inner_empty_func"
)

type initIota int

//the Cmd of type
const (
	CmdApp initIota = iota
	CmdFunc
)
