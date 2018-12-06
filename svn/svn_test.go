package svn

import (
	"fmt"
	"testing"
)

// @Date：   2018/12/6 0006 13:41
// @Author:  Joshua Conero
// @Name:    名称描述

func TestVersion(t *testing.T) {
	fmt.Println(Version())
}

func TestCall(t *testing.T) {
	// version
	out, er := Call("--version", "--quiet")
	if er != nil {
		t.Log(er.Error())
		t.Fail()
	}
	ver := Version()
	if ver != out {
		t.Log(out, ver)
		t.Fail()
	}
}
