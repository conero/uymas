package main

import (
	"fmt"
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/cli/ansi"
	"gitee.com/conero/uymas/v2/cli/chest"
	"gitee.com/conero/uymas/v2/cli/gen"
	"gitee.com/conero/uymas/v2/culture/pinyin"
	"gitee.com/conero/uymas/v2/logger/lgr"
	"gitee.com/conero/uymas/v2/rock"
	"gitee.com/conero/uymas/v2/str"
	"gitee.com/conero/uymas/v2/util/tm"
	"os"
	"strconv"
	"strings"
)

type pinyinOption struct {
	File      string   `cmd:"file,f help:将文件的内容作为输入以获取字典拼音"`
	IsUnicode bool     `cmd:"unicode,U help:将unicode码转为字符串实体"`
	Utf16     bool     `cmd:"utf16 help:同步将输入的汉字转或任何字符转\sUnicode\s代码"`
	Search    string   `cmd:"search,S help:根据输入的拼音进行搜索"`
	Number    bool     `cmd:"number,n help:输出数字型pinyin"`
	Alpha     bool     `cmd:"alpha,a help:输出不带音标拉丁文pinyin"`
	All       bool     `cmd:"all,A help:输出所有类型的拼音"`
	Seps      []string `cmd:"sep,S help:设置查询到的字符分割，默认为空"`
	Words     string   `cmd:"words isdata"`
	globalOption
}

func cmdPinyin(args cli.ArgsParser) {
	spendFn := tm.SpendFn()
	var opt pinyinOption
	err := gen.ArgsDress(args, &opt)
	if err != nil {
		lgr.Error(err.Error())
		return
	}

	words := opt.Words
	if opt.File != "" {
		bys, err := os.ReadFile(opt.File)
		if err != nil {
			lgr.Error("读取文件错误，%s", ansi.Style(err, ansi.Red))
			return
		}
		lgr.Info("已读取文件 %s 的内容", opt.File)
		words = string(bys)
	}

	if words == "" {
		words = chest.InputRequire("请输入中文汉字：", nil)
	}

	if opt.IsUnicode {
		lgr.Info("%s", str.Str(words).ParseUnicode())
		return
	}

	if opt.Utf16 {
		var codeList []string
		var strList []string
		for _, r := range []rune(words) {
			codeList = append(codeList, fmt.Sprintf("U+%s", strconv.FormatInt(int64(r), 16)))
			strList = append(strList, fmt.Sprintf("\\u%s", strconv.FormatInt(int64(r), 16)))
		}
		lgr.Info("%s 转utf16如：\n %s\n", words,
			ansi.Style(strings.Join(strList, ""), ansi.Green))

		if opt.IsVerbose {
			lgr.Info("%s 转utf16 (unicode 风格)如：\n %s\n", words,
				ansi.Style(strings.Join(codeList, " "), ansi.Green))
		}
	}

	pinyinCache := getPinyin()
	if opt.Search != "" {
		list := pinyinCache.SearchAlpha(opt.Search, 1000)
		textList := list.Text()
		//lgr.Info("搜索到字符如下：\n%s\n", bin.FormatQue(textList))
		lgr.Info("搜索到字符如下：\n%s\n", strings.Join(textList, " "))
		fmt.Printf("\n   搜索到 %d 个字，限制搜索字数字 %d，用时 %v\n", len(textList), 1000, spendFn())
		return
	}

	var line string
	if opt.Number {
		line = pinyinCache.SearchByGroup(words).Number(opt.Seps...)
	} else if opt.Alpha {
		line = pinyinCache.SearchByGroup(words).Alpha(opt.Seps...)
	} else if opt.All {
		vList := pinyinCache.SearchByGroup(words)
		line = "原始拼音：" + vList.Tone(opt.Seps...) + "\n" +
			"数字声调拼音：" + vList.Number(opt.Seps...) + "\n" +
			"字母拼音：" + vList.Alpha(opt.Seps...)

		// 多音字
		pyList := vList.Polyphony(pinyin.PinyinTone)
		pyCount := len(pyList)
		if len(pyList) > 0 {
			line += fmt.Sprintf("\n  多音字共 %d 组，详细如：\n原始拼音组：\n%s\n数字拼音组：\n%s\n字母拼音组：\n%s\n",
				pyCount, rock.FormatList(pyList, "  "),
				rock.FormatList(vList.Polyphony(pinyin.PinyinNumber), "  "),
				rock.FormatList(rock.ListNoRepeat(vList.Polyphony(pinyin.PinyinAlpha)), "  "))
		}
	} else {
		if opt.IsVerbose {
			list := pinyinCache.SearchByGroup(words)
			line = rock.FormatList(list.Polyphony(pinyin.PinyinTone, opt.Seps...), " ")
		} else {
			line = pinyinCache.SearchByGroup(words).Tone(opt.Seps...)
		}

	}

	fmt.Println(line)
	fmt.Printf("\n   字长%d（Unicode），用时 %v\n", len(words), spendFn())

}
