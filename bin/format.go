package bin

import (
	"fmt"
	"github.com/conero/uymas/number"
	"github.com/conero/uymas/str"
	"strconv"
	"strings"
)

/**
 * @DATE        2019/6/5
 * @NAME        Joshua Conero
 * @DESCRIPIT   命令行输出格式化
 **/

// 获取字符串格式化
// [[k,  v]]
// Deprecated: Use FormatKv instead.
func FormatStr(d string, ss ...[][]string) string {
	if d == "" {
		// 4 个空格
		d = "   "
	}
	bit := d[0:1]

	// 或者最大长度
	maxLen := 0
	for _, sg := range ss {
		for _, s := range sg {
			kLen := len(s[0])
			if kLen > maxLen {
				maxLen = kLen
			}
		}
	}

	maxLen += len(d)

	// 格式化
	var contents string
	for _, sg := range ss {
		for _, s := range sg {
			ss1 := s[0] + strings.Repeat(bit, maxLen-len(s[0])) + s[1] + "\n"
			contents += ss1
		}
	}

	return contents
}

// The `k-v` data format to beautiful str.
//
// FormatKv(kv map[string]interface{}, pref string)				含前缀的字符输出.
// FormatKv(kv map[string]interface{}, pref string, md string)	含前缀和中间连接符号的字符输出.
func FormatKv(kv map[string]interface{}, params ...string) string {
	var s, pref, d = "", "", ""
	var pLen = len(params)
	if pLen > 0 {
		pref = params[0]
	}
	if pLen > 1 {
		d = params[1]
	}

	// 计算最大长度
	// 最大长度
	maxLen := len(pref)
	for k, _ := range kv {
		kLen := len(k)
		if kLen > maxLen {
			maxLen = kLen
		}
	}

	if d == "" {
		// 4 个空格
		d = "   "
	}
	bit := d[0:1]
	maxLen += len(d)

	// 格式化
	for k, v := range kv {
		if s != "" {
			s += "\n"
		}
		s += pref + k + strings.Repeat(bit, maxLen-len(k)) + fmt.Sprintf("%v", v)
	}
	return s
}

// 格式化数组字符
// 用于命令行输出
// prefs 为 "" 时默认以数组索引开头；否则默给定的输出
func FormatQue(que []interface{}, prefs ...string) string {
	pref := ""  // 开头符号
	dter := " " // 空格
	if prefs != nil && len(prefs) > 0 {
		pref = prefs[0]
		if len(prefs) > 1 {
			dter = prefs[1]
		}
	}
	s := ""
	queLen := len(que)
	mdLen := 4 + len(strconv.Itoa(queLen))
	for i, q := range que {
		if pref == "" {
			iStr := strconv.Itoa(i) + "."
			s += iStr + strings.Repeat(dter, mdLen-len(iStr)) + fmt.Sprintf(" %v\n", q)
		} else {
			s += pref + strings.Repeat(dter, mdLen-len(pref)) + fmt.Sprintf(" %v\n", q)
		}
	}
	return s
}

// Bug(FormatQue): 中文长度无法使字符串字符对齐

// 表格格式化
// (data, bool) 是否使用 idx
func FormatTable(data [][]interface{}, args ...interface{}) string {
	useIdxMk := true
	if args != nil {
		if v, isBool := args[0].(bool); isBool {
			useIdxMk = v
		}
	}

	// 数据处理
	data2Str := [][]string{}
	maxLenQue := []int{}
	for _, dd := range data {
		ddStr := []string{}
		for i, d := range dd {
			vStr := fmt.Sprintf("%v", d)
			ddStr = append(ddStr, vStr)
			ddStrLen := len(vStr)
			if len(maxLenQue) > i {
				if maxLenQue[i] < ddStrLen {
					maxLenQue[i] = ddStrLen
				}
			} else {
				maxLenQue = append(maxLenQue, ddStrLen)
			}
		}
		data2Str = append(data2Str, ddStr)
	}

	var s string
	dCtt := len(data)
	maxLen := number.SumQInt(maxLenQue) + dCtt*2
	if useIdxMk {
		dCttLen := len(strconv.Itoa(dCtt) + ".")
		maxLen += dCttLen + dCtt*2
		maxLenQue = append([]int{dCttLen}, maxLenQue...)
	} else {
		maxLen += (dCtt - 1) * 2
	}

	for j, sdd := range data2Str {
		line := ""
		tLen := maxLen
		if useIdxMk {
			jStr := strconv.Itoa(j + 1)
			tLen -= tLen
			jStr = str.PadRight(jStr, " ", maxLenQue[0]+2)
			s += jStr
		}
		for i, sd := range sdd {
			maxCol := maxLenQue[i]
			if useIdxMk {
				maxCol = maxLenQue[i+1]
			}
			s += str.PadRight(sd, " ", maxCol+2)
		}
		s += line + "\n"
	}
	return s
}
