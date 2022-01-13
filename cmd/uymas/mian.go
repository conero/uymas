package main

import (
	"bufio"
	"fmt"
	"gitee.com/conero/uymas"
	"gitee.com/conero/uymas/bin"
	"gitee.com/conero/uymas/bin/butil"
	"gitee.com/conero/uymas/bin/parser"
	"gitee.com/conero/uymas/culture/pinyin"
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

//the cli app tools
func application() {
	cli = bin.NewCLI()
	//app App 应用
	cli.RegisterApp(new(App), "app")
	cli.RegisterAny(&defaultApp{})
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

type defaultApp struct {
	bin.CliApp
}

func (c *defaultApp) Construct() {
	cc := c.Cc
	cc.CommandAliasAll(map[string][]string{
		"pinyin": {"py"},
		"scan":   {"sc"},
		"cache":  {"cc"},
		"uls":    {"uymas-ls"},
	})
	/*cc.CommandAlias("pinyin", "py").
	CommandAlias("scan", "sc").
	CommandAlias("cache", "cc").
	CommandAlias("uls", "uymas-ls")*/
}

// DefaultIndex index
func (c *defaultApp) DefaultIndex() {
	cc := c.Cc
	if cc.CheckSetting("v", "version") {
		fmt.Printf("v%v/%v\n", uymas.Version, uymas.Release)
		return
	} else if cc.CheckSetting("h", "help") {
		cc.CallCmd("help")
		return
	}

	fmt.Printf(" wecolme use the <%v>. \n", uymas.Name)
	fmt.Printf(" %v/%v\n", uymas.Version, uymas.Release)
	fmt.Printf(" Power by %v.\n", uymas.Author)
}

// DefaultHelp help
func (c *defaultApp) DefaultHelp() {
	cc := c.Cc
	lang := cc.ArgRaw("lang", "L")
	content := bin.GetHelpEmbed(s, lang)
	fmt.Println(content)
}

func (c *defaultApp) DefaultUnmatched() {
	fmt.Printf(" Oop, 命令'%v'还没有实现呢，老兄！\n", c.Cc.Command)
	fmt.Println()
}

// Test test 用于命令行解析等测试
func (c *defaultApp) Test() {
	cc := c.Cc
	pwd, _ := os.Getwd()

	fmt.Println(" 命令行测试")
	fmt.Printf("  SubCommand: %v \n", cc.SubCommand)
	fmt.Printf("  Option: %#v \n", cc.Setting)
	fmt.Printf("  Raw: %#v \n", cc.Raw)
	fmt.Printf("  DataRaw: %#v \n", cc.DataRaw)
	fmt.Printf("  Data: %#v \n", cc.Data)
	fmt.Printf("  Input: %#v \n", strings.Join(cc.Raw, " "))
	fmt.Printf("  Basedir : %v \n", butil.GetBasedir())
	fmt.Printf("  PWD : %v \n", pwd)

	fmt.Printf(" %v \n", getSpendStr())
	fmt.Println()
}

// Repl REPL
func (c *defaultApp) Repl() {
	var input = bufio.NewScanner(os.Stdin)
	fmt.Println("驻留式命令行程序")
	fmt.Print("$ uymas> ")

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
		fmt.Print("$ uymas> ")
	}
}

// Pinyin 拼音
func (c *defaultApp) Pinyin() {
	cc := c.Cc
	words := cc.SubCommand
	if words != "" {
		pinyinCache = getPinyin()
		fmt.Println(pinyinCache.GetPyTone(words))
	} else {
		fmt.Println("请输入: $ pinyin <汉字>")
	}
}

// Uls command allis: uls,uymas ls
func (c *defaultApp) Uls() {
	fmt.Println("  " + strings.Join(cli.GetCmdList(), "\n  "))
}

// Cache cache namespace@key.key setValue
func (c *defaultApp) Cache() {
	cc := c.Cc
	key := cc.SubCommand
	value := cc.Next(key)
	hasCache, ccValue := getCache(key, value)
	if value != "" {
		if hasCache {
			fmt.Printf("%v\n", ccValue)
		} else {
			fmt.Printf("%v 没有设置值\n", key)
		}
	} else {
		fmt.Printf("%v, %v 键值对已保存!\n", key, value)
	}
}

// Scan scan,sc
func (c *defaultApp) Scan() {
	cc := c.Cc
	memSubCall := gMu.GetSysMemSub()
	baseDir := cc.SubCommand
	if baseDir == "" {
		baseDir = "./"
	}
	dd := fs.NewDirScanner(baseDir)
	dd.CddChanMax = cc.ArgInt("bufsize", "B")

	//过滤
	dd.Exclude(cc.ArgStringSlice("exclude")...)
	dd.Include(cc.ArgStringSlice("include")...)

	var er error
	var isParallel = "否"
	if cc.CheckSetting("parallel", "ll") {
		er = dd.ScanParallel()
		isParallel = "是"
	} else {
		er = dd.Scan()
	}

	if er == nil {
		var table = [][]interface{}{{"Path", "Size", "Depth"}}
		for key, tcd := range dd.TopChildDick {
			table = append(table, []interface{}{key, number.Bytes(tcd.Size), tcd.Depth})
		}

		fmt.Println(bin.FormatTable(table, false))
		fmt.Printf(" 文件扫目标目录： %v，是否并发: %v(线程分配 %v).\n", dd.BaseDir(), isParallel, dd.ChanNumber())
		fmt.Printf(" 文件扫描数： %v, 目录: %v, 文件： %v.\n", dd.AllItem, dd.AllDirItem, dd.AllFileItem)
		fmt.Printf(" 目录大小: %v.\n", number.Bytes(dd.AllSize))
		fmt.Printf(" 使用时间： %v.\n", dd.Runtime)
	}
	fmt.Printf(" 内存消耗：%v\n", memSubCall())

}

func init() {
	//时间统计
	gSpendMem = gMu.GetSysMemSub()
	gSpendTm = util.SpendTimeDiff()
}
