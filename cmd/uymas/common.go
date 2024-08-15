package main

import (
	"errors"
	"fmt"
	"gitee.com/conero/uymas/util/rock"
	"regexp"
	"strconv"
	"strings"
)

// 命令行内部函数

// 但进制转换进制转换
func baseTransfer(exp, base string) (string, error) {
	exp = strings.ReplaceAll(exp, ",", "")
	req := regexp.MustCompile(`(?i)^(0[bodxh])?[\dabcdef]+(\.[\dabcdef])*`)
	if !req.MatchString(exp) {
		return "", errors.New("数值进制写法错误，支持0b、0o、0d、0x(h)等开头")
	}

	var valBase = "od"
	var vb = 10
	if exp[:1] == "0" {
		valBase = exp[:2]
		exp = strings.ReplaceAll(exp, valBase, "")
	}
	switch strings.ToUpper(valBase) {
	case "0B":
		vb = 2
	case "0O":
		vb = 8
	case "0D":
		vb = 10
	case "0X":
		vb = 16
	}

	vNum, err := strconv.ParseInt(exp, vb, 64)
	if err != nil {
		return "", fmt.Errorf("%d进制解析错误，%v", vb, err)
	}

	matchBase := strings.ToUpper(base)
	switch matchBase {
	case "0B", "B", "2": // 二进制
		return strconv.FormatInt(vNum, 2), nil
	case "0O", "O", "8": // 八进制
		return strconv.FormatInt(vNum, 8), nil
	case "0X", "X", "0H", "H", "16": // 八进制
		return strconv.FormatInt(vNum, 16), nil
	default:
		if matchBase != "" && rock.ListIndex([]string{"0D", "D", "10"}, matchBase) == -1 {
			return "", errors.New("进制规则仅支持如：0b/2、0o/8、0d/10、0x/16(或0h)，不区分大小写")
		}
		return strconv.FormatInt(vNum, 10), nil
	}
}
