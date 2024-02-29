package digit

var zhIndexDick = map[int]string{
	0:    "零",
	1:    "一",
	2:    "二",
	3:    "三",
	4:    "四",
	5:    "五",
	6:    "六",
	7:    "七",
	8:    "八",
	9:    "九",
	10:   "十",
	100:  "百",
	1000: "千",
}

// LowerIndex convert serial number to Chinese digits
// @todo need to do
func LowerIndex(i int) string {
	var zh string
	if i <= 10 {
		return zhIndexDick[i]
	}
	return zh
}
