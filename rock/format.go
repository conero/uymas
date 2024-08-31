package rock

import (
	"fmt"
	"gitee.com/conero/uymas/v2/rock/constraints"
	"strings"
)

// FormatKv @todo neet to test
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
