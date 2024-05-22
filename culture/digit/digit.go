// Package digit Provide digitization tools such as converting numbers to Text uppercase and commonly used strings.
// link: http://daxie.gjcha.com/
package digit

const (
	UnitSValue = 10
	UnitBValue = 100
	UnitQValue = 1_000
	UnitWValue = 10_000
	UnitYValue = 100_000_000
)

// chinese Uppercase Digit Dictionary
var vUpperMap = map[uint32]string{
	0:          "零",
	1:          "壹",
	2:          "贰",
	3:          "叁",
	4:          "肆",
	5:          "伍",
	6:          "陆",
	7:          "柒",
	8:          "捌",
	9:          "玖",
	UnitSValue: "拾",
	UnitBValue: "佰",
	UnitQValue: "仟",
	UnitWValue: "万",
	UnitYValue: "亿",
}

// chinese Lowercase Digit Dictionary
var vLowerMap = map[uint32]string{
	0:          "〇",
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
