//Package svn command parser.
package svn

import (
	"bytes"
	"os/exec"
)

// @Date：   2018/12/5 0005 22:25
// @Author:  Joshua Conero
// @Name:    Svn 包(基于 svn 命令)
const (
	CliName     = "svn"   // 命令名称
	BaseVersion = "1.8.3" // 依赖版本信息
)

// 获取版本信息
func Version() string {
	c := exec.Command(CliName, "--version", "--quiet")
	var out bytes.Buffer
	c.Stdout = &out
	if err := c.Run(); err != nil {
		return ""
	}
	return out.String()
}

// 调用命令
func Call(args ...string) (string, error) {
	c := exec.Command(CliName, args...)
	var out bytes.Buffer
	c.Stdout = &out
	if err := c.Run(); err != nil {
		return "", err
	}
	return out.String(), nil
}
