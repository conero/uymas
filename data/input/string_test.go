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
	vInt = Stringer("1982.2344").Uint32()
	if vInt != 1982 {
		t.Errorf("%d != 1982", vInt)
	}

	// case - 64位数据测试
	vInt = Stringer("1234567890123456789").Uint32()
	if vInt != math.MaxUint32 {
		t.Errorf("%d != 0", vInt)
	}
}

func TestStringer_Int(t *testing.T) {
	vInt := Stringer("123").Int()
	if vInt != 123 {
		t.Errorf("%d != 123", vInt)
	}

	// case
	vInt = Stringer("-100").Int()
	if vInt != -100 {
		t.Errorf("%d != -100", vInt)
	}

	// case
	vInt = Stringer("-100.123").Int()
	if vInt != -100 {
		t.Errorf("%d != -100", vInt)
	}

	// case
	vInt = Stringer("3.14").Int()
	if vInt != 3 {
		t.Errorf("%d != 3", vInt)
	}

	// case
	vInt = Stringer("12_000").Int()
	if vInt != 12000 {
		t.Errorf("%d != 12000", vInt)
	}
}

func TestStringer_Float(t *testing.T) {
	vFloat := Stringer("123.123").Float()
	if vFloat != 123.123 {
		t.Errorf("%f != 123.123", vFloat)
	}

	//  case
	vFloat = Stringer("-100.123").Float()
	if vFloat != -100.123 {
		t.Errorf("%f != -100.123", vFloat)
	}

	// case
	vFloat = Stringer("3").Float()
	if vFloat != 3 {
		t.Errorf("%f != 3", vFloat)
	}
}
