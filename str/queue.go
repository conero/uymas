package str

import "strings"

// @Date：   2018/11/7 0007 11:38
// @Author:  Joshua Conero
// @Name:    字符互队列

/**
字符串对是否存在
不存在返回 -1
*/
func InQue(s string, que []string) int {
	idx := -1
	for i, v := range que {
		if s == v {
			idx = i
			break
		}
	}
	return idx
}

// 不区分大小写
func InQuei(s string, que []string) int {
	idx := -1
	s = strings.ToLower(s)
	for i, v := range que {
		if s == strings.ToLower(v) {
			idx = i
			break
		}
	}
	return idx
}

// 删除队列
func DelQue(que []string, ss ...string) []string {
	var value []string
	if que != nil && ss != nil {
		for _, s := range que {
			if InQue(s, ss) == -1 {
				if value == nil {
					value = []string{}
				}
				value = append(value, s)
			}
		}
	}
	return value
}
