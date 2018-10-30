package util

// @Date：   2018/10/30 0030 13:26
// @Author:  Joshua Conero
// @Name:    工具栏

/**
数组中是否存在
不存在返回 -1
*/
func InQue(val interface{}, que []interface{}) int {
	idx := -1
	if que != nil {
		for i, v := range que {
			if v == val {
				idx = i
				break
			}
		}
	}
	return idx
}

/**
字符串对是否存在
不存在返回 -1
*/
func InStrQue(s string, que []string) int {
	idx := -1
	for i, v := range que {
		if s == v {
			idx = i
			break
		}
	}
	return idx
}
