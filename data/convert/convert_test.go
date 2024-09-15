package convert

import (
	"reflect"
	"testing"
)

func TestSetByStrSlice(t *testing.T) {
	var vI64 []int64
	vStr := []string{"187", "1893", "100", "61"}
	SetByStrSlice(reflect.ValueOf(&vI64), vStr)

	// case 1
	if len(vI64) != len(vStr) {
		t.Errorf("字符串转 int64 错误：%v, 格式不匹配", vI64)
	} else {
		t.Logf("设置的值：%#v", vI64)
	}

	// case
	vStr = []string{"2024", "09", "15"}
	SetByStrSlice(reflect.ValueOf(&vI64), vStr)
	if len(vI64) != len(vStr) {
		t.Errorf("字符串转 int64 错误：%v, 格式不匹配", vI64)
	} else {
		t.Logf("设置的值：%#v", vI64)
	}

}
