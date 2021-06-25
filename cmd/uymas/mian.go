package main

import (
	"bufio"
	"fmt"
	"github.com/conero/uymas"
	"github.com/conero/uymas/bin"
	"github.com/conero/uymas/bin/parser"
	"github.com/conero/uymas/culture/pinyin"
	"github.com/conero/uymas/fs"
	"github.com/conero/uymas/number"
	"github.com/conero/uymas/storage"
	"github.com/conero/uymas/util"
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

//the cli app tools
func application() {
	cli = bin.NewCLI()
	//pinyin
	cli.RegisterFunc(func(cc *bin.CliCmd) {
		words := cc.SubCommand
		if words != "" {
			pinyinCache = getPinyin()
			fmt.Println(pinyinCache.GetPyTone(words))
		} else {
			fmt.Println("请输入: $ pinyin <汉字>")
		}
	}, "pinyin")

	//empty
	//option: -v,--version; -h,--help
	cli.RegisterEmpty(func(cc *bin.CliCmd) {
		if cc.CheckSetting("v", "version") {
			fmt.Printf("v%v/%v\r\n", uymas.Version, uymas.Release)
			return
		} else if cc.CheckSetting("h", "help") {
			cc.CallCmd("help")
			return
		}

		fmt.Printf(" wecolme use the <%v>. \r\n", uymas.Name)
		fmt.Printf(" %v/%v\r\n", uymas.Version, uymas.Release)
		fmt.Printf(" Power by %v.\r\n", uymas.Author)
	})

	//uls,uymas ls
	cli.RegisterFunc(func(cc *bin.CliCmd) {
		fmt.Println("  " + strings.Join(cli.GetCmdList(), "\r\n  "))
	}, "uls", "uymas-ls")

	//cache namespace@key.key setValue
	cli.RegisterFunc(func(cc *bin.CliCmd) {
		key := cc.SubCommand
		value := cc.Next(key)
		hasCache, ccValue := getCache(key, value)
		if value != "" {
			if hasCache {
				fmt.Printf("%v\r\n", ccValue)
			} else {
				fmt.Printf("%v 没有设置值\r\n", key)
			}
		} else {
			fmt.Printf("%v, %v 键值对已保存!\r\n", key, value)
		}
	}, "cache", "cc")

	//scan, sc
	cli.RegisterFunc(func(cc *bin.CliCmd) {
		memSubCall := gMu.GetSysMemSub()
		baseDir := cc.SubCommand
		if baseDir == "" {
			baseDir = "./"
		}
		dd := fs.NewDirScanner(baseDir)

		//过滤
		dd.Exclude(cc.ArgRaw("exclude"))
		dd.Include(cc.ArgRaw("include"))

		if er := dd.Scan(); er == nil {
			var table [][]interface{}
			for key, tcd := range dd.TopChildDick {
				table = append(table, []interface{}{key, number.Bytes(tcd.Size)})
			}

			fmt.Println(bin.FormatTable(table, " "))
			fmt.Printf(" 文件扫目标目录： %v.\r\n", dd.BaseDir())
			fmt.Printf(" 文件扫描数： %v, 目录: %v, 文件： %v.\r\n", dd.AllItem, dd.AllDirItem, dd.AllFileItem)
			fmt.Printf(" 目录大小: %v.\r\n", number.Bytes(dd.AllSize))
			fmt.Printf(" 使用时间： %v.\r\n", dd.Runtime)
		}
		fmt.Printf(" 内存消耗：%v\r\n", memSubCall())

	}, "scan", "sc")

	//REPL
	cli.RegisterFunc(func(cc *bin.CliCmd) {
		var input = bufio.NewScanner(os.Stdin)
		fmt.Println("驻留式命令行程序")
		fmt.Print("$ uymas>")

		for input.Scan() {
			text := input.Text()
			text = strings.TrimSpace(text)

			switch text {
			case "exit":
				os.Exit(0)
			default:
				var cmdsList = parser.NewParser(text)
				for _, cmdArgs := range cmdsList {
					cli.Run(cmdArgs...)
				}
			}

			fmt.Println()
			fmt.Println()
			fmt.Print("$ uymas>")
		}
	}, "repl")

	//help
	cli.RegisterFunc(func(cc *bin.CliCmd) {
		lang := cc.ArgRaw("lang", "L")
		content := bin.GetHelpEmbed(s, lang)
		fmt.Println(content)
	}, "help", "?")

	//test 用于命令行解析等测试
	cli.RegisterFunc(func(cc *bin.CliCmd) {
		fmt.Println(" 命令行测试")
		fmt.Printf("  SubCommand: %v \r\n", cc.SubCommand)
		fmt.Printf("  Option: %v \r\n", cc.Setting)
		fmt.Printf("  DataRaw: %v \r\n", cc.DataRaw)
		fmt.Printf("  Data: %#v \r\n", cc.Data)
		fmt.Printf("  Input: %#v \r\n", strings.Join(cc.Raw, " "))
		fmt.Printf(" %v \r\n", getSpendStr())
		fmt.Println()
	}, "test")

	//app App 应用
	cli.RegisterApp(new(App), "app")

	cli.Run()
}

func getPinyin() *pinyin.Pinyin {
	if pinyinCache == nil {
		pinyinCache = pinyin.NewPinyin("./resource/culture/pinyin.txt")
	}
	return pinyinCache
}

//the uymas cmd message
func main() {
	application()
}

//获取的缓存
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

//消耗时间、内存等计算
func getSpendStr() string {
	return fmt.Sprintf("时间和内存消耗，用时 %v, 内存消耗 %v", gSpendTm(), gSpendMem())
}

func init() {
	//时间统计
	gSpendMem = gMu.GetSysMemSub()
	gSpendTm = util.SpendTimeDiff()
}
