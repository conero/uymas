// @Date：   2018/10/30 0030 13:20
// @Author:  Joshua Conero

// Package bin is sample command application lib, provides functional and classic style Apis.
package bin

import (
	"regexp"
	"strings"
)

const (
	AppMethodInit     = "Init"
	AppMethodRun      = "Run"
	AppMethodNoSubC   = "SubCommandUnfind"
	AppMethodHelp     = "Help"
	FuncRegisterEmpty = "_inner_empty_func"
)

type initIota int

// the Cmd of type
const (
	CmdApp initIota = iota
	CmdFunc
)

// Cmd2StringMap command string turn to map string, for standard go method name.
// like:
//
//	`get-videos` -> `GetVideos`
//	`get_videos` -> `GetVideos`
func Cmd2StringMap(c string) string {
	reg := regexp.MustCompile(`([-_]+)|(\s{2,})`)
	c = reg.ReplaceAllString(c, " ")

	var words []string
	for _, v := range strings.Split(c, " ") {
		if v != "" {
			words = append(words, strings.Title(v))
		}
	}

	return strings.Join(words, "")
}
