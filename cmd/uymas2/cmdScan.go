package main

import (
	"fmt"
	"gitee.com/conero/uymas/v2/app/scan"
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/cli/gen"
	"gitee.com/conero/uymas/v2/logger/lgr"
	"gitee.com/conero/uymas/v2/number"
	"gitee.com/conero/uymas/v2/rock"
	"gitee.com/conero/uymas/v2/util/fs"
	"gitee.com/conero/uymas/v2/util/tm"
)

type scanOption struct {
	Bufsize    int      `cmd:"bufsize,B default:5000 help:缓存数默认"`
	Exclude    []string `cmd:"exclude help:排除"`
	Include    []string `cmd:"include help:包含"`
	IsParallel bool     `cmd:"parallel,ll help:[实验性]的并行扫描"`
	Dir        string   `cmd:"dir isdata"`
}

func cmdScan(args cli.ArgsParser) {
	spendFn := tm.SpendFn()
	var opt scanOption
	err := gen.ArgsDress(args, &opt)
	if err != nil {
		lgr.Error(err.Error())
		return
	}

	var mU fs.MemUsage
	memSubCall := mU.GetSysMemSub()
	baseDir := opt.Dir
	if baseDir == "" {
		baseDir = "./"
	}
	dd := scan.NewDirScanner(baseDir)
	dd.CddChanMax = opt.Bufsize

	//过滤
	dd.Exclude(opt.Exclude...)
	dd.Include(opt.Include...)

	var er error
	var isParallel = "否"
	if opt.IsParallel {
		er = dd.ScanParallel()
		isParallel = "是"
	} else {
		er = dd.Scan()
	}

	if er == nil {
		var table = [][]any{{"Path", "Size", "Depth"}}
		for key, tcd := range dd.TopChildDick {
			table = append(table, []any{key, number.Bytes(tcd.Size), tcd.Depth})
		}

		fmt.Println(rock.FormatTable(table))
		fmt.Println()
		fmt.Printf(" 文件扫目标目录： %v，是否并发: %v(线程分配 %v).\n", dd.BaseDir(), isParallel, dd.ChanNumber())
		fmt.Printf(" 文件扫描数： %v, 目录: %v, 文件： %v.\n", dd.AllItem, dd.AllDirItem, dd.AllFileItem)
		fmt.Printf(" 目录大小: %v.\n", number.Bytes(dd.AllSize))
		fmt.Printf(" 使用时间： %v.\n", dd.Runtime)
	}
	fmt.Printf(" 内存消耗：%v，用时:%s\n", memSubCall(), spendFn())
}
