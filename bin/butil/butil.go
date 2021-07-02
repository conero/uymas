// Package butil bin util package
// will not run the init(), but bin will
package butil

import (
	"fmt"
	"github.com/conero/uymas/bin/parser"
	"github.com/conero/uymas/fs"
	"os"
	"path"
)

var (
	cacheBaseDir string
)

// GetBasedir get the root base Dir
func GetBasedir() string {
	if cacheBaseDir != "" {
		return cacheBaseDir
	}
	rwd := os.Args[0]
	rwd = fs.StdPathName(rwd)
	rwd = path.Dir(rwd)
	cacheBaseDir = rwd
	return rwd
}

// GetPathDir the path dir by application same location.
func GetPathDir(vPath string) string {
	return fmt.Sprintf("%v%v", GetBasedir(), vPath)
}

// StringToArgs make the string to bin/Args, it's used in interactive cli
func StringToArgs(str string) []string {
	args := parser.ParseLine(str)
	if len(args) > 0 {
		return args[0]
	}
	return nil
}

// StringToMultiArgs string line parse multi line, support ";" split.
func StringToMultiArgs(str string) [][]string {
	return parser.ParseLine(str)
}
