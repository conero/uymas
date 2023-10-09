package main

import (
	"encoding/json"
	"fmt"
	"gitee.com/conero/uymas/bin"
	"gitee.com/conero/uymas/logger/lgr"
	"gitee.com/conero/uymas/parser/xini"
	"gitee.com/conero/uymas/util"
)

type ActionIni struct {
	bin.CliApp
}

func (c *ActionIni) DefaultHelp() {
	fmt.Println("  ini 命令，实现对文件ini文件的加载解析")
	fmt.Println("  ini [file]   文件解析")
	fmt.Println("      --output,-O   是否打印内容")
	fmt.Println("      --restore,-R  反序列恢复")
}

func (c *ActionIni) DefaultUnmatched() {
	file := c.Cc.SubCommand
	if file == "" {
		return
	}

	lgr.Info("正在读取文件 %s ……", file)

	timeTck := util.SpendTimeDiff()
	psr := xini.NewParser()
	psr.OpenFile(file)
	if !psr.IsValid() {
		lgr.Error("文件加载失败!\n  %s", psr.ErrorMsg())
		return
	}

	// 计算成功后显示信息
	// @todo 显示文件大写之类的。
	lgr.Info("文件加载成功！\n  用时：%s", timeTck())

	isOut := c.Cc.CheckSetting("output", "O")
	isRestore := c.Cc.CheckSetting("restore", "R")
	allData := psr.GetData()
	if isOut {
		jsonBy, jsonEr := json.Marshal(allData)

		lgr.Info("解析后的内容：\n\n--------------[RAW]--------------\n%#v\n\n--------------[JSON]--------------\n%s",
			allData, string(jsonBy))
		if jsonEr != nil {
			lgr.Error("json 编码错误，%s", jsonEr.Error())
		}
	}
	if isRestore {
		bys, err := xini.Marshal(allData)
		if err != nil {
			lgr.Error("xini 序列化生成错误，%s", err.Error())
		} else {
			lgr.Info("xini 序列化生成\n\n------------[ini]-----------\n%s\n", string(bys))
		}
	}
}

func (c *ActionIni) DefaultIndex() {
	c.DefaultHelp()

}
