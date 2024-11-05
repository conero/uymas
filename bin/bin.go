// Package bin is sample command application lib, provides functional and classic style Apis.
package bin

// @Date：   2018/10/30 0030 13:20
// @Author:  Joshua Conero

import (
	"gitee.com/conero/uymas/str"
	"regexp"
	"strings"
)

// the Cmd of type
const (
	CmdApp int = iota
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
			words = append(words, str.Ucfirst(v))
		}
	}

	return strings.Join(words, "")
}
