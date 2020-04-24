package parser

import (
	"regexp"
	"strings"
)

//cli command 解析器
//解析规则：
//		单行以“;"分割
//		多行以“行\n分割"分割
func NewParser(script string) [][]string {
	var cmds [][]string
	var reg = regexp.MustCompile(`\n|;`)         //分行
	var regComment = regexp.MustCompile(`#.*$+`) //注释
	var regSpan = regexp.MustCompile(`\s{1,}`)
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
		tmpArr := []string{}
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
