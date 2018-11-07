package fs

import (
	"github.com/conero/uymas/util"
	"github.com/conero/uymas/str"
	"os"
	"strings"
)

// @Date：   2018/11/7 0007 11:27
// @Author:  Joshua Conero
// @Name:    操作系统相关的数据读写

const (
	V_EnvPath = "path"
)

// 获取项目路径
func EnvPath() []string {
	value := []string{}
	path := os.Getenv(V_EnvPath)
	if path != "" {
		value = strings.Split(path, ";")
	}
	return value
}

// 设置环境变量
func AddEnvPath(paths ...string) error {
	refPath := EnvPath()
	addMk := false
	for _, path := range paths {
		if str.InQuei(path, refPath) == -1 {
			refPath = append(refPath, path)
			addMk = true
		}
	}
	if addMk {
		err := os.Setenv(V_EnvPath, strings.Join(refPath, ";"))
		return err
	}
	return &util.BaseError{"输入为空或者环境变量已经存在"}
}

// 删除环境变量
func DelEnvPath(paths ...string) error {
	return nil
}
