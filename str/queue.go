package str

import (
	"gitee.com/conero/uymas/v2/rock"
	"strings"
)

// @Date：   2018/11/7 0007 11:38
// @Author:  Joshua Conero
// @Name:    字符互队列

// InQuei checkout substring exist in array that case insensitive
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

// DelQue del queue from array
func DelQue(que []string, ss ...string) []string {
	var value []string
	if que != nil && ss != nil {
		for _, s := range que {
			if rock.ListIndex(ss, s) == -1 {
				if value == nil {
					value = []string{}
				}
				value = append(value, s)
			}
		}
	}
	return value
}

// StrQueueToAny string slice convert to nany slice
func StrQueueToAny(args []string) []any {
	var anyQueue []any
	for _, s := range args {
		anyQueue = append(anyQueue, s)
	}
	return anyQueue
}
