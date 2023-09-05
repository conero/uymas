package main

import (
	"fmt"
	"gitee.com/conero/uymas/bin"
	"gitee.com/conero/uymas/logger/lgr"
	"gitee.com/conero/uymas/str"
	"regexp"
	"strconv"
	"strings"
)

type defaultApp struct {
	bin.CliApp
}

func (c *defaultApp) DefaultIndex() {
	fmt.Printf("demo 子命令")
}

func (c *defaultApp) DefaultHelp() {
	fmt.Println("cal [equal]  计算器")
}

func (c *defaultApp) Cal() {
	equal := c.Cc.SubCommand
	if equal != "" {
		spanReg := regexp.MustCompile(`\s+`)
		equal = spanReg.ReplaceAllString(equal, "")
	}
	if equal == "" {
		lgr.Error("请输入等式符号！")
		return
	}

	// 计算器（加减乘除）
	// 字符串转数字
	str2F64 := func(s string) float64 {
		f64, _ := strconv.ParseFloat(s, 10)
		return f64
	}

	//计算
	regBracketSign := regexp.MustCompile(`^\(.*\)$`)
	// add, subtract, multiply and divide => +-*/
	regMd := regexp.MustCompile(`(\d(\.\d)?)[*/](\d(\.\d)?)`)
	mdSig := regexp.MustCompile(`[*/]`)
	regAs := regexp.MustCompile(`(\d(\.\d)?)[+-](\d(\.\d)?)`)
	asSig := regexp.MustCompile(`[+-]`)
	asmdSig := regexp.MustCompile(`[+\-*/]+`)
	calFn := func(bkt string) string {
		bkt = strings.TrimSpace(bkt)
		// 去括号
		if regBracketSign.MatchString(bkt) && len(bkt) > 2 {
			bkt = bkt[1:]
			bkt = bkt[:len(bkt)-1]
		}

		// 乘除运算
		mdLs := regMd.FindAllString(bkt, -1)
		for _, md := range mdLs {
			mdSigLs := mdSig.FindAllString(md, -1)
			if len(mdSigLs) != 1 {
				bkt = strings.ReplaceAll(bkt, md, "0")
				continue
			}
			parts := mdSig.Split(bkt, -1)
			if len(parts) != 2 {
				bkt = strings.ReplaceAll(bkt, md, "0")
				continue
			}

			var mgCtt float64 = 0
			if mdSigLs[0] == "*" {
				mgCtt = str2F64(parts[0]) * str2F64(parts[1])
			} else {
				mgCtt = str2F64(parts[0]) / str2F64(parts[1])
			}

			bkt = strings.ReplaceAll(bkt, md, str.FloatSimple(fmt.Sprintf("%.5f", mgCtt)))

		}

		// 加减运算
		asLs := regAs.FindAllString(bkt, -1)
		for _, as := range asLs {
			asSigLs := asSig.FindAllString(as, -1)
			if len(asSigLs) != 1 {
				bkt = strings.ReplaceAll(bkt, as, "0")
				continue
			}
			parts := asSig.Split(bkt, -1)
			if len(parts) != 2 {
				bkt = strings.ReplaceAll(bkt, as, "0")
				continue
			}

			var mgCtt float64 = 0
			if asSigLs[0] == "+" {
				mgCtt = str2F64(parts[0]) + str2F64(parts[1])
			} else {
				mgCtt = str2F64(parts[0]) - str2F64(parts[1])
			}

			mgCttStr := fmt.Sprintf("%.5f", mgCtt)
			bkt = strings.ReplaceAll(bkt, as, str.FloatSimple(mgCttStr))
		}

		return bkt
	}

	// 循环
	regBracket := regexp.MustCompile(`([^()]+)`)
	regEqual := regexp.MustCompile(``)
	countBad := 0
	for {
		isEnd := false

		bracket := regBracket.FindAllString(equal, -1)
		for _, bkt := range bracket {
			newBkt := calFn(bkt)
			equal = strings.ReplaceAll(equal, bkt, newBkt)
		}

		equal = calFn(equal)

		if !asmdSig.MatchString(equal) {
			break
		}

		if equal == "" {
			break
		}

		if isEnd {
			break
		}

		if !regEqual.MatchString(equal) {
			break
		}

		// 大于5000直接退出
		if countBad > 5000 {
			break
		}
		countBad += 1
	}
	lgr.Info("输入等式：%s\n    => %s", c.Cc.SubCommand, equal)
}
