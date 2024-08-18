// Package butil bin util package
// will not run the init(), but bin will
package butil

import (
	"gitee.com/conero/uymas/v2/bin/parser"
	"gitee.com/conero/uymas/v2/util/fs"
	"os"
	"path"
)

// BinInfo by parse binary.
type BinInfo struct {
	BaseDir      string
	Name         string
	NameNoSuffix string
}

var (
	current *BinInfo
)

// Basedir get application binary root dir.
func Basedir() string {
	baseDir := fs.RootPath()
	if baseDir != "" {
		return baseDir
	}
	// Notice: When the system is running in a cmd environment,
	// it may not be possible to obtain the current directory.
	// Therefore, at this point, read the cwd of the current running environment
	cwd, err := os.Getwd()
	if err == nil {
		basedir := fs.StdPathName(cwd + "/")
		return basedir
	}
	return "./"
}

// DetectPath detect path by give relative or absolute path, can correct incorrect paths normally.
func DetectPath(vPath string) string {
	vLen := len(vPath)
	if vLen == 0 {
		return fs.RootPath(vPath)
	}

	first := vPath[:2]
	if vLen > 1 && first == "./" {
		return vPath
	}

	// let path like `p1/px` is base exec bir
	if vPath[:1] != "/" {
		return fs.RootPath(vPath)
	}

	if path.IsAbs(vPath) || fs.ExistPath(vPath) {
		return vPath
	}

	return fs.RootPath(vPath)
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
