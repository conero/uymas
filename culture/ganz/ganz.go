// Package ganz Chinese Traditional Heavenly Stems and Earthly Branches(天干地支).
// Refer to <GB/T 33661-2017> http://c.gb688.cn/bzgk/gb/showGb?type=online&hcno=E107EA4DE9725EDF819F33C60A44B296
package ganz

import (
	"fmt"
	"strings"
	"time"
)

const (
	// TianGan Heavenly Stems（H-S）
	// reference link: https://mp.weixin.qq.com/s?__biz=MzA4MDM1NzYyNQ==&mid=2650691138&idx=1&sn=ca580627fbd7d910d282763450fa065f&chksm=87af87d4b0d80ec2f8a2335259dac634390bd215f45ffd973542047dbc271f052c1d5554c74c&scene=27
	// 甲   1H, the first of the ten Heavenly Stems.
	// 乙   2H, the first of the ten Heavenly Stems.
	TianGan = "甲乙丙丁戊己庚辛壬癸"
	// DiZhi Earthly Branches （E-B）
	// 子   1E, the first of the twelve Earthly Branches.
	// 丑   2E, the 2nd of the twelve Earthly Branches.
	DiZhi      = "子丑寅卯辰巳午未申酉戌亥"
	DiZhiMonth = "冬月,腊月,正月,二月,三月,四月,五月,六月,七月,八月,九月,十月"
	DiZhiTime  = "夜半,鸡鸣,平旦,日出,食时,隅中,日中,日昳,哺时,日入,黄昏,人定"
	Zodiac     = "鼠牛虎兔龙蛇马羊猴鸡狗猪"
)

var (
	cacheTgList     []string
	cacheDzList     []string
	cacheZodiacList []string
	cacheDzTimeDick []DzTime
)

func TgList() []string {
	if cacheTgList == nil {
		cacheTgList = strings.Split(TianGan, "")
	}
	return cacheTgList
}

func DzList() []string {
	if cacheDzList == nil {
		cacheDzList = strings.Split(DiZhi, "")
	}
	return cacheDzList
}

func ZodiacList() []string {
	if cacheZodiacList == nil {
		cacheZodiacList = strings.Split(Zodiac, "")
	}
	return cacheZodiacList
}

// PnPart parse PN(Positive Negative)-parts by define string
func PnPart(def string) []string {
	list := strings.Split(def, "")
	var pnPart []string

	vLen := len(list)
	for i := 0; i < vLen; i++ {
		j := i * 2
		if j >= vLen {
			break
		}
		pnPart = append(pnPart, fmt.Sprintf("%s%s", list[j], list[j+1]))
	}

	return pnPart
}

// GzList List of stems and branches obtained through pairing with heavenly stems and earthly branches
func GzList() []string {
	var list []string
	tgPars := PnPart(TianGan)
	dzPars := PnPart(DiZhi)

	dzIdx := 0
	dzLen := len(dzPars)
	for x := 0; x <= 5; x++ {
		for _, tg := range tgPars {
			pnPars := strings.Split(tg, "")
			dzPnPars := strings.Split(dzPars[dzIdx], "")

			// 同极相配，异极互斥
			// Positive
			pars := fmt.Sprintf("%s%s", pnPars[0], dzPnPars[0])
			list = append(list, pars)

			// Negative
			pars = fmt.Sprintf("%s%s", pnPars[1], dzPnPars[1])
			list = append(list, pars)

			dzIdx += 1
			if dzIdx == dzLen {
				dzIdx = 0
			}
		}
	}

	return list
}

// CountGzAndZodiac Calculate the Heavenly Stems, Earthly Branches, and Zodiac Phases Based on the Year
// reference link: http://www.360doc.com/content/19/0131/15/30390538_812371769.shtml
func CountGzAndZodiac(year int) (gz string, zodiac string) {
	tgLs := TgList()
	dzLs := DzList()

	tg := (year - 3) % 10
	tg = tg - 1
	if tg < 0 {
		tg = len(tgLs) - 1
	}
	dz := (year - 3) % 12
	dz = dz - 1
	if dz < 0 {
		dz = len(dzLs) - 1
	}

	zodiacLs := ZodiacList()
	gz = fmt.Sprintf("%s%s", tgLs[tg], dzLs[dz])
	zodiac = zodiacLs[dz]
	return
}

type DzTime struct {
	Name      string // dizhi name like '子'
	Alias     string
	Range     [2]int
	Month     int
	MonthName string
	Zodiac    string
}

func (c *DzTime) String() string {
	return fmt.Sprintf("%s时, [%d:00-%d:00]", c.Name, c.Range[0], c.Range[1])
}

// DzTimeList reference link: https://mbd.baidu.com/newspage/data/dtlandingsuper?nid=dt_3438222656616746597
// 十二时辰
// https://zhuanlan.zhihu.com/p/490496771
func DzTimeList() []DzTime {
	if cacheDzTimeDick != nil {
		return cacheDzTimeDick
	}

	dztList := strings.Split(DiZhiTime, ",")
	monthList := strings.Split(DiZhiMonth, ",")
	zodiacLs := ZodiacList()
	for i, dz := range DzList() {
		name := dz

		begin := 23 + i*2
		end := begin + 2
		vRang := [2]int{begin % 24, end % 24}
		cacheDzTimeDick = append(cacheDzTimeDick, DzTime{
			Name:      name,
			Alias:     dztList[i],
			Range:     vRang,
			Month:     (11 + i) % 12,
			MonthName: monthList[i],
			Zodiac:    zodiacLs[i],
		})
	}

	return cacheDzTimeDick
}

// TimeParse the parsing date is in stem format
func TimeParse(tm time.Time) string {
	year := tm.Year()
	gz, zodiac := CountGzAndZodiac(year)

	dzTimeLs := DzTimeList()

	month := tm.Month()
	mth := dzTimeLs[month]

	return fmt.Sprintf("农历%s年（%s年）%s(%s月)%s时\n", gz, zodiac, mth.MonthName, mth.Zodiac, mth.Name)
}
