package str

import (
	"regexp"
	"strconv"
	"strings"
)

// ParseUnicode parse unicode like `\u00001` to real char
func ParseUnicode(s string) string {
	reg := regexp.MustCompile(`(?i)\\u[\da-f]+`)
	for _, fv := range reg.FindAllString(s, -1) {
		rpl := strings.ReplaceAll(strings.ToUpper(fv), "\\U", "")
		v, err := strconv.ParseInt(rpl, 16, 64)
		if err != nil {
			continue
		}

		s = strings.ReplaceAll(s, fv, string(rune(v)))
	}
	return s
}
