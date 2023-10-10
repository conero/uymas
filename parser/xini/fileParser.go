package xini

import (
	"fmt"
	"regexp"
	"strings"
)

// @Date：   2018/8/19 0019 10:57
// @Author:  Joshua Conero
// @Name:    文件解析器

// FileParser File parser
type FileParser interface {
	Line() int // 获取总行数
	GetData() map[string]any
}

// base 文件解析
type baseFileParse struct {
	line    int               // 总行数
	comment int               // 注释行
	equal   int               // 等式行
	data    map[string]any    // 解析以后的数据
	rawData map[string]string // 原始数据
	section []string          // 节
	err     error             // 错误信息
}

// 变量解析支持
func (p *baseFileParse) supportVariable(s string) string {
	reg := getRegByKey("reg_var_support")
	if reg == nil || !reg.MatchString(s) {
		return s
	}

	// 变量
	regVal := getRegByKey("reg_var_support_val")
	if regVal != nil {
		// 变量
		for _, vl := range regVal.FindAllString(s, -1) {
			name := strings.TrimSpace(vl[1:])
			vAny, exist := p.data[name]
			rpl := ""
			if exist && vAny != nil {
				rpl = fmt.Sprintf("%v", vAny)
			}
			s = strings.ReplaceAll(s, vl, rpl)
		}
	}

	// 变量引用
	regRef := getRegByKey("reg_var_support_ref")
	if regRef != nil {
		// 变量
		for _, vl := range regRef.FindAllString(s, -1) {
			name := strings.TrimSpace(vl[1:])
			vAny, exist := p.data[name]
			rpl := ""
			if exist && vAny != nil {
				rpl = fmt.Sprintf("%v", vAny)
			}
			s = strings.ReplaceAll(s, vl, rpl)
		}
	}
	return s
}

// 文件读取
func (p *baseFileParse) read(filename string) *baseFileParse {
	if p.data == nil {
		p.data = map[string]any{}
	}

	if p.rawData == nil {
		p.rawData = map[string]string{}
	}

	if p.section == nil {
		p.section = []string{}
	}
	ln := NewLnRer(filename)
	// 行扫描
	secTmpDd := map[string]any{}
	isSecMk := false
	var section string
	ln.Scan(func(line string) {
		p.line += 1
		str := strings.TrimSpace(line)
		// 空行过滤
		if "" == str {
			p.comment += 1
			return
		}
		// 注释过滤
		if matched, _ := regexp.MatchString(baseCommentReg, str); matched {
			return
		}
		// 节处理
		if matched, _ := regexp.MatchString(baseSectionReg, str); matched {
			// section 加到 data 中
			if isSecMk {
				p.data[baseSecRegPref+section] = secTmpDd
			}

			// 值重置
			secTmpDd = map[string]any{}
			isSecMk = true
			section = str[1 : len(str)-1]
			p.section = append(p.section, section)

			return
		}
		// 赋值
		idx := strings.Index(str, baseEqualToken)
		key := strings.TrimSpace(str[:idx])
		value := lnTrim(str[idx+1:])
		// 变量命令替换
		value = p.supportVariable(value)
		// 赋值
		if isSecMk {
			secTmpDd[key] = parseValue(value)
		} else {
			p.data[key] = parseValue(value)
		}

		p.rawData[key] = value
	})

	// section 加到 data 中
	if isSecMk {
		p.data[baseSecRegPref+section] = secTmpDd
	}

	// 错误信息
	p.err = ln.error

	return p
}

// Line get the file line count
func (p *baseFileParse) Line() int {
	return p.line
}

// GetData get all data by parse file.
func (p *baseFileParse) GetData() map[string]any {
	return p.data
}
