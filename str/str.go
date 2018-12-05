package str

import (
	"regexp"
	"strings"
)

// @Date：   2018/10/30 0030 15:14
// @Author:  Joshua Conero
// @Name:    字符串

// 首字母大写
func Ucfirst(str string) string {
	idx := strings.Index(str, " ")
	if idx > -1 {
		newStr := []string{}
		for _, s := range strings.Split(str, " ") {
			newStr = append(newStr, Ucfirst(s))
		}
		str = strings.Join(newStr, "")
	} else {
		if len(str) > 0 {
			str = strings.ToUpper(str[0:1]) + str[1:]
		}
	}
	return str
}

// 首字母小写
func Lcfirst(str string) string {
	idx := strings.Index(str, " ")
	if idx > -1 {
		newStr := []string{}
		for _, s := range strings.Split(str, " ") {
			newStr = append(newStr, Lcfirst(s))
		}
		str = strings.Join(newStr, "")
	} else {
		if len(str) > 0 {
			str = strings.ToLower(str[0:1]) + str[1:]
		}
	}
	return str
}

// 安全字符串分割
func SplitSafe(s, sep string) []string  {
	var dd []string
	s = ClearSpace(s)
	dd = strings.Split(s, sep)
	return dd
}

// 清除空格
func ClearSpace(s string) string {
	s = strings.TrimSpace(s)
	if strings.Index(s, " ") > -1{
		spaceReg := regexp.MustCompile("\\s")
		s = spaceReg.ReplaceAllString(s, "")
	}
	return s
}