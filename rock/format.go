package rock

import (
	"fmt"
	"gitee.com/conero/uymas/v2/rock/constraints"
	"strings"
)

// FormatKv Formatting string alignment, currently valid only for Latin characters, not for mixed characters.
//
// BUG: Mixed strings in Chinese and English (Latin) are invalid
func FormatKv[K constraints.KeyIterable, V any](data map[K]V) string {
	var maxLen = 0
	var stringMap = map[string]V{}
	for k, v := range data {
		kStr := fmt.Sprintf("%v", k)
		kLen := len(kStr)
		if kLen > maxLen {
			maxLen = kLen
		}
		stringMap[kStr] = v
	}

	var lines []string
	for k, v := range stringMap {
		line := fmt.Sprintf("%-"+fmt.Sprintf("%d", maxLen+4)+"s%v", k, v)
		lines = append(lines, line)
	}

	return strings.Join(lines, "\n")
}

func FormatList[V any](vList []V, joins ...string) string {
	join := Param("   ", joins...)
	vLen := len(vList)
	if vLen == 0 {
		return ""
	}

	maxLen := len(fmt.Sprintf("%v", vLen))
	vFmt := "%-" + fmt.Sprintf("%d", maxLen+1) + "s" + join + "%v"
	var queue []string
	for i, v := range vList {
		vs := fmt.Sprintf("%d.", i+1)
		queue = append(queue, fmt.Sprintf(vFmt, vs, v))
	}

	return strings.Join(queue, "\n")
}

// FormatTable two-dimensional array formatted output
func FormatTable[V any](table [][]V, seps ...int) string {
	var maxLenLs []int
	var data [][]string
	// calculate the maximum length value
	for _, vc := range table {
		var line []string
		for i, v := range vc {
			vStr := fmt.Sprintf("%v", v)
			maxLen := ListGetOr(maxLenLs, i, 0)
			vLen := len(vStr)
			if vLen > maxLen {
				maxLen = vLen
				lnLen := len(maxLenLs)
				if lnLen <= i {
					maxLenLs = append(maxLenLs, maxLen)
				} else {
					maxLenLs[i] = maxLen
				}
			}
			line = append(line, vStr)
		}
		data = append(data, line)
	}

	// Output value construction
	var outputLns []string
	numCount := len(maxLenLs)
	joinSep := Param(2, seps...)
	for _, sArr := range data {
		var lnStr string
		for i, sLn := range sArr {
			maxLen := maxLenLs[i]
			sep := 0
			if i+1 < numCount {
				sep = joinSep
			}
			vFmt := "%-" + fmt.Sprintf("%d", maxLen+sep) + "s"
			lnStr += fmt.Sprintf(vFmt, sLn)
		}
		outputLns = append(outputLns, lnStr)
	}
	return strings.Join(outputLns, "\n")
}
