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

// FormatKv The `k-v` data format to beautiful str.
//FormatKv(kv map[string]interface{}, pref string)				 provide pref param form FormatKv.
//FormatKv(kv map[string]interface{}, pref string, md string)	     provide pref and middle param form FormatK.
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

// FormatKvSort The `k-v` data format to beautiful str.
//FormatKvSort(kv map[string]interface{}, pref string)				 provide pref param form FormatKv.
//FormatKvSort(kv map[string]interface{}, pref string, md string)	     provide pref and middle param form FormatK.
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

// FormatQue format the string array, using for cli output pretty.
// where prefs is empty default use the array index
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

// FormatTable Table format output by slice:
// 	(table, bool) if is use the idx, table is 2 dimensional array.
// Bug(FormatQue): chinese text cannot alignment
func FormatTable(table interface{}, args ...interface{}) string {
	useIdxMk := true
	if args != nil {
		if v, isBool := args[0].(bool); isBool {
			useIdxMk = v
		}
	}

	rv := reflect.ValueOf(table)
	var vLen int
	//Only Support Array/Slice, other output itself.
	if rv.Kind() == reflect.Array || rv.Kind() == reflect.Slice {
		vLen = rv.Len()
	} else {
		return fmt.Sprintf("%v", table)
	}

	var data2Str [][]string
	var maxLenQue []int

	for i := 0; i < vLen; i++ {
		arr := rv.Index(i).Interface()
		rvD1 := reflect.ValueOf(arr)
		//Only Support Array/Slice, other output itself.
		var ddStr []string
		var vStr string
		if rvD1.Kind() == reflect.Array || rvD1.Kind() == reflect.Slice {
			vLenD1 := rvD1.Len()
			for j := 0; j < vLenD1; j++ {
				vD1 := rvD1.Index(j).Interface()
				if vD1 == nil {
					vD1 = ""
				}
				vStr = fmt.Sprintf("%v", vD1)
				ddStr = append(ddStr, vStr)
				ddStrLen := len(vStr)
				if len(maxLenQue) > j {
					if maxLenQue[j] < ddStrLen {
						maxLenQue[j] = ddStrLen
					}
				} else {
					maxLenQue = append(maxLenQue, ddStrLen)
				}
			}
		} else {
			if arr == nil {
				arr = ""
			}
			vStr = fmt.Sprintf("%v", arr)
			ddStr = append(ddStr, vStr)
		}
		data2Str = append(data2Str, ddStr)
	}

	var s string
	dCtt := vLen
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
