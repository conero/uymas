// Package butil bin util package
// will not run the init(), but bin will
package butil

import (
	"fmt"
	"gitee.com/conero/uymas/bin/parser"
	"gitee.com/conero/uymas/fs"
	"os"
	"path"
)

// current application by parse binary.
type application struct {
	baseDir string
	name    string
}

var (
	current *application
)

// parse the current application
func parseCurrent(force bool) {
	if current != nil && !force {
		return
	}

	rwd := os.Args[0]
	rwd = fs.StdPathName(rwd)
	vDir, vFile := path.Split(rwd)
	current = &application{
		baseDir: vDir,
		name:    vFile,
	}

}

// Deprecated: get the root base dir, will rename to `Basedir()`
func GetBasedir() string {
	return Basedir()
}

// Basedir get application binary root dir.
func Basedir() string {
	return current.baseDir
}

// AppName get current binary application name
func AppName() string {
	return current.name
}

// Deprecated: the path dir by application same location, please replace use  `RootPath`.
func GetPathDir(vPath string) string {
	return fmt.Sprintf("%v%v", Basedir(), vPath)
}

// RootPath the path dir by application same location.
func RootPath(vPath string) string {
	return fmt.Sprintf("%v%v", Basedir(), vPath)
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

func init() {
	parseCurrent(false)
}
