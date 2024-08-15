package main

import (
	"encoding/base64"
	"fmt"
	"gitee.com/conero/uymas"
	"gitee.com/conero/uymas/bin"
	"gitee.com/conero/uymas/bin/butil"
	"gitee.com/conero/uymas/bin/color"
	"gitee.com/conero/uymas/bin/data"
	"gitee.com/conero/uymas/bin/doc"
	"gitee.com/conero/uymas/culture/digit"
	"gitee.com/conero/uymas/culture/ganz"
	"gitee.com/conero/uymas/culture/pinyin"
	"gitee.com/conero/uymas/fs"
	"gitee.com/conero/uymas/logger/lgr"
	"gitee.com/conero/uymas/number"
	"gitee.com/conero/uymas/str"
	"gitee.com/conero/uymas/util"
	"gitee.com/conero/uymas/util/rock"
	"io"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type defaultApp struct {
	bin.CliApp
	isVerbose bool // 详细输出
}

func (c *defaultApp) Construct() {
	cc := c.Cc
	cc.CommandAliasAll(map[string][]string{
		"pinyin":   {"py"},
		"scan":     {"sc"},
		"cache":    {"cc"},
		"uls":      {"uymas-ls"},
		"digit":    {"dg"},
		"datediff": {"dd"},
		"base64":   {"b64"},
	})
	/*cc.CommandAlias("pinyin", "py").
	CommandAlias("scan", "sc").
	CommandAlias("cache", "cc").
	CommandAlias("uls", "uymas-ls")*/
	c.isVerbose = c.Cc.CheckSetting("vv", "verbose")
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
	fmt.Println()
	fmt.Printf(" %v [comand] [option]    执行应用命令\n", butil.AppName())
	fmt.Println()
	fmt.Printf(" v%v/%v\n", uymas.Version, uymas.Release)
	fmt.Printf(" Power by %v.\n", uymas.Author)
	fmt.Println()
	fmt.Println(" 命令行设置，通环境变量如:")
	fmt.Println("   UYMAS_CMD_UYMAS_COLON=false # 支持 ':' 等式，或者0/1")
	fmt.Println("   UYMAS_CMD_UYMAS_LONG=true # 支持长选项和段选项，使用 false/0 关闭")
}

// DefaultHelp help
func (c *defaultApp) DefaultHelp() {
	lng := c.Cc.ArgRaw("lang", "L")
	help := doc.FromLine(s, lng)
	if lng != "" && !help.Support(lng) {
		lgr.Error("%s 不支持该语言", lng)
		return
	}

	// help
	fnCmd := c.Cc.SubCommand
	if fnCmd == "" {
		img := help.Help(lng)
		fmt.Println(img)
		return
	}

	helpDoc, exist := help.Search(fnCmd)
	if !exist {
		lgr.Error("%s Not Find", fnCmd)
		return
	}

	fmt.Println(helpDoc)
}

func (c *defaultApp) DefaultUnmatched() {
	fmt.Printf(" Oop, 命令'%v'还没有实现呢，老兄！\n", c.Cc.Command)
	fmt.Println()
}

// Test 用于命令行解析等测试
func (c *defaultApp) Test() {
	cc := c.Cc
	pwd, _ := os.Getwd()
	dataMng := data.CliManager()

	fmt.Println(" 命令行测试")
	fmt.Printf("  SubCommand: %v \n", cc.SubCommand)
	fmt.Printf("  Option: %#v \n", cc.Setting)
	fmt.Printf("  Raw: %#v \n", cc.Raw)
	fmt.Printf("  DataRaw: %#v \n", cc.DataRaw)
	fmt.Printf("  Data: %#v \n", cc.Data)
	fmt.Printf("  Input: %#v \n", strings.Join(cc.Raw, " "))
	fmt.Printf("  Basedir : %v \n", butil.Basedir())
	fmt.Printf("  DataManager : %v \n", dataMng.Dir())
	fmt.Printf("  PWD : %v \n", pwd)
	fmt.Printf("  NextLing : %#v \n", cc.NextList())
	fmt.Printf("  Args : %#v \n", os.Args[1:])
	fmt.Printf("  ArgCofnig : %#v \n", c.Cc.Config())

	// var
	varName := cc.ArgRaw("var", "v")
	if varName != "" {
		fmt.Printf("  Var : %s =>  %s \n", varName, cc.ArgRaw(varName))
		fmt.Printf("  Var/Raw : %s =>  %#v \n", varName, cc.Arg(varName))
	}

	fmt.Println()
	fmt.Printf(" %v \n", getSpendStr())
	fmt.Println()
	fmt.Println()
	fmt.Println("输入命令 “--var,-v $name”   可用于读取 $name 的option参数")
	fmt.Println("环境变量（powershell 版本）：\n  $env:UYMAS_CMD_UYMAS_LONG=1/true\n  $env:UYMAS_CMD_UYMAS_COLON=0/false")
}

// Repl REPL
func (c *defaultApp) Repl() {
	bin.ReplCommand("uymas", cli)
}

// Pinyin 拼音
func (c *defaultApp) Pinyin() {
	tmSpend := util.SpendTimeDiff()
	cc := c.Cc
	words := cc.SubCommand
	fl := cc.ArgRaw("file", "f")
	if fl != "" {
		bys, err := os.ReadFile(fl)
		if err != nil {
			lgr.Error("读取文件错误，%s", color.StyleByAnsi(color.AnsiTextRed, err))
			return
		}
		lgr.Info("已读取文件 %s 的内容", fl)
		words = string(bys)
	}

	if words == "" {
		words = butil.InputRequire("请输入中文汉字：", nil)
	}

	if c.Cc.CheckSetting("from-unicode", "fu", "U", "unicode") {
		lgr.Info("%s", str.ParseUnicode(words))
		return
	}

	words = str.ParseUnicode(words)

	// 输出utf16 代码
	if c.Cc.CheckSetting("utf16", "utf") {
		var codeList []string
		var strList []string
		for _, r := range []rune(words) {
			codeList = append(codeList, fmt.Sprintf("U+%s", strconv.FormatInt(int64(r), 16)))
			strList = append(strList, fmt.Sprintf("\\u%s", strconv.FormatInt(int64(r), 16)))
		}
		lgr.Info("%s 转utf16如：\n %s\n", words,
			color.StyleByAnsi(color.AnsiTextGreen, strings.Join(strList, "")))

		if c.isVerbose {
			lgr.Info("%s 转utf16 (unicode 风格)如：\n %s\n", words,
				color.StyleByAnsi(color.AnsiTextGreen, strings.Join(codeList, " ")))
		}
	}

	pinyinCache = getPinyin()

	// 拼音搜索
	pinyinArgs := []string{"pyin", "P"}
	searchAlpha := c.Cc.ArgRaw(pinyinArgs...)
	if searchAlpha == "" {
		searchAlpha = words
	}
	if c.Cc.CheckSetting(pinyinArgs...) {
		limit := c.Cc.DefInt(pinyin.SearchAlphaLimit, "limit", "L")
		list := pinyinCache.SearchAlpha(searchAlpha, limit)
		textList := list.Text()
		//lgr.Info("搜索到字符如下：\n%s\n", bin.FormatQue(textList))
		lgr.Info("搜索到字符如下：\n%s\n", strings.Join(textList, " "))
		fmt.Printf("\n   搜索到 %d 个字，限制搜索字数字 %d，用时 %v\n", len(textList), limit, tmSpend())
		return
	}

	isOld := c.Cc.CheckSetting("old")
	vFmt := ""
	switch c.Cc.DefString("", "fmt", "F") {
	case "title":
		vFmt = pinyin.SepTitle
	}
	seps := []string{
		c.Cc.DefString("", "sep", "S"),
		vFmt,
	}
	var line string
	if c.Cc.CheckSetting("number", "n") {
		if isOld {
			line = pinyinCache.GetPyToneNumber(words)
		} else {
			line = pinyinCache.SearchByGroup(words).Number(seps...)
		}
	} else if c.Cc.CheckSetting("alpha", "a") {
		if isOld {
			line = pinyinCache.GetPyToneAlpha(words)
		} else {
			line = pinyinCache.SearchByGroup(words).Alpha(seps...)
		}
	} else if c.Cc.CheckSetting("all", "A") {
		if isOld {
			line = "原始拼音：" + pinyinCache.GetPyTone(words) + "\n" +
				"数字声调拼音：" + pinyinCache.GetPyToneNumber(words) + "\n" +
				"字母拼音：" + pinyinCache.GetPyToneAlpha(words)
		} else {
			vList := pinyinCache.SearchByGroup(words)
			line = "原始拼音：" + vList.Tone(seps...) + "\n" +
				"数字声调拼音：" + vList.Number(seps...) + "\n" +
				"字母拼音：" + vList.Alpha(seps...)

			// 多音字
			pyList := vList.Polyphony(pinyin.PinyinTone)
			pyCount := len(pyList)
			if len(pyList) > 0 {
				line += fmt.Sprintf("\n  多音字共 %d 组，详细如：\n原始拼音组：\n%s\n数字拼音组：\n%s\n字母拼音组：\n%s\n",
					pyCount, bin.FormatQue(pyList, "  "),
					bin.FormatQue(vList.Polyphony(pinyin.PinyinNumber), "  "),
					bin.FormatQue(rock.ListNoRepeat(vList.Polyphony(pinyin.PinyinAlpha)), "  "))
			}
		}
	} else {
		if isOld {
			line = pinyinCache.GetPyTone(words)
		} else if c.isVerbose {
			list := pinyinCache.SearchByGroup(words)
			line = bin.FormatQue(list.Polyphony(pinyin.PinyinTone, seps...), " ")
		} else {
			line = pinyinCache.SearchByGroup(words).Tone(seps...)
		}

	}

	fmt.Println(line)
	fmt.Printf("\n   字长%d（Unicode），用时 %v\n", len(words), tmSpend())
}

// Uls command allis: uls,uymas ls
func (c *defaultApp) Uls() {
	fmt.Println("  " + strings.Join(cli.GetCmdList(), "\n  "))

	// plugin-sub-command 支持
	plugs := bin.PlgCmdList()
	if len(plugs) > 0 {
		fmt.Println("\nPlugin Sub command:\n" + "  " + strings.Join(plugs, "\n  "))
	}
}

// Cache namespace@key.key setValue
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
		var table = [][]any{{"Path", "Size", "Depth"}}
		for key, tcd := range dd.TopChildDick {
			table = append(table, []any{key, number.Bytes(tcd.Size), tcd.Depth})
		}

		fmt.Println(bin.FormatTable(table, false))
		fmt.Printf(" 文件扫目标目录： %v，是否并发: %v(线程分配 %v).\n", dd.BaseDir(), isParallel, dd.ChanNumber())
		fmt.Printf(" 文件扫描数： %v, 目录: %v, 文件： %v.\n", dd.AllItem, dd.AllDirItem, dd.AllFileItem)
		fmt.Printf(" 目录大小: %v.\n", number.Bytes(dd.AllSize))
		fmt.Printf(" 使用时间： %v.\n", dd.Runtime)
	}
	fmt.Printf(" 内存消耗：%v\n", memSubCall())

}

// Digit 数字转
func (c *defaultApp) Digit() {
	value := c.Cc.SubCommand
	if value == "" {
		lgr.Error("请指定阿拉伯数字或中文数字！")
		return
	}

	isMatch, _ := regexp.MatchString(`\d+(.?\d+)?`, value)
	if isMatch {
		lgr.Info("识别为：阿拉伯数字转中文数字")
		vNum, err := strconv.ParseFloat(value, 64)
		if err != nil {
			lgr.Error("%s 不是有效数字!", value)
			return
		}
		isRmb := c.Cc.CheckSetting("rmb", "r")
		var cv = digit.Cover(vNum)
		if c.Cc.CheckSetting("lower", "l") {
			var valueStr string
			if isRmb {
				valueStr = cv.ToRmbLower()
			} else {
				valueStr = cv.ToChnRoundLower()
			}
			lgr.Info("转化中文小写数字成功！\n\n %v", valueStr)
		}
		if c.Cc.CheckSetting("both", "b") {
			if isRmb {
				lgr.Info("转化中文大小写数字成功！\n\n %v\n %v\n %v\n %v",
					cv.ToChnRoundUpper(), cv.ToChnRoundLower(), cv.ToRmbUpper(), cv.ToRmbLower())
				return
			}
			lgr.Info("转化中文大小写数字成功！\n\n %v\n %v", cv.ToChnRoundUpper(), cv.ToChnRoundLower())
			return
		}

		var valueStr string
		if isRmb {
			valueStr = cv.ToRmbUpper()
		} else {
			valueStr = cv.ToChnRoundUpper()
		}
		lgr.Info("转化中文大写数成功！\n\n %v", valueStr)
		return
	}

}

func (c *defaultApp) DefaultEnd() {
	fmt.Println()
}

// Datediff 时间日期差计算
func (c *defaultApp) Datediff() {
	date := c.Cc.SubCommand
	if date == "" {
		lgr.Error("请输入日期！")
		return
	}
	endDate := c.Cc.ArgRaw("end", "e")

	// 日期解析
	var tm time.Time
	var err error
	if strings.ToLower(date) == "now" {
		tm = time.Now()
	} else {
		tm, err = str.TimeParse(date)
		if err != nil {
			lgr.Error("日期格式不支持！")
			return
		}
	}

	// 日期运算
	add := c.Cc.ArgRaw("add", "a")
	if add != "" {
		dura, err := str.ParseDuration(add)
		if err != nil {
			lgr.Error("加入日期错误，%s", err)
			return
		}

		newTm := tm.Add(dura)
		lgr.Info("当前时间：%s, 运算：%s，得\n    %s",
			color.StyleByAnsi(color.AnsiTextBlackBr, tm.Format(time.DateTime)),
			color.StyleByAnsi(color.AnsiTextBlackBr, add+"("+dura.String()+")"),
			color.StyleByAnsi(color.AnsiTextGreenBr, newTm.Format(time.DateTime)))
		return
	}

	now := time.Now()
	if endDate != "" {
		tmEnd, err := str.TimeParse(endDate)
		if err != nil {
			lgr.Error("日期格式不支持！")
			return
		}
		now = tmEnd
	}
	diff := tm.Sub(now)
	d3 := NewD3(diff)

	var diffType string
	if diff > 0 {
		diffType = "之后"
	} else {
		diffType = "之前"
	}

	// 输出
	cmdLsString := d3.cmdListing()
	if cmdLsString != "" {
		cmdLsString = " " + cmdLsString + "\n"
	}
	lgr.Info("%s 距今(%s)比较：\n 时间总差：%v\n 差别类型：%s\n%s",
		date, now.Format("2006-01-02"), diff, diffType, cmdLsString)
}

func (c *defaultApp) Hash() {
	timeDiffFn := util.SpendTimeDiff()
	vPath := c.Cc.SubCommand
	if vPath == "" {
		lgr.Error("请指定路径文件或目录")
		return
	}

	fh := &FileHash{
		Vtype: c.Cc.ArgRaw("type", "t"),
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

	lgr.Info("文件读取(%s)成功，如列表下：\n%s\n", fh.Vtype, bin.FormatTable(tableData))
	lgr.Info("用时 %s", timeDiffFn())
}

func (c *defaultApp) Ganz() {
	year := c.Cc.SubCommand
	if year == "" {
		lgr.Info("请输入年份，来计算干支纪元法。默认为当年")
	}

	y, _ := strconv.Atoi(year)
	if y <= 0 {
		y = time.Now().Year()
	}

	gz, zod := ganz.CountGzAndZodiac(y)

	fmt.Printf("  %d年，干支纪元%s年，属%s.\n", y, gz, zod)
	fmt.Printf("\n天干：%s\n地支：%s\n属相：%s\n", ganz.TianGan, ganz.DiZhi, ganz.Zodiac)
}

func tryReadFileOrText(text string) (string, error, string) {
	fl, err := os.Open(text)
	if err != nil {
		return text, fmt.Errorf("文件读取错误，%v", err), ""
	}
	defer fl.Close()
	bys, err := io.ReadAll(fl)
	if err != nil {
		return text, fmt.Errorf("文件内容读取错误，%v", err), ""
	}

	mime := http.DetectContentType(bys)

	return string(bys), nil, fmt.Sprintf("data:%s;base64,", mime)
}

func (c *defaultApp) Base64() {
	isDec := c.Cc.CheckSetting("decode", "d")
	outOptions := []string{"out", "o"}
	outName := c.Cc.ArgRaw(outOptions...)
	if c.Cc.CheckSetting(outOptions...) {
		if isDec {
			outName = "dec.raw"
		} else {
			outName = "enc.base64"
		}
	}
	text := c.Cc.SubCommand
	var prefix string
	if text == "" {
		if isDec {
			lgr.Info("解密时请提供编码文本")
			return
		}
		var rd str.RandString
		text = "这是一个实例文本，" + rd.String(rand.Intn(35))
		lgr.Info("空文本，设置默认文本：%s", text)
	} else if c.Cc.CheckSetting("file", "f") {
		var err error
		text, err, _ = tryReadFileOrText(text)
		if err != nil {
			lgr.Error(err.Error())
			return
		}
	} else {
		text, _, prefix = tryReadFileOrText(text)
	}

	if isDec {
		by, err := base64.StdEncoding.DecodeString(text)
		if err != nil {
			lgr.Error("编码错误，%v", err)
			return
		}

		lgr.Info("解码成功：\n%s", string(by))
		return
	}

	uriContent := base64.StdEncoding.EncodeToString([]byte(text))
	uriContent = fmt.Sprintf("%s%s", prefix, uriContent)
	if outName != "" {
		err := fs.Put(outName, uriContent)
		if err != nil {
			lgr.Error("文件保存错误(%s)\n  %v", outName, uriContent)
			return
		}
		lgr.Info("已保存内容到(%s)", outName)
		return
	}
	lgr.Info("数据已编码！编码内容(URI)：\n%s\n", uriContent)
}

// Cal 等式计算
//
// 支持进制转换，Decimal/Binary/Octal/Hexadecimal
func (c *defaultApp) Cal() {
	equal := c.Cc.SubCommand
	if equal == "" {
		equal = butil.InputRequire("请输入等式符号：", nil)
	}
	// 清理 ',/_'
	equal = str.NumberClear(equal)

	var baseExp string
	toArgs := []string{"to", "t"}
	baseTsf := c.Cc.CheckSetting(toArgs...)
	if baseTsf {
		baseExp = c.Cc.ArgRaw(toArgs...)
		rslt, err := baseTransfer(equal, baseExp)
		if err != nil {
			lgr.Error("进制转换错误。\n  %v", err)
			return
		}
		lgr.Info("进制转换：%s", color.StyleByAnsi(color.AnsiTextGreen, rslt))
		return
	}

	spanReg := regexp.MustCompile(`\s+`)
	equal = spanReg.ReplaceAllString(equal, "")

	calc := str.NewCalc(equal)
	rslt := calc.Count()
	lgr.Info("输入等式：\n%s => %s， %s",
		color.StyleByAnsi(color.AnsiTextCyan, equal),
		color.StyleByAnsi(color.AnsiTextGreen, str.NumberSplitFormat(rslt)),
		color.StyleByAnsi(color.AnsiTextBlackBr, str.FloatSimple(calc.String())))
}
