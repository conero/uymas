// Package digit Provide digitization tools such as converting numbers to Chinese uppercase and commonly used strings.
// link: http://daxie.gjcha.com/
package digit

// chinese Uppercase Digit Dictionary
var vUpperMap = map[int8]string{
	0:  "零",
	1:  "壹",
	2:  "贰",
	3:  "叁",
	4:  "肆",
	5:  "伍",
	6:  "陆",
	7:  "柒",
	8:  "捌",
	9:  "玖",
	10: "拾",
}

// UnitUpperS Equivalent data unit 10
const UnitUpperS = "拾"

// UnitUpperB Equivalent data unit 100
const UnitUpperB = "佰"
const UnitSValue = 10
const UnitBValue = 100

// UnitUpperQ Equivalent data unit 1,000
const UnitUpperQ = "仟"
const UnitQValue = 1_000

// UnitUpperW Equivalent data unit 10,000
const UnitUpperW = "万"
const UnitWValue = 10_000

// UnitUpperY Equivalent data unit 10,000,000
const UnitUpperY = "亿"
const UnitYValue = 100_000_000

// UnitLowerS Equivalent data unit 10
const UnitLowerS = "十"

// UnitLowerB Equivalent data unit 100
const UnitLowerB = "百"

// UnitLowerQ Equivalent data unit 1,000
const UnitLowerQ = "千"

// UnitLowerW Equivalent data unit 10,000
const UnitLowerW = "万"

// UnitLowerY Equivalent data unit 10,000,000
const UnitLowerY = "亿"

// chinese Lowercase Digit Dictionary
var vLowerMap = map[int8]string{
	0:  "〇",
	1:  "一",
	2:  "二",
	3:  "三",
	4:  "四",
	5:  "五",
	6:  "六",
	7:  "七",
	8:  "八",
	9:  "九",
	10: "十",
}
