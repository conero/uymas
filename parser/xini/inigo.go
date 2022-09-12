// Package xini parse ini config files and utils.
package xini

import (
	"regexp"
	"strconv"
	"strings"
)

// @Date：   2018/8/19 0019 10:58
// @Author:  Joshua Conero
// @Name:    库主文件

// NewParser instantiate the Parser
// param format(single param)
//
//	opts map[string]string{}|string
//		driver SupportNameRong SupportNameIni
//
// default BaseParser
func NewParser(params ...any) Parser {
	var driver string
	var opts map[string]any
	if params == nil {
		return new(BaseParser)
	} else {
		paramsLen := len(params)
		if driverTmp, isStr := params[0].(string); isStr {
			driver = driverTmp
		}

		if optsTmp, isOpt := params[0].(map[string]any); isOpt && driver == "" {
			opts = optsTmp
			if driverTmp, isset := opts["driver"]; isset {
				driver = driverTmp.(string)
			}
		}

		if paramsLen > 1 {
			if driverTmp, isStr := params[1].(string); isStr {
				driver = driverTmp
			}
		}
	}

	switch driver {
	case SupportNameRong:
		return new(RongParser)
	case SupportNameToml:
		return new(TomlParser)
	default:
		return new(BaseParser)
	}
}

// ParseValue parse the data value, the format like
//
//	bool:   true/false/TRUE/FALSE
//	string: '字符串', "字符串" 以及无法解析的参数
//	int64: 47, 52, -49552
//	float64: 3.14, -0.24552
//	null: nil     空值时，默认为 nil
func ParseValue(v string) any {
	var value any = nil
	v = strings.TrimSpace(v)
	isStr := false
	// 预处理，函数格式化字符
	// 删除，首尾字符
	if v != "" {
		tmpV := StrClear(v)
		if tmpV != v {
			v = tmpV
			isStr = true
		}
	}

	// 字符串
	if isStr {
		value = v
		v = ""
	}

	// 解析非空字符串
	if v != "" {
		lowStr := strings.ToLower(v)
		// 布尔对象
		if lowStr == "false" {
			value = false
		} else if lowStr == "true" {
			value = true
		} else {
			parseMk := false

			// 数字解析
			reg := regexp.MustCompile(`^[-]*[\d.]+$`)
			if reg.MatchString(v) {
				if strings.Index(v, ".") > -1 {
					// float64
					if f64, er := strconv.ParseFloat(v, 64); er == nil {
						value = f64
						parseMk = true
					}
				} else {
					if i64, er := strconv.ParseInt(v, 10, 64); er == nil {
						value = i64
						parseMk = true
					}
				}
			} else if strings.Index(v, ",") > -1 { // []int
				sQue := []string{}
				if matched, err := regexp.MatchString(`^[-\d]+[-\d,]+[-\d]+$`, v); err == nil && matched {
					sQue = strings.Split(v, ",")
					iQue := []int{}
					for _, v0 := range sQue {
						if i, err := strconv.Atoi(v0); err == nil {
							iQue = append(iQue, i)
						}
					}
					value = iQue
					parseMk = true
				} else if strings.Index(v, ".") > -1 {
					// []float64
					if matched, er := regexp.MatchString(`^[-\d.]+[-\d,.]{3,}[-\d,.]{3,}$`, v); er == nil && matched {
						sQue := strings.Split(v, ",")
						fQue := []float64{}
						for _, v0 := range sQue {
							if f64, er := strconv.ParseFloat(v0, 64); er == nil {
								fQue = append(fQue, f64)
							}
						}
						value = fQue
						parseMk = true
					}
				}

				if !parseMk {
					rpls := "_JC::JC_"
					vt0 := strings.Replace(v, "\\,", rpls, -1)
					reg2 := regexp.MustCompile(`['"]+[^'"]+['"]+`)

					for _, v1 := range reg2.FindAllString(vt0, -1) {
						v2 := strings.Replace(v1, ",", rpls, -1)
						vt0 = strings.Replace(vt0, v1, v2, -1)
					}

					if strings.Index(vt0, ",") > -1 {
						ss := []string{}
						for _, tV := range strings.Split(vt0, ",") {
							tV = StrClear(tV)
							tV = strings.ReplaceAll(tV, rpls, ",")
							ss = append(ss, tV)
						}
						value = ss
						parseMk = true
					}
				}
			}

			if !parseMk {
				value = v
			}
		}
	}
	return value
}

// StrClear the string data clear
func StrClear(s string) string {
	s = strings.TrimSpace(s)
	if s != "" {
		if matched, er := regexp.MatchString(`^'[^']*'$`, s); er == nil && matched {
			s = s[1 : len(s)-1]
		} else if matched, er := regexp.MatchString(`^"[^"]*"$`, s); er == nil && matched {
			s = s[1 : len(s)-1]
		}
	}
	return s
}
