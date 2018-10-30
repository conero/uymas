package str

import "strings"

// @Date：   2018/10/30 0030 15:14
// @Author:  Joshua Conero
// @Name:    字符串

func Ucfirst(str string) string  {
	idx := strings.Index(str, " ")
	if idx > -1{
		newStr := []string{}
		for _, s := range strings.Split(str, " "){
			newStr = append(newStr, Ucfirst(s))
		}
		str = strings.Join(newStr, "")
	}else{
		if len(str) > 0{
			str = strings.ToUpper(str[0:1]) + str[1:]
		}
	}
	return str
}
