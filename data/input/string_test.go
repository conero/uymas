package input

import (
	"math"
	"testing"
)

func TestStringer_Uint32(t *testing.T) {
	vInt := Stringer("123").Uint32()
	if vInt != 123 {
		t.Error("vInt != 123")
	}

	// case
	vInt = Stringer("-100").Uint32()
	if vInt != 0 {
		t.Error("vInt != 0")
	}

	// case
	// @todo 兼容 float
	//vInt = Stringer("1982.2344").Uint32()
	//if vInt != 1982 {
	//	t.Errorf("%d != 1982", vInt)
	//}

	// case - 64位数据测试
	vInt = Stringer("1234567890123456789").Uint32()
	if vInt != math.MaxUint32 {
		t.Errorf("%d != 0", vInt)
	}
}
