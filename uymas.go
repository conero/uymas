// Package uymas is summary util library from the work experience,
// provides base version, author information and so on.
//
// Major functional like `bin` (CLI-APP), `str`(string util), other more. Some of originated from work experience
// to enhance code reuse.
// The final directory is a convenient tool to realize command-line program development and other code.
package uymas

// @Date：   2018/10/30 0030 12:58
// @Author:  Joshua Conero

const (
	Version        = "1.4.2"    // dev is not product but development tag.
	Release        = "20241020" // dev|20060102
	Since          = "20181030"
	Author         = "Joshua Conero"
	Email          = "conero@163.com"
	Name           = "uymas"
	PkgName        = "conero/uymas"
	TimeLayoutDate = "20060102" // date layout-20060102
)

// Data for injection
// # go build -ldflags "-s -w -X 'gitee.com/conero/uymas.gitHash=$gitHash' -X 'gitee.com/conero/uymas.buildData=$buildData' -X 'gitee.com/conero/uymas.buildAuthor=$buildAuthor'" ./cmd/...
var (
	gitHash     string
	buildData   string
	buildAuthor string
)

func GetBuildInfo() string {
	info := ""
	if gitHash != "" && buildData != "" {
		info = "(" + buildData + " " + gitHash + ")"
	}
	if buildAuthor != "" {
		if info != "" {
			info += "    "
		}
		info += "Power by " + buildAuthor
	}
	return info
}
