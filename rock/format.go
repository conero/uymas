package rock

import (
	"fmt"
	"gitee.com/conero/uymas/v2/rock/constraints"
	"strings"
)

// FormatKv @todo neet to do
func FormatKv[K constraints.KeyIterable, V any](data map[K]V) string {
	var s string
	return s
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
