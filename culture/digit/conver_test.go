package digit

import "testing"

func TestCover_ToChnUpper(t *testing.T) {
	var test Cover

	test = 6
	test.ToChnUpper()

	test = 9001
	test.ToChnUpper()

	test = 9070
	test.ToChnUpper()

	test = 98_710_016
	test.ToChnUpper()

	test = 105_070_401
	test.ToChnUpper()

	test = 10_005_070_401
	test.ToChnUpper()

	test = 2391792.0872
	test.ToChnUpper()
}

// 参考工具： http://daxie.gjcha.com/
func TestCover_ToChnRoundUpper(t *testing.T) {
	var value Cover
	var ref, actl string

	// case
	value = 999_999_999_999_999_999
	t.Logf("cover: %v => %v", value, value.ToChnRoundUpper())

	// case
	value = 9
	ref = "玖"
	actl = value.ToChnRoundUpper()
	if actl != ref {
		t.Fatalf("值 [%v], 转换后 [%v]; 与参考值 [%v] 不符合！", value, actl, ref)
	}

	// case
	value = 98
	ref = "玖拾捌"
	actl = value.ToChnRoundUpper()
	if actl != ref {
		t.Fatalf("值 [%v], 转换后 [%v]; 与参考值 [%v] 不符合！", value, actl, ref)
	}

	// case
	value = 185
	ref = "壹佰捌拾伍"
	actl = value.ToChnRoundUpper()
	if actl != ref {
		t.Fatalf("值 [%v], 转换后 [%v]; 与参考值 [%v] 不符合！", value, actl, ref)
	}

	// 9001
	value = 9001
	ref = "玖仟零壹"
	actl = value.ToChnRoundUpper()
	if actl != ref {
		t.Fatalf("值 [%v], 转换后 [%v]; 与参考值 [%v] 不符合！", value, actl, ref)
	}

	// case
	value = 1_387_001
	ref = "壹佰叁拾捌万柒仟零壹"
	actl = value.ToChnRoundUpper()
	if actl != ref {
		t.Fatalf("值 [%v], 转换后 [%v]; 与参考值 [%v] 不符合！", value, actl, ref)
	}
}
