// @DATE        2019/6/5
// @NAME        Joshua Conero

package bin

import (
	"fmt"
	"github.com/conero/uymas/number"
	"github.com/conero/uymas/str"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

/*
The `k-v` data format to beautiful str.
	FormatKv(kv map[string]interface{}, pref string)				 provide pref param form FormatKv.
	FormatKv(kv map[string]interface{}, pref string, md string)	     provide pref and middle param form FormatK.
*/
func FormatKv(kv interface{}, params ...string) string {
	var vf = reflect.ValueOf(kv)
	if vf.Kind() != reflect.Map {
		return ""
	}
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
	for mr := vf.MapRange(); mr.Next(); {
		kLen := len(fmt.Sprintf("%v", mr.Key()))
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
	for mr := vf.MapRange(); mr.Next(); {
		if s != "" {
			s += "\n"
		}
		k := fmt.Sprintf("%v", mr.Key())
		v := fmt.Sprintf("%v", mr.Value())
		s += pref + k + strings.Repeat(bit, maxLen-len(k)) + fmt.Sprintf("%v", v)
	}
	return s
}

/*
The `k-v` data format to beautiful str.
	FormatKvSort(kv map[string]interface{}, pref string)				 provide pref param form FormatKv.
	FormatKvSort(kv map[string]interface{}, pref string, md string)	     provide pref and middle param form FormatK.
*/
func FormatKvSort(kv interface{}, params ...string) string {
	var vf = reflect.ValueOf(kv)
	if vf.Kind() != reflect.Map {
		return ""
	}
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
	var sortKeys []string
	for mr := vf.MapRange(); mr.Next(); {
		k := fmt.Sprintf("%v", mr.Key())
		sortKeys = append(sortKeys, k)
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

	sort.Strings(sortKeys)
	// 格式化
	for mr := vf.MapRange(); mr.Next(); {
		k := fmt.Sprintf("%v", mr.Key())
		if s != "" {
			s += "\n"
		}
		s += pref + k + strings.Repeat(bit, maxLen-len(k)) + fmt.Sprintf("%v", mr.Value())
	}
	return s
}

// 格式化数组字符
// 用于命令行输出
// prefs 为 "" 时默认以数组索引开头；否则默给定的输出
func FormatQue(que interface{}, prefs ...string) string {
	pref := ""  // 开头符号
	dter := " " // 空格
	if prefs != nil && len(prefs) > 0 {
		pref = prefs[0]
		if len(prefs) > 1 {
			dter = prefs[1]
		}
	}
	s := ""
	vt := reflect.ValueOf(que)
	var queLen int
	//Only Support Array/Slice, other output itself.
	if vt.Kind() == reflect.Array || vt.Kind() == reflect.Slice {
		queLen = vt.Len()
	} else {
		return fmt.Sprintf("%v", que)
	}

	mdLen := 4 + len(strconv.Itoa(queLen))
	for i := 0; i < queLen; i++ {
		qVal := vt.Index(i).Interface()
		if pref == "" {
			iStr := strconv.Itoa(i) + "."
			s += iStr + strings.Repeat(dter, mdLen-len(iStr)) + fmt.Sprintf(" %v\n", qVal)
		} else {
			s += pref + strings.Repeat(dter, mdLen-len(pref)) + fmt.Sprintf(" %v\n", qVal)
		}
	}
	return s
}

// Bug(FormatQue): chinese text cannot alignment
//
// Table format output by slice:
// 	(data, bool) if is use the idx
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
