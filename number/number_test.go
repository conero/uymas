package number

import "testing"

// @Date：   2018/12/20 0020 16:36
// @Author:  Joshua Conero
// @Name:    数据测试

func TestSumQueue(t *testing.T) {
	var axp, out any
	var in []any
	// int
	in = []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	axp = 55
	out = SumQueue(in)
	if axp != out {
		t.Fail()
	}

	// float 32
	in = []any{1.2, 2.3, 3.4, 4.5, 5.6, 6.7, 7.8, 8.9, 9.1, 10.11}
	axp = 59.61
	out = SumQueue(in)
	if axp != out {
		t.Fail()
	}
}
