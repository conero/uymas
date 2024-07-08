package main

import (
	"fmt"
	"gitee.com/conero/uymas/bin"
	"gitee.com/conero/uymas/culture/pinyin"
	"gitee.com/conero/uymas/culture/pinyin/material"
	"gitee.com/conero/uymas/fs"
	"gitee.com/conero/uymas/number"
	"gitee.com/conero/uymas/storage"
	"gitee.com/conero/uymas/util"
	"os"
	"strings"
	"time"
)

var (
	cli         *bin.CLI
	pinyinCache *pinyin.Pinyin = nil
	gMu         fs.MemUsage
	gSpendTm    func() time.Duration
	gSpendMem   func() number.BitSize
)

// the cli app tools
func application() {
	cli = bin.NewCLI()
	// app App 应用
	cli.RegisterApp(new(App), "app")
	cli.RegisterApp(new(ActionIni), "ini")
	cli.RegisterAny(&defaultApp{})

	// 设置 cli
	long := os.Getenv("UYMAS_CMD_UYMAS_LONG")
	colon := os.Getenv("UYMAS_CMD_UYMAS_COLON")
	cli.RunWith(bin.ArgConfig{
		LongOption: strings.ToLower(long) != "false" && strings.ToLower(long) != "0",
		EqualColon: strings.ToLower(colon) == "true" || strings.ToLower(colon) == "1",
	})
}

func getPinyin() *pinyin.Pinyin {
	if pinyinCache == nil {
		py := material.NewPinyin()
		pinyinCache = &py
	}
	return pinyinCache
}

// the uymas cmd message
func main() {
	application()
}

// 获取的缓存
func getCache(key, value string) (bool, storage.Any) {
	var namespace string
	var nsSplit = "@"
	if strings.Contains(key, nsSplit) {
		tapQueue := strings.Split(key, nsSplit)
		namespace = strings.TrimSpace(tapQueue[0])
		key = strings.TrimSpace(tapQueue[1])
	}

	store := storage.GetStorage(namespace)
	if value == "" {
		if store != nil {
			return true, store.GetValue(key)
		}
		return false, ""
	} else {
		if store == nil {
			store = storage.NewStorage(namespace)
		}
		return true, store.SetValue(key, storage.NewLite(value).GetAny())
	}
}

// 消耗时间、内存等计算
func getSpendStr() string {
	return fmt.Sprintf("时间和内存消耗，用时 %v, 内存消耗 %v", gSpendTm(), gSpendMem())
}

func init() {
	//时间统计
	gSpendMem = gMu.GetSysMemSub()
	gSpendTm = util.SpendTimeDiff()
}
