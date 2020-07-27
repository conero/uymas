package xini

import (
	"bufio"
	"strings"
)

// @Date：   2018/8/19 0019 10:57
// @Author:  Joshua Conero
// @Name:    字符串解析器

// 字符串解析器
type StrParser interface {
	Line() int
	GetData() map[interface{}]interface{}
	LoadContent(content string) StrParser
}

// 字符串遍历行
func str2lines(content string, callback func(line string)) {
	buf := bufio.NewReader(strings.NewReader(content))
	for {
		line, err2 := buf.ReadString('\n')
		callback(line)
		// 错误
		if err2 != nil {
			break
		}
	}
}
