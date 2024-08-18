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

func RootPath(joins ...string) string {
	if gRootPath == "" {
		parseRoot()
	}
	return gRootPath
}

func AppName() string {
	if gRootApp == "" {
		parseRoot()
	}
	return gRootApp
}

func RootFile() string {
	if gRootFile == "" {
		parseRoot()
	}
	return gRootFile
}
