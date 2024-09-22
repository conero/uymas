// Package tm assistant methods related to time reporting
package tm

import (
	"errors"
	"gitee.com/conero/uymas/v2/data/input"
	"regexp"
	"strings"
	"time"
)

// SpendFn Get the program spend time for any format.
func SpendFn() func() time.Duration {
	now := time.Now()
	return func() time.Duration {
		return time.Since(now)
	}
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
			value := input.Stringer(s).Float()
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
			vDay := input.Stringer(s).Float()
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
			vHour := input.Stringer(s).Float()
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
			vMinute := input.Stringer(s).Float()
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
			vSecond := input.Stringer(s).Float()
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
