package util

import (
	"math"
	"strings"
	"time"
)

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

// 不区分大小写
func InStrQuei(s string, que []string) int {
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


//返回秒用于计算程序用时,参数为0时返回当前的毫秒，否则返回计算后的秒差
func Sec(start float64) float64 {
	t := time.Now()
	ns := float64(t.Nanosecond())
	ms := ns / math.Pow10(6) //1ms = 10^6ns
	if start == 0 {
		return ms
	}
	ds := (ms - start) / math.Pow10(3)
	ds = Round(ds, 5)
	return ds
}

//字符串方法处理float等长数据 规定位数
func Round(num float64, b int) float64 {
	if b == 0 {
		return float64(int(num))
	}
	n2t := int(num * math.Pow10(b))    //num转换数
	base := int(num * math.Pow10(b+1)) //四舍五入的最后一位数
	base = int(math.Abs(float64(base - n2t*10)))
	if base > 5 {
		n2t += 1
	}
	num = float64(int(num)) + float64(n2t)/float64(math.Pow10(b))
	return num
}