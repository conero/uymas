package main

import (
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/cli/gen"
	"gitee.com/conero/uymas/v2/internal/recipe"
	"gitee.com/conero/uymas/v2/logger/lgr"
	"gitee.com/conero/uymas/v2/util/tm"
	"os"
)

type cmdHashOpt struct {
	Vtype string `cmd:"type,t help:支持md5, sha1, sha256, sha512等"`
}

func cmdHash(args cli.ArgsParser) {
	var opt cmdHashOpt
	err := gen.ArgsDress(args, &opt)
	if err != nil {
		lgr.Error(err.Error())
		return
	}
	timeDiffFn := tm.SpendFn()
	vPath := args.SubCommand()
	if vPath == "" {
		pwdDir, _ := os.Getwd()
		vPath = pwdDir
	}

	fh := &recipe.FileHash{
		Vtype: opt.Vtype,
	}
	list, err := fh.PathList(vPath)
	if err != nil {
		lgr.Error("%v", err)
		return
	}

	var tableData [][]string
	for _, ls := range list {
		tableData = append(tableData, []string{ls.Hash, ls.Filename})
	}

	if len(tableData) == 0 {
		lgr.Info("未发现文件：%s", vPath)
		return
	}

	lgr.Info("文件读取(%s)成功，如列表下：\n%v\n", fh.Vtype, tableData)
	lgr.Info("用时 %s", timeDiffFn())
}
