package fs

import (
	"gitee.com/conero/uymas/v2/str"
	"os"
	"strings"
)

// @Date：   2018/11/7 0007 11:27
// @Author:  Joshua Conero
// @Name:    操作系统相关的数据读写

const (
	VEnvPath = "path"
)

// EnvPath get os env path list
func EnvPath() []string {
	var value []string
	path := os.Getenv(VEnvPath)
	if path != "" {
		path = strings.ReplaceAll(path, "\\", "/")
		value = strings.Split(path, ";")
	}
	return value
}

// AddEnvPath and path to env
func AddEnvPath(paths ...string) string {
	paths = StdPathList(paths...)
	refPath := EnvPath()
	for _, path := range paths {
		if str.InQuei(path, refPath) == -1 {
			refPath = append(refPath, path)
		}
	}
	return strings.Join(refPath, ";")
}

// DelEnvPath del the path from env path list.
func DelEnvPath(paths ...string) string {
	var newPath []string
	paths = StdPathList(paths...)
	for _, pth := range EnvPath() {
		if str.InQuei(pth, paths) != -1 {
			continue
		}
		newPath = append(newPath, pth)
	}
	return strings.Join(newPath, ";")
}

// StdPathList turn path list to standard path
func StdPathList(paths ...string) []string {
	var newPath []string
	for i, path := range paths {
		paths[i] = StdPathName(path)
	}
	return newPath
}
