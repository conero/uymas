// Package butil bin util package
// will not run the init(), but bin will
package butil

import (
	"fmt"
	"gitee.com/conero/uymas/bin/parser"
	"gitee.com/conero/uymas/fs"
	"gitee.com/conero/uymas/util/rock"
	"os"
	"path"
	"strings"
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

// parse the current application
func parseCurrent(forces ...bool) {
	force := rock.ExtractParam(false, forces...)
	if current != nil && !force {
		return
	}

	rwd := os.Args[0]
	rwd = fs.StdPathName(rwd)
	vDir, vFile := path.Split(rwd)
	var nameNoSuffix string
	if vFile != "" {
		nameNoSuffix = strings.ReplaceAll(vFile, path.Ext(vFile), "")
	}

	current = &BinInfo{
		BaseDir:      vDir,
		Name:         vFile,
		NameNoSuffix: nameNoSuffix,
	}

}

// Basedir get application binary root dir.
func Basedir() string {
	if current == nil {
		parseCurrent(true)
	}
	if current.BaseDir != "" {
		return current.BaseDir
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

// AppName get current binary application name
func AppName() string {
	return current.NameNoSuffix
}

func AppFilename() string {
	return current.Name
}

// RootPath the path dir by application same location.
func RootPath(vPath string) string {
	return fmt.Sprintf("%v%v", Basedir(), vPath)
}

// DetectPath detect path by give relative or absolute path, can correct incorrect paths normally.
func DetectPath(vPath string) string {
	vLen := len(vPath)
	if vLen == 0 {
		return RootPath(vPath)
	}

	first := vPath[:2]
	if vLen > 1 && first == "./" {
		return vPath
	}

	// let path like `p1/px` is base exec bir
	if vPath[:1] != "/" {
		return RootPath(vPath)
	}

	if path.IsAbs(vPath) || fs.ExistPath(vPath) {
		return vPath
	}

	return RootPath(vPath)
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

// Current get the current application's `BinInfo`
func Current() BinInfo {
	return *current
}

func init() {
	parseCurrent()
}

// Pwd try to get the current working directory
func Pwd(joins ...string) string {
	vDir, err := os.Getwd()
	if err != nil {
		vDir = "./"
	}

	joins = append([]string{vDir}, joins...)
	return fs.StdPathName(path.Join(joins...))
}
