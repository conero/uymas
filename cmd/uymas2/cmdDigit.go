package main

import (
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/cli/gen"
	"gitee.com/conero/uymas/v2/culture/digit"
	"gitee.com/conero/uymas/v2/logger/lgr"
	"regexp"
	"strconv"
)

type digitOption struct {
	Rmb   bool `cmd:"rmb,r help:输出人民币格式"`
	Lower bool `cmd:"lower,l help:转为小写字母"`
	Both  bool `cmd:"both,b help:同时转为大小写"`
}

func cmdDigit(args cli.ArgsParser) {
	var opt digitOption
	err := gen.ArgsDress(args, &opt)
	if err != nil {
		lgr.Error(err.Error())
		return
	}
	value := args.SubCommand()
	if value == "" {
		lgr.Error("请指定阿拉伯数字或中文数字！")
		return
	}

	isMatch, _ := regexp.MatchString(`\d+(.?\d+)?`, value)
	if isMatch {
		lgr.Info("识别为：阿拉伯数字转中文数字")
		vNum, err := strconv.ParseFloat(value, 64)
		if err != nil {
			lgr.Error("%s 不是有效数字!", value)
			return
		}
		isRmb := opt.Rmb
		var cv = digit.Cover(vNum)
		if opt.Lower {
			var valueStr string
			if isRmb {
				valueStr = cv.ToRmbLower()
			} else {
				valueStr = cv.ToChnRoundLower()
			}
			lgr.Info("转化中文小写数字成功！\n\n %v", valueStr)
		}
		if opt.Both {
			if isRmb {
				lgr.Info("转化中文大小写数字成功！\n\n %v\n %v\n %v\n %v",
					cv.ToChnRoundUpper(), cv.ToChnRoundLower(), cv.ToRmbUpper(), cv.ToRmbLower())
				return
			}
			lgr.Info("转化中文大小写数字成功！\n\n %v\n %v", cv.ToChnRoundUpper(), cv.ToChnRoundLower())
			return
		}

		var valueStr string
		if isRmb {
			valueStr = cv.ToRmbUpper()
		} else {
			valueStr = cv.ToChnRoundUpper()
		}
		lgr.Info("转化中文大写数成功！\n\n %v", valueStr)
		return
	}
}
