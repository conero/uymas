package fs

import (
	"errors"
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
		value = strings.Split(path, ";")
	}
	return value
}

// AddEnvPath and path to env
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
		err := os.Setenv(VEnvPath, strings.Join(refPath, ";"))
		// [TIP] 仅仅在当前进程中/实例中有效
		// fmt.Println(os.Getenv(V_EnvPath))
		return err
	}
	return errors.New("输入为空或者环境变量已经存在")
}

// DelEnvPath del the path from env path list.
func DelEnvPath(paths ...string) error {
	return nil
}
