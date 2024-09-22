package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

type DateDiffDesc struct {
	diff     time.Duration
	allDay   float64 // 所有天数，将小时转化为她
	allWeek  float64 // 总周
	allMonth float64 // 总月
	allYear  float64 // 总年
	isPast   bool    // 是否为已过去的天数（即负数）
}

// 运算
func (c *DateDiffDesc) calculate() {
	c.allDay = float64(c.diff) / (24 * float64(time.Hour))
	if c.allDay > 0 {
		c.allWeek = c.allDay / 7     // 一周7天
		c.allMonth = c.allDay / 30.4 // `365/12 = 30.4`
		c.allYear = c.allDay / 365   // 一年365天
	}
}

// 命令类别形式函数
func (c *DateDiffDesc) cmdListing() string {
	var queue []string
	// 年计算
	if c.allYear >= 1 {
		year, yearLst := math.Modf(c.allYear)
		lstStr := ""
		for {
			day, _ := math.Modf(yearLst * 365)
			if day > 30.4 {
				mth, mthLst := math.Modf(day / 30.4)
				lstStr += fmt.Sprintf("%d个月", int(mth))
				day = mthLst * 30.4
			}
			if day > 7 {
				wk, wkLst := math.Modf(day / 7)
				lstStr += fmt.Sprintf("%d周", int(wk))
				day = wkLst * 7
			}
			if day > 0 {
				lstStr += fmt.Sprintf("%d天", int(day))
			}
			break
		}
		queue = append(queue, fmt.Sprintf("按年计算: %d年%s", int(year), lstStr))
	}
	// 月计算
	if c.allMonth >= 1 {
		mth, mthLst := math.Modf(c.allMonth)
		lstStr := ""
		for {
			day, _ := math.Modf(mthLst * 30.4)
			if day > 7 {
				wk, wkLst := math.Modf(day / 7)
				lstStr += fmt.Sprintf("%d周", int(wk))
				day = wkLst * 7
			}
			if day > 0 {
				lstStr += fmt.Sprintf("%d天", int(day))
			}
			break
		}
		queue = append(queue, fmt.Sprintf("按月计算: %d个月%s", int(mth), lstStr))
	}
	// 周计算
	if c.allWeek >= 1 {
		wk, wkLst := math.Modf(c.allWeek)
		lstStr := ""
		for {
			day, _ := math.Modf(wkLst * 7)
			if day > 0 {
				lstStr += fmt.Sprintf("%d天", int(day))
			}
			break
		}
		queue = append(queue, fmt.Sprintf("按周计算: %d周%s", int(wk), lstStr))
	}
	if c.allDay >= 1 {
		queue = append(queue, fmt.Sprintf("按天计算: %d天", int(c.allDay)))
	}
	return strings.Join(queue, "\n ")
}

func NewD3(diff time.Duration) *DateDiffDesc {
	var isPast = false
	if diff < 0 {
		isPast = true
		diff = time.Duration(math.Abs(float64(diff)))
	}
	d3 := &DateDiffDesc{
		diff:   diff,
		isPast: isPast,
	}
	d3.calculate()
	return d3
}
