package bin

import (
	"fmt"
	"gitee.com/conero/uymas/str"
	"strings"
)

// GetHelpEmbed GetHelpEmbed(content string, lang string)
func GetHelpEmbed(content string, args ...string) string {
	queue := strings.Split(content, "\n")
	argsLine := ""
	if len(args) > 0 {
		argsLine = strings.ToLower(args[0])
	}
	var (
		needStrQueue  []string
		isNeedMk      = false
		supportLang   []string
		supportLangMk bool
	)
	for _, line := range queue {
		ln := strings.TrimSpace(line)
		if ln == "" {
			continue
		}
		first := ln[:1]
		if first == ";" {
			continue
		}
		if first == ":" {
			var (
				matchStr string
				idx      int
			)
			if !supportLangMk {
				matchStr = ":lang-support="
				idx = strings.Index(ln, matchStr)
				if idx > -1 && argsLine != "" {
					support := strings.ToLower(ln[len(matchStr):])
					supportLang = strings.Split(support, ",")
					if str.InQuei(argsLine, supportLang) == -1 {
						panic(fmt.Sprintf("Lang un-support: %v, lang-list: %v",
							argsLine, strings.Join(supportLang, ",")))
					}
					supportLangMk = true
					continue
				}
			}
			matchStr = ":lang="
			idx = strings.Index(ln, matchStr)
			if idx > -1 {
				if isNeedMk {
					break
				}
				lang := strings.ToLower(ln[len(matchStr):])
				if lang == argsLine || argsLine == "" {
					isNeedMk = true
					continue
				}
			}
		}
		if isNeedMk {
			needStrQueue = append(needStrQueue, line)
		}
	}
	return strings.Join(needStrQueue, "\n")
}
