package digit

import "strings"

var zhIndexDick = map[int]string{
	0:          "零",
	1:          "一",
	2:          "二",
	3:          "三",
	4:          "四",
	5:          "五",
	6:          "六",
	7:          "七",
	8:          "八",
	9:          "九",
	UnitSValue: "十",
	UnitBValue: "百",
	UnitQValue: "千",
	UnitWValue: "万",
	UnitYValue: "亿",
}

// LowerIndex convert serial number to Chinese digits
func LowerIndex(i int) string {
	if i <= 10 {
		return zhIndexDick[i]
	}

	zh := NumberCoverChnDigit(float64(i), false)
	// replace `〇` -> `零`
	zh = strings.ReplaceAll(zh, "〇", "零")
	// simple `一十` as `十`
	if strings.Index(zh, "一十") == 0 {
		zh = strings.Replace(zh, "一十", "十", 1)
	}

	return zh
}
