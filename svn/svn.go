package svn

import (
	"bytes"
	"os/exec"
)

// @Date：   2018/12/5 0005 22:25
// @Author:  Joshua Conero
// @Name:    Svn 包(基于 svn 命令)
const (
	CliName = "svn"
)

func Version() string {
	c := exec.Command(CliName, "--version", "--quiet")
	var out bytes.Buffer
	c.Stdout = &out
	if err := c.Run(); err != nil{
		return ""
	}
	return out.String()
}