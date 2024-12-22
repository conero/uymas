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
	"sort"
)

type scanOption struct {
	Bufsize    int      `cmd:"bufsize,B default:5000 help:缓存数默认"`
	Exclude    []string `cmd:"exclude mark:... help:排除"`
	Include    []string `cmd:"include mark:... help:包含"`
	IsParallel bool     `cmd:"parallel,ll help:[实验性]的并行扫描"`
	NoSort     bool     `cmd:"not-sort help:禁止文件大小排序"`
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
	if er != nil {
		lgr.Error("目录扫描异常，%v", er)
		return
	}

	var table = [][]any{{"Path", "Size", "Depth"}}
	if opt.NoSort {
		for key, tcd := range dd.TopChildDick {
			table = append(table, []any{key, number.Bytes(tcd.Size), tcd.Depth})
		}
	} else {
		var diskAttr = map[int64]scan.ChildDirData{}
		var sizeList []int64
		var repeatArr = map[int64][]scan.ChildDirData{}
		for _, tcd := range dd.TopChildDick {
			_, exist := diskAttr[tcd.Size]
			if exist {
				repeat := repeatArr[tcd.Size]
				repeat = append(repeat, tcd)
				repeatArr[tcd.Size] = repeat
				continue
			}
			diskAttr[tcd.Size] = tcd
			sizeList = append(sizeList, tcd.Size)
		}
		sort.Slice(sizeList, func(i, j int) bool {
			return sizeList[i] > sizeList[j]
		})
		for _, vSize := range sizeList {
			tcd := diskAttr[vSize]
			table = append(table, []any{tcd.Name, number.Bytes(tcd.Size), tcd.Depth})
			repeat := repeatArr[tcd.Size]
			for _, same := range repeat {
				table = append(table, []any{same.Name, number.Bytes(same.Size), same.Depth})
			}
		}
	}

	fmt.Println(rock.FormatTable(table))
	fmt.Println()
	fmt.Printf(" 文件扫目标目录： %v，是否并发: %v(线程分配 %v).\n", dd.BaseDir(), isParallel, dd.ChanNumber())
	fmt.Printf(" 文件扫描数： %v, 目录: %v, 文件： %v.\n", dd.AllItem, dd.AllDirItem, dd.AllFileItem)
	fmt.Printf(" 目录大小: %v.\n", number.Bytes(dd.AllSize))
	fmt.Printf(" 使用时间： %v.\n", dd.Runtime)

	fmt.Printf(" 内存消耗：%v，用时:%s\n", memSubCall(), spendFn())
}
