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
	//value = 99_999_999_999_999_999 // 溢出
	value = 8_909_970_069_905_049
	t.Logf("cover: %f => %v", value, value.ToChnRoundUpper())

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
	value = 100
	ref = "壹佰"
	actl = value.ToChnRoundUpper()
	if actl != ref {
		t.Fatalf("值 [%v], 转换后 [%v]; 与参考值 [%v] 不符合！", value, actl, ref)
	}

	value = 101
	ref = "壹佰零壹"
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

// 参考工具： http://daxie.gjcha.com/
func TestCover_ToChnRoundLower(t *testing.T) {
	var value Cover
	var ref, actl string

	// case
	value = 8_909_970_069_905_049
	t.Logf("cover: %v => %v", value, value.ToChnRoundLower())

	// case
	value = 9
	ref = "九"
	actl = value.ToChnRoundLower()
	if actl != ref {
		t.Fatalf("值 [%v], 转换后 [%v]; 与参考值 [%v] 不符合！", value, actl, ref)
	}

	// case
	value = 98
	ref = "九十八"
	actl = value.ToChnRoundLower()
	if actl != ref {
		t.Fatalf("值 [%v], 转换后 [%v]; 与参考值 [%v] 不符合！", value, actl, ref)
	}

	// case
	value = 185
	ref = "一百八十五"
	actl = value.ToChnRoundLower()
	if actl != ref {
		t.Fatalf("值 [%v], 转换后 [%v]; 与参考值 [%v] 不符合！", value, actl, ref)
	}

	// 9001
	value = 9001
	ref = "九千〇一"
	actl = value.ToChnRoundLower()
	if actl != ref {
		t.Fatalf("值 [%v], 转换后 [%v]; 与参考值 [%v] 不符合！", value, actl, ref)
	}

	// case
	value = 1_387_001
	ref = "一百三十八万七千〇一"
	actl = value.ToChnRoundLower()
	if actl != ref {
		t.Fatalf("值 [%v], 转换后 [%v]; 与参考值 [%v] 不符合！", value, actl, ref)
	}
}

func TestNumberCoverRmb(t *testing.T) {
	var ref, atl string
	var value float64

	value = 1903.089
	ref = "壹仟玖佰零叁元捌分"
	atl = NumberCoverRmb(value)
	if atl != ref {
		t.Fatalf("TestNumberCoverRmb -> 值 [%v], 转换后 [%v]; 与参考值 [%v] 不符合！", value, atl, ref)
	}

	value = 13.01
	ref = "壹拾叁元壹分"
	atl = NumberCoverRmb(value)
	if atl != ref {
		t.Fatalf("TestNumberCoverRmb -> 值 [%v], 转换后 [%v]; 与参考值 [%v] 不符合！", value, atl, ref)
	}

	value = 13.02
	ref = "壹拾叁元贰分"
	atl = NumberCoverRmb(value)
	if atl != ref {
		t.Fatalf("TestNumberCoverRmb -> 值 [%v], 转换后 [%v]; 与参考值 [%v] 不符合！", value, atl, ref)
	}

	value = 1629.49
	ref = "壹仟陆佰贰拾玖元肆角玖分"
	atl = NumberCoverRmb(value)
	if atl != ref {
		t.Fatalf("TestNumberCoverRmb -> 值 [%v], 转换后 [%v]; 与参考值 [%v] 不符合！", value, atl, ref)
	}

	value = 18.3
	ref = "壹拾捌元叁角"
	atl = NumberCoverRmb(value)
	if atl != ref {
		t.Fatalf("TestNumberCoverRmb -> 值 [%v], 转换后 [%v]; 与参考值 [%v] 不符合！", value, atl, ref)
	}
}
