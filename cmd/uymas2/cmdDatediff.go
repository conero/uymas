package main

import (
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/cli/ansi"
	"gitee.com/conero/uymas/v2/cli/gen"
	"gitee.com/conero/uymas/v2/logger/lgr"
	"gitee.com/conero/uymas/v2/str"
	"gitee.com/conero/uymas/v2/util/tm"
	"strings"
	"time"
)

/*
*
--end,-e    结束数据，默认为当天

	--add,-a    日期加，采用 d/h/m/s 表示天/时/分秒，减则采用负号（"-/减/负"）。默认为天
*/
type ddOption struct {
	End  string `cmd:"end,e help:结束数据，默认为当天"`
	Add  string `cmd:"add,a help:日期加，采用\\sd/h/m/s\\s表示天/时/分秒，减则采用负号（“-/减/负”）。默认为天"`
	Date string `cmd:"date isdata"`
}

func cmdDatediff(args cli.ArgsParser) {
	spendFn := tm.SpendFn()
	var opt ddOption
	err := gen.ArgsDress(args, &opt)
	if err != nil {
		lgr.Error(err.Error())
		return
	}

	date := opt.Date
	if date == "" {
		lgr.Error("请输入日期！")
		return
	}
	endDate := opt.End

	// 日期解析
	var vTm time.Time
	if strings.ToLower(date) == "now" {
		vTm = time.Now()
	} else {
		vTm, err = str.TimeParse(date)
		if err != nil {
			lgr.Error("日期格式不支持！")
			return
		}
	}

	// 日期运算
	add := opt.Add
	if add != "" {
		dura, err := tm.ParseDuration(add)
		if err != nil {
			lgr.Error("加入日期错误，%s", err)
			return
		}

		newTm := vTm.Add(dura)
		lgr.Info("当前时间：%s, 运算：%s，得\n    %s",
			ansi.Style(vTm.Format(time.DateTime), ansi.BlackBr),
			ansi.Style(add+"("+dura.String()+")", ansi.BlackBr),
			ansi.Style(newTm.Format(time.DateTime)), ansi.GreenBr)
		return
	}

	now := time.Now()
	if endDate != "" {
		tmEnd, err := str.TimeParse(endDate)
		if err != nil {
			lgr.Error("日期格式不支持！")
			return
		}
		now = tmEnd
	}
	diff := vTm.Sub(now)
	d3 := NewD3(diff)

	var diffType string
	if diff > 0 {
		diffType = "之后"
	} else {
		diffType = "之前"
	}

	// 输出
	cmdLsString := d3.cmdListing()
	if cmdLsString != "" {
		cmdLsString = " " + cmdLsString + "\n"
	}
	lgr.Info("%s 距今(%s)比较：\n 时间总差：%v\n 差别类型：%s\n%s\n用时：%v",
		date, now.Format("2006-01-02"), diff, diffType, cmdLsString, spendFn())
}
