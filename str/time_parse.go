package str

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	strAsFloatReg *regexp.Regexp
)

func TimeParse(tmStr string) (time.Time, error) {
	layout, err := TimeParseLayout(tmStr)
	if err != nil {
		return time.Time{}, err
	}
	return time.Parse(layout, tmStr)
}

type TimeLayoutDetector struct {
	input  string
	layout string
}

// HH:ii:ss
func (c *TimeLayoutDetector) timeDetect(tm string) string {
	var layout string
	if tm == "" {
		return layout
	}
	tmSpl := ":"
	// time
	if strings.Contains(tm, tmSpl) {
		for j, ss := range strings.Split(tm, tmSpl) {
			ssLn := len(ss)
			if j == 0 {
				if ssLn == 2 {
					layout += "15"
				} else if ssLn == 1 {
					layout += "3"
				}
			} else if j == 1 {
				if ssLn == 2 {
					layout += tmSpl + "04"
				} else if ssLn == 1 {
					layout += tmSpl + "4"
				}
			} else if j == 2 {
				if ssLn == 2 {
					layout += tmSpl + "05"
				} else if ssLn == 1 {
					layout += tmSpl + "5"
				}
			}
		}
	}

	return layout
}

// yyyy-mm-dd HH:ii:ss
// yyyy/mm/dd HH:ii:ss
func (c *TimeLayoutDetector) layoutFmt1(spl string) string {
	tmStr := c.input
	var layout string
	var layoutQueue []string
	// yyyy-mm-dd
	if !strings.Contains(tmStr, spl) {
		return layout
	}

	for idx, s := range strings.Split(tmStr, spl) {
		sLen := len(s)
		if idx == 0 { // yyyy|yy
			if sLen == 4 {
				layoutQueue = append(layoutQueue, "2006")
			} else {
				layoutQueue = append(layoutQueue, "06")
			}
		} else if idx == 1 {
			if sLen == 1 {
				layoutQueue = append(layoutQueue, "1")
			} else {
				layoutQueue = append(layoutQueue, "01")
			}
		} else if idx == 2 { // Is there a time format present
			if sLen == 1 {
				layoutQueue = append(layoutQueue, "2")
			} else {
				spanSpl := " "
				var dayStr = s
				var tStr = s
				var part3 string
				cIdx := strings.Index(s, spanSpl)
				if cIdx > -1 {
					dayStr = s[:cIdx]
					tStr = s[cIdx+1:]
				}

				dayStrLen := len(dayStr)
				if dayStrLen == 2 {
					part3 += "02"
				} else if dayStrLen == 1 {
					part3 += "2"
				}

				tmLayout := c.timeDetect(tStr)
				if tmLayout != "" {
					if cIdx > -1 {
						tmLayout = spanSpl + tmLayout
					}
					part3 += tmLayout
				}
				layoutQueue = append(layoutQueue, part3)
			}
		}
	}

	return strings.Join(layoutQueue, spl)
}

// yyyymmdd
func (c *TimeLayoutDetector) layoutFmt2() string {
	input := c.input
	var layout string
	isMatched, _ := regexp.MatchString(`^(\d{2}|\d{4})\d*(\s?\d*)?$`, input)
	if !isMatched {
		return ""
	}

	spaceSpl := " "
	idx := strings.Index(input, spaceSpl)
	var dtStr, tmStr string = input, ""
	if idx > -1 {
		dtStr = input[:idx]
		tmStr = input[idx+1:]
	}

	// date detected
	rIdx := 0
	for {
		dtLn := len(dtStr)
		if rIdx == 0 {
			if dtLn == 8 {
				layout = "20060102"
				break
			} else if dtLn == 6 && tmStr == "" {
				layout = "060102"
				break
			} else if dtLn >= 4 {
				layout += "2006"
				dtStr = dtStr[4:]
			}
		} else if rIdx == 1 {
			if dtLn >= 2 {
				layout += "01"
				dtStr = dtStr[2:]
			}
		} else if rIdx == 2 {
			if dtLn >= 2 {
				layout += "02"
				dtStr = dtStr[:2]
			}
		} else {
			break
		}

		rIdx += 1
	}

	// time detected
	if tmStr != "" {
		if idx > -1 {
			layout += " "
		}
		rIdx = 0
		for {
			tmLn := len(tmStr)
			if rIdx == 0 {
				if tmLn == 6 {
					layout += "150405"
					break
				} else if tmLn >= 2 {
					layout += "15"
					tmStr = tmStr[2:]
				}
			} else if rIdx == 1 {
				if tmLn >= 2 {
					layout += "04"
					tmStr = tmStr[2:]
				}
			} else if rIdx == 2 {
				if tmLn >= 2 {
					layout += "05"
					tmStr = tmStr[2:]
				}
			} else {
				break
			}
			rIdx += 1
		}
	}

	return layout
}

// yyyy年mm月dd日
func (c *TimeLayoutDetector) layoutFmt3() string {
	input := c.input
	var layout string
	isMatched, _ := regexp.MatchString(`^(\d{2}|\d{4})年(\d{1,2}月)?(\d{1,2}日?)?$`, input)
	if !isMatched {
		return ""
	}

	// year
	splYear := "年"
	queue := strings.Split(input, splYear)
	yearVal := queue[0]
	vLn := len(yearVal)
	if vLn == 4 {
		layout += "2006年"
	} else if vLn == 2 {
		layout += "06年"
	}
	input = queue[1]

	// month
	splMon := "月"
	queue = strings.Split(input, splMon)
	val := queue[0]
	ln := len(val)
	if ln == 2 {
		layout += "01月"
	} else if vLn == 2 {
		layout += "1月"
	}
	if len(queue) > 1 {
		input = queue[1]
	} else {
		input = ""
	}

	// month
	splMon = "日"
	queue = strings.Split(input, splMon)
	val = queue[0]
	ln = len(val)
	if ln == 2 {
		layout += "02"
	} else if vLn == 2 {
		layout += "2"
	}
	if strings.Contains(input, splMon) {
		layout += "日"
	}
	return layout
}

func (c *TimeLayoutDetector) Parse() (string, error) {
	c.layout = ""
	if c.input == "" {
		return "", errors.New("输入的日期格式为空，解析失败！")
	}

	var lot string
	if lot = c.layoutFmt1("-"); lot != "" {
		c.layout = lot
		return lot, nil
	}

	if lot = c.layoutFmt1("/"); lot != "" {
		c.layout = lot
		return lot, nil
	}

	if lot = c.layoutFmt2(); lot != "" {
		c.layout = lot
		return lot, nil
	}

	if lot = c.layoutFmt3(); lot != "" {
		c.layout = lot
		return lot, nil
	}

	return "", errors.New("日期格式失败！")
}

func (c *TimeLayoutDetector) Layout() string {
	return c.layout
}

func NewTimeLotDtr(tmStr string) *TimeLayoutDetector {
	return &TimeLayoutDetector{
		input: tmStr,
	}
}

// TimeParseLayout Resolve the given time format into (convert to) standard format.
// The supported formats are as follows(time can be missing):
// - 2006-01-02 15:04:06
// - 2006-01-02
// - 2006-1-2
// - 2006/01/02 15:04:06
// - 2006年01月02日 15:04:06
// - 20060102 15:04:06
func TimeParseLayout(tmStr string) (string, error) {
	tld := NewTimeLotDtr(tmStr)
	return tld.Parse()
}

// ParseDuration parse duration by input string data like day/d(default)/天,hour/h/时,minute/m/分,second/s/秒
//
// like: -10d73s, 73h11m1002s, 13分12秒。"-"/"减"/“负” are similar terms
func ParseDuration(dura string) (time.Duration, error) {
	if dura == "" {
		return 0, nil
	}
	symbol := 1
	queue := strings.Split(dura, "")
	first := queue[0]
	if first == "-" || first == "减" || first == "负" {
		dura = strings.Join(queue[1:], "")
		symbol = -1
	}

	if dura == "" {
		return 0, errors.New("duration data is error that no number")
	}

	var duration time.Duration

	timeFormatReg := regexp.MustCompile(`\d+(\.\d+)*:\d+(\.\d+)*:\d+(\.\d+)*`)
	if timeFormatReg.MatchString(dura) {
		var durationTmp time.Duration
		for i, s := range strings.Split(dura, ":") {
			value := strAsFloat(s)
			if value == 0 {
				continue
			}
			durationTmp = time.Duration(value * 1000)
			switch i {
			case 0: // hour
				duration += durationTmp * time.Millisecond * 60 * 60
			case 1: // minute
				duration += durationTmp * time.Millisecond * 60
			case 2: // second
				duration += durationTmp * time.Millisecond
			}
		}
		dura = ""
	}

	// day
	reg := regexp.MustCompile(`\d+(\.\d+)*\s?(天|d)`)
	if reg.MatchString(dura) {
		for _, s := range reg.FindAllString(dura, -1) {
			vDay := strAsFloat(s)
			if vDay == 0 {
				continue
			}
			duration += time.Millisecond * time.Duration(vDay*24*60*60*1000)
			dura = strings.ReplaceAll(dura, s, "")
		}
	}

	// hour
	reg = regexp.MustCompile(`\d+(\.\d+)*\s?(时|h)`)
	if reg.MatchString(dura) {
		for _, s := range reg.FindAllString(dura, -1) {
			vHour := strAsFloat(s)
			if vHour == 0 {
				continue
			}
			duration += time.Millisecond * time.Duration(vHour*60*60*1000)
			dura = strings.ReplaceAll(dura, s, "")
		}
	}

	// minute
	reg = regexp.MustCompile(`\d+(\.\d+)*\s?(分|m)`)
	if reg.MatchString(dura) {
		for _, s := range reg.FindAllString(dura, -1) {
			vMinute := strAsFloat(s)
			if vMinute == 0 {
				continue
			}
			duration += time.Millisecond * time.Duration(vMinute*60*1000)
			dura = strings.ReplaceAll(dura, s, "")
		}
	}

	// second
	reg = regexp.MustCompile(`\d+(\.\d+)*\s?(秒|s)`)
	if reg.MatchString(dura) {
		for _, s := range reg.FindAllString(dura, -1) {
			vSecond := strAsFloat(s)
			if vSecond == 0 {
				continue
			}
			duration += time.Millisecond * time.Duration(vSecond*1000)
			dura = strings.ReplaceAll(dura, s, "")
		}
	}
	duration *= time.Duration(symbol)

	return duration, nil

}

func strAsFloat(s string) float64 {
	if strAsFloatReg == nil {
		strAsFloatReg = regexp.MustCompile(`-?\d+(\.\d+)*`)
	}

	for _, eS := range strAsFloatReg.FindAllString(s, 1) {
		vf, _ := strconv.ParseFloat(eS, 64)
		return vf
	}

	return 0
}
