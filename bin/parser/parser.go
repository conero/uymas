package parser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Transferred Escape symbol
var Transferred map[string]string = map[string]string{
	`\'`: "_Sg_.Qmark_", // Single quotation mark
	`\"`: "_Db_.Qmark_", // Double quotation marks
}

// NewParser cli command 解析器
// 解析规则：
//
//	单行以“;"分割
//	多行以“行\n分割"分割
func NewParser(script string) [][]string {
	var cmds [][]string
	var reg = regexp.MustCompile(`[\n;]`)        //分行
	var regComment = regexp.MustCompile(`#.*$+`) //注释
	var regSpan = regexp.MustCompile(`\s+`)
	//var regSign = regexp.MustCompile(`('[^\']*')|("[^\"]*")`)
	var strArr = reg.Split(script, -1)

	for _, sa := range strArr {
		sa = strings.TrimSpace(sa)
		sa = regComment.ReplaceAllString(sa, "")
		sa = strings.TrimSpace(sa)
		if sa == "" {
			continue
		}

		//
		//fmt.Println(sa)
		tmpRawArr := regSpan.Split(sa, -1)
		var tmpArr []string
		for _, t := range tmpRawArr {
			t = strings.TrimSpace(t)
			if t == "" {
				continue
			}
			tmpArr = append(tmpArr, t)
		}
		if len(tmpArr) > 0 {
			cmds = append(cmds, tmpArr)
		}
	}
	return cmds
}

// NewScriptFile parse the script file, the syntax like shell.
//
//	"#"   comment line
func NewScriptFile(filename string) []string {
	var cmds []string
	if fl, er := os.Open(filename); er == nil {
		buf := bufio.NewReader(fl)
		for {
			line, err2 := buf.ReadString('\n')
			line = strings.TrimSpace(line)
			// empty line
			if line == "" {
				continue
			}

			// comment line
			// # 1.case comment line
			// [command] --data "# is not comment" 	# 2.case comment line
			if line[:1] == "#" {
				continue
			}
			cmds = append(cmds, line)

			// 错误
			if err2 != nil {
				break
			}
		}
	}
	return cmds
}

// ParseLine parse shell line syntax to option
func ParseLine(line string) [][]string {
	var args [][]string
	line = strings.TrimSpace(line)
	//Transferred meaning
	for sym, tran := range Transferred {
		line = strings.ReplaceAll(line, sym, tran)
	}

	//temp part dist
	var tmpDick = map[string]string{}
	// `"string"`
	dbExp := regexp.MustCompile(`"[^"]+"`)
	handlerLn := line
	index := 0
	for _, sign := range dbExp.FindAllString(handlerLn, -1) {
		idxKey := fmt.Sprintf("_INDEX_%v_", index)
		handlerLn = strings.ReplaceAll(handlerLn, sign, idxKey)
		tmpDick[idxKey] = sign
		index += 1
	}

	// `'string'`
	dbExp = regexp.MustCompile(`'[^']+'`)
	for _, sign := range dbExp.FindAllString(handlerLn, -1) {
		idxKey := fmt.Sprintf("_INDEX_%v_", index)
		handlerLn = strings.ReplaceAll(handlerLn, sign, idxKey)
		tmpDick[idxKey] = sign
		index += 1
	}

	// blank
	dbExp = regexp.MustCompile(`\s{2,}`)
	handlerLn = dbExp.ReplaceAllString(handlerLn, " ")
	for _, ln := range strings.Split(handlerLn, ";") {
		cQueue := strings.Split(ln, " ")
		var nQueue []string
		for _, cWord := range cQueue {
			if dickVal, existDick := tmpDick[cWord]; existDick {
				cWord = dickVal
			}
			for dk, dv := range tmpDick {
				cWord = strings.ReplaceAll(cWord, dk, dv)
			}
			//Transferred meaning
			for sym, tran := range Transferred {
				cWord = strings.ReplaceAll(cWord, tran, sym)
			}
			nQueue = append(nQueue, cWord)
		}
		if len(nQueue) > 0 {
			args = append(args, nQueue)
		}
	}
	return args
}
