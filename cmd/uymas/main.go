package main

import (
	"fmt"
	"gitee.com/conero/uymas/v2/bin"
	"gitee.com/conero/uymas/v2/culture/pinyin"
	"gitee.com/conero/uymas/v2/fs"
	"gitee.com/conero/uymas/v2/number"
	"gitee.com/conero/uymas/v2/storage"
	"gitee.com/conero/uymas/v2/util"
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
	//app App 应用
	cli.RegisterApp(new(App), "app")
	cli.RegisterApp(new(ActionIni), "ini")
	cli.RegisterAny(&defaultApp{})
	cli.Run()
}

func getPinyin() *pinyin.Pinyin {
	if pinyinCache == nil {
		pinyinCache = pinyin.NewPinyin("./resource/culture/pinyin.txt")
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
	if strings.Index(key, nsSplit) > -1 {
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
