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
	Version = "2.0.0" //dev is not product but development tag.
	Release = "dev"   // dev|20060102
	Since   = "20181030"
	Author  = "Joshua Conero"
	Email   = "conero@163.com"
	Name    = "uymas"
	PkgName = "conero/uymas/v2"
)

var (
	gitHash     string
	buildData   string
	buildAuthor string
)

// GetBuildInfo Get the relevant version information after injection with the -ldflags parameters gitHash, buildData, buildAuthor.
//
// shell like:
//
// $gitHash = $(git rev-parse --short HEAD); $buildData = $(get-date -Format 'yyyy-MM-dd');$buildAuthor = 'Joshua Conero';
//
// go build -ldflags "-s -w -X 'gitee.com/conero/uymas/v2.gitHash=$gitHash' -X 'gitee.com/conero/uymas/v2.buildData=$buildData' -X 'gitee.com/conero/uymas/v2.buildAuthor=$buildAuthor'" ./cmd/...
//
// Output format such as: "(buildData gitHash)"
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
