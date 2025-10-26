package main

import (
	"fmt"
	"math/rand"
	"strings"

	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/logger/lgr"
	"gitee.com/conero/uymas/v2/str"
	"gitee.com/conero/uymas/v2/util/tm"
)

func cmdTest(arg cli.ArgsParser) {
	if arg.Switch("for") {
		cmdTestFor(arg)
		return
	}
	defer tm.SpendDefer("本次执行耗时：")()
	if arg.Switch("verbose", "V") {
		fmt.Println()
		fmt.Println("参数解析，数据如下")
		fmt.Println()
		fmt.Printf("value: %v\n", arg.Values())
		fmt.Printf("mapValue: %v\n", arg.MapValue())
		fmt.Printf("option: %#v\n", arg.Option())
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
		return
	}
}

// for 版本测试
//
// 命令：for ($i = 0; $i -lt 20; $i++){$get = .\zig-out\bin\zuymas.exe -test -for 0.034597401B -sum -inline;echo "「$($i+1)」-> $get";};
func cmdTestFor(arg cli.ArgsParser) {
	spendFn := tm.SpendFn()
	vFor := arg.Int64Def(1_000_000_000, "for")
	var i int64 = 0
	var count uint64
	shouldSum := arg.Switch("sum")
	for ; i < vFor; i++ {
		if shouldSum {
			count += uint64(i) + 1
		}
	}

	// 行内测试
	if arg.Switch("inline", "I") {
		if shouldSum {
			fmt.Printf("本次耗时：%v， 累加值：%d，循环数 %d", spendFn(), count, vFor)
			return
		}

		fmt.Printf("本次耗时：%v，循环数 %d", spendFn(), vFor)
		return
	}
	if shouldSum {
		fmt.Printf("累加值：%v\n", count)
	}
	fmt.Printf("本次耗时：%v，循环数 %d\n", spendFn(), vFor)
}
