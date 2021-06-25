package parser

import (
	"testing"
)

func TestNewParser(t *testing.T) {
	script := "build --name 'name.py';run -tsz" + `
	test -xyz										# Joshua Conero
	tar ./target/dir -o target.zip
	echo 'this is a good idea'
	# 支持代码注释，单行注释
	`
	tmpArr := NewParser(script)
	t.Log(tmpArr)
}
