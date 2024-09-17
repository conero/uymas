package main

import (
	"fmt"
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/logger/lgr"
	"gitee.com/conero/uymas/v2/str"
	"gitee.com/conero/uymas/v2/util/tm"
	"math/rand"
	"strings"
)

func cmdTest(arg cli.ArgsParser) {
	spendFn := tm.SpendFn()
	if arg.Switch("verbose", "V") {
		fmt.Println()
		fmt.Println("参数解析，数据如下")
		fmt.Println()
		fmt.Printf("value: %v\n", arg.Values())
		fmt.Printf("option: %v\n", arg.Option())
		fmt.Printf("CommandList: %v\n", arg.CommandList())
	}
	option := arg.List("option", "O")
	if len(option) > 0 {
		fmt.Printf("Read option: %v\n", arg.Get(option...))
	}

	vNumber := arg.Int("make-number", "M")
	if vNumber > 0 {
		var mkOptionList = []string{"uymas", "test"}
		for i := 0; i < vNumber; i++ {
			mkKey := str.RandStr.SafeStr(rand.Intn(40))
			mkQueue := []string{"--" + mkKey}
			if rand.Intn(4)%2 == 0 {
				mkQueue = append(mkQueue, fmt.Sprintf("%d", rand.Intn(999999)))
			}
			mkOptionList = append(mkOptionList, mkQueue...)
		}
		lgr.Info("创建生成测试命令如下：\n%s\n\n", strings.Join(mkOptionList, " "))
		fmt.Printf("消耗时间：%s\n", spendFn())
		return
	}

	fmt.Println()
	fmt.Println()
	lgr.Info("使用时间：%s\n", spendFn())
}
