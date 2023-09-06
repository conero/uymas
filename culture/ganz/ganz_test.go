package ganz

import (
	"fmt"
	"testing"
)

func TestZodiacList(t *testing.T) {
	rsl := ZodiacList()
	ref := []string{"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"}
	if fmt.Sprintf("%#v", rsl) != fmt.Sprintf("%#v", ref) {
		t.Errorf("Zodiac 分解错误，%#v\n  ≠ %#v", rsl, ref)
	}
}

func TestPnPart(t *testing.T) {
	rsl := PnPart(DiZhi)
	ref := []string{"子丑", "寅卯", "辰巳", "午未", "申酉", "戌亥"}
	if fmt.Sprintf("%#v", rsl) != fmt.Sprintf("%#v", ref) {
		t.Errorf("Zodiac 分解错误，%#v\n  ≠ %#v", rsl, ref)
	}
}

func TestGzList(t *testing.T) {
	list := GzList()
	//t.Logf("%d => %#v\n", len(list), list)
	if len(list) != 60 {
		t.Error("Gzlist 配对数量有误，应为60！")
		return
	}

	if list[11] != "乙亥" {
		t.Error("Gzlist 配对有误，12 ≠ 乙亥")
	}

	if list[22] != "丙戌" {
		t.Error("Gzlist 配对有误，23 ≠ 丙戌")
	}

	if list[33] != "丁酉" {
		t.Error("Gzlist 配对有误，34 ≠ 丁酉")
	}

	if list[44] != "戊申" {
		t.Error("Gzlist 配对有误，45 ≠ 戊申")
	}

	if list[55] != "己未" {
		t.Error("Gzlist 配对有误，56 ≠ 己未")
	}

	if list[46] != "庚戌" {
		t.Error("Gzlist 配对有误，47 ≠ 庚戌")
	}

	if list[37] != "辛丑" {
		t.Error("Gzlist 配对有误，38 ≠ 辛丑")
	}

	if list[28] != "壬辰" {
		t.Error("Gzlist 配对有误，29 ≠ 壬辰")
	}

	if list[19] != "癸未" {
		t.Error("Gzlist 配对有误，20 ≠ 癸未")
	}
}

func TestCountGzAndZodiac(t *testing.T) {
	var rslGz, rslSx, refGz, refSx string
	var ipt int

	// case
	ipt = 1991
	refGz, refSx = "辛未", "羊"
	rslGz, rslSx = CountGzAndZodiac(ipt)
	if rslGz != refGz || rslSx != refSx {
		t.Errorf("input -> rsl ≠ ref: %d -> (%s, %s) ≠ (%s, %s)", ipt, rslGz, rslSx, refGz, refSx)
	}

	// case
	ipt = 1982
	refGz, refSx = "壬戌", "狗"
	rslGz, rslSx = CountGzAndZodiac(ipt)
	if rslGz != refGz || rslSx != refSx {
		t.Errorf("input -> rsl ≠ ref: %d -> (%s, %s) ≠ (%s, %s)", ipt, rslGz, rslSx, refGz, refSx)
	}

	// case
	ipt = 1977
	refGz, refSx = "丁巳", "蛇"
	rslGz, rslSx = CountGzAndZodiac(ipt)
	if rslGz != refGz || rslSx != refSx {
		t.Errorf("input -> rsl ≠ ref: %d -> (%s, %s) ≠ (%s, %s)", ipt, rslGz, rslSx, refGz, refSx)
	}

	// case
	ipt = 1964
	refGz, refSx = "甲辰", "龙"
	rslGz, rslSx = CountGzAndZodiac(ipt)
	if rslGz != refGz || rslSx != refSx {
		t.Errorf("input -> rsl ≠ ref: %d -> (%s, %s) ≠ (%s, %s)", ipt, rslGz, rslSx, refGz, refSx)
	}

	// case
	ipt = 1894
	refGz, refSx = "甲午", "马"
	rslGz, rslSx = CountGzAndZodiac(ipt)
	if rslGz != refGz || rslSx != refSx {
		t.Errorf("input -> rsl ≠ ref: %d -> (%s, %s) ≠ (%s, %s)", ipt, rslGz, rslSx, refGz, refSx)
	}

	// case
	ipt = 2024
	refGz, refSx = "甲辰", "龙"
	rslGz, rslSx = CountGzAndZodiac(ipt)
	if rslGz != refGz || rslSx != refSx {
		t.Errorf("input -> rsl ≠ ref: %d -> (%s, %s) ≠ (%s, %s)", ipt, rslGz, rslSx, refGz, refSx)
	}

	// case
	ipt = 2052
	refGz, refSx = "壬申", "猴"
	rslGz, rslSx = CountGzAndZodiac(ipt)
	if rslGz != refGz || rslSx != refSx {
		t.Errorf("input -> rsl ≠ ref: %d -> (%s, %s) ≠ (%s, %s)", ipt, rslGz, rslSx, refGz, refSx)
	}
}
