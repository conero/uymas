package fs

import (
	"os"
	"path"
)

var (
	gRootPath string
	gRootApp  string
	gRootFile string
)

func parseRoot() {
	rwd := os.Args[0]
	gRootFile = StdPathName(rwd)
	gRootPath, _, gRootApp = Split(gRootFile)
}

// Split Decompose the path into base directory, file name, and file name without suffix
func Split(vPath string) (baseDir, file, name string) {
	rootPath := StdPathName(vPath)
	baseDir, file = path.Split(rootPath)
	baseDir = StdDir(baseDir)
	ext := path.Ext(file)
	name = file
	if ext != "" {
		name = name[:len(name)-len(ext)]
	}
	return
}

// RootPath Get the directory where the application is located
func RootPath(joins ...string) string {
	if gRootPath == "" {
		parseRoot()
	}

	joins = append([]string{gRootPath}, joins...)
	return path.Join(joins...)
}

// RunDir Prioritize obtaining the current working directory,
// and try to obtain the directory where the application is located after failure,
// ensuring that the directory exists as much as possible.
func RunDir(joins ...string) string {
	pwd, err := os.Getwd()
	if err == nil {
		joins = append([]string{pwd}, joins...)
		return path.Join(joins...)
	}
	return RootPath(joins...)
}

// AppName Get the name of the current application without any suffix
func AppName() string {
	if gRootApp == "" {
		parseRoot()
	}
	return gRootApp
}

// RootFile Get the full path of the current application
func RootFile() string {
	if gRootFile == "" {
		parseRoot()
	}
	return gRootFile
}
