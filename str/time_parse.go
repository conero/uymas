package str

import (
	"errors"
	"strings"
	"time"
)

func TimeParse(tmStr string) (time.Time, error) {
	layout, err := TimeParseLayout(tmStr)
	if err != nil {
		return time.Time{}, err
	}
	return time.Parse(layout, tmStr)
}

type TimeLayoutDetector struct {
	input   string
	dateLyt string
	timeLyt string
	layout  string
}

// yyyy-yy-mm HH:ii:ss
func (c *TimeLayoutDetector) layoutD1() string {
	tmStr := c.input
	var layout string
	var layoutQueue []string
	spl := "-"
	// yyyy-yy-mm
	if strings.Index(tmStr, spl) > -1 {
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

					tmSpl := ":"
					// time
					if strings.Index(tStr, tmSpl) > -1 {
						if cIdx > -1 {
							part3 += spanSpl
						}
						for j, ss := range strings.Split(tStr, tmSpl) {
							ssLn := len(ss)
							if j == 0 {
								if ssLn == 2 {
									part3 += "15"
								} else if ssLn == 1 {
									part3 += "3"
								}
							} else if j == 1 {
								if ssLn == 2 {
									part3 += tmSpl + "04"
								} else if ssLn == 1 {
									part3 += tmSpl + "4"
								}
							} else if j == 2 {
								if ssLn == 2 {
									part3 += tmSpl + "05"
								} else if ssLn == 1 {
									part3 += tmSpl + "5"
								}
							}
						}
					}
					layoutQueue = append(layoutQueue, part3)
				}
			}
		}

		return strings.Join(layoutQueue, spl)
	}

	return layout
}

func (c *TimeLayoutDetector) Parse() (string, error) {
	if c.input == "" {
		return "", errors.New("输入的日期格式为空，解析失败！")
	}

	if lot := c.layoutD1(); lot != "" {
		return lot, nil
	}

	return "", errors.New("日期格式失败！")
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
