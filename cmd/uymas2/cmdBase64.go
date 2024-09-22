package main

import (
	"encoding/base64"
	"fmt"
	"gitee.com/conero/uymas/v2/cli"
	"gitee.com/conero/uymas/v2/cli/gen"
	"gitee.com/conero/uymas/v2/logger/lgr"
	"gitee.com/conero/uymas/v2/str"
	"gitee.com/conero/uymas/v2/util/fs"
	"gitee.com/conero/uymas/v2/util/tm"
	"io"
	"math/rand"
	"net/http"
	"os"
)

type base64Option struct {
	IsDecode bool   `cmd:"decode,d help:读取文件或字符串的base64"`
	IsFile   bool   `cmd:"file,f help:是否为文件，不设置是自动检测。用于产生混淆时"`
	Output   string `cmd:"out,o help:保存内容到文件，未提供名称是默认为\\senc.base64\\s或\\sdec.raw"`
	Input    string `cmd:"data isdata"`
}

func cmdBase64(args cli.ArgsParser) {
	spendFn := tm.SpendFn()
	var opt base64Option
	err := gen.ArgsDress(args, &opt)
	if err != nil {
		lgr.Error(err.Error())
		return
	}

	isDec := opt.IsDecode
	outOptions := []string{"out", "o"}
	outName := opt.Output
	if args.Switch(outOptions...) {
		if isDec {
			outName = "dec.raw"
		} else {
			outName = "enc.base64"
		}
	}
	text := opt.Input
	var prefix string
	if text == "" {
		if isDec {
			lgr.Info("解密时请提供编码文本")
			return
		}
		var rd str.RandString
		text = "这是一个实例文本，" + rd.String(rand.Intn(35))
		lgr.Info("空文本，设置默认文本：%s", text)
	} else if opt.IsFile {
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
	lgr.Info("数据已编码！编码内容(URI)：\n%s\n\n用时：%v", uriContent, spendFn())
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
