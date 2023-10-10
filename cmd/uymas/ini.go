package main

import (
	"encoding/json"
	"fmt"
	"gitee.com/conero/uymas"
	"gitee.com/conero/uymas/bin"
	"gitee.com/conero/uymas/logger/lgr"
	"gitee.com/conero/uymas/parser/xini"
	"gitee.com/conero/uymas/str"
	"gitee.com/conero/uymas/util"
	"math/rand"
	"time"
)

type ActionIni struct {
	bin.CliApp
}

func (c *ActionIni) DefaultHelp() {
	fmt.Println("  ini 命令，实现对文件ini文件的加载解析")
	fmt.Println("  ini create   ini文件随机生成")
	fmt.Println("      --num,-N [1000]  设置需要生成的数量集合")
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

// Create 文件创建，用于测试
func (c *ActionIni) Create() {
	tmMark := util.SpendTimeDiff()
	number := c.Cc.ArgInt("num", "N")
	if number < 1 {
		number = 1000
	}

	var data = map[string]any{}
	tyNum := 5
	var rdStr str.RandString
	for i := 0; i < number; i++ {
		vCount := rand.Intn(50)
		if vCount < 2 {
			vCount = 2
		}
		key := fmt.Sprintf("%s%v", rdStr.SafeStr(rand.Intn(25)), i)
		var value any
		switch rand.Intn(tyNum) {
		case 0: // i64
			value = rand.Intn(999999)
		case 1: //bool
			value = false
			if rand.Intn(2) == 0 {
				value = true
			}
		case 3: // float-64
			value = float64(rand.Intn(10000)) * rand.Float64()
		case 4: // []int64
			var i64Que []int64
			for j := 0; j < vCount; j++ {
				i64Que = append(i64Que, int64(rand.Intn(999999)))
			}
			value = i64Que
		case 5: // []float64
			var i64Que []float64
			for j := 0; j < vCount; j++ {
				i64Que = append(i64Que, float64(rand.Intn(10000))*rand.Float64())
			}
			value = i64Que
		case 6: // []string
			var i64Que []string
			for j := 0; j < vCount; j++ {
				i64Que = append(i64Que, rdStr.SafeStr(rand.Intn(15)))
			}
			value = i64Que
		default: // string
			value = rdStr.SafeStr(rand.Intn(50))
		}
		data[key] = value
	}

	by, err := xini.Marshal(data)
	if err != nil {
		lgr.Error("data 序列化为ini错误，%v", err.Error())
		return
	}

	fmt.Println("; ------------------[ini]---------------------")
	fmt.Printf("; CREATE AT %s, BY %v/%v.\n", time.Now().Format("2006-01-02 15:04:05"), uymas.Version, uymas.Release)
	fmt.Println()
	fmt.Println()
	fmt.Printf("%s", string(by))
	fmt.Println()
	fmt.Println()
	fmt.Println("; ")
	fmt.Printf("; 本次生成ini变量个数 %v，用时 %v.\n", number, tmMark())

}

func (c *ActionIni) DefaultIndex() {
	c.DefaultHelp()

}
