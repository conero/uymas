package xini

import (
	"regexp"
	"strings"
)

// @Date：   2018/8/19 0019 10:57
// @Author:  Joshua Conero
// @Name:    文件解析器

// FileParser File parser
type FileParser interface {
	Line() int // 获取总行数
	GetData() map[interface{}]interface{}
}

// base 文件解析
type baseFileParse struct {
	line    int                         // 总行数
	comment int                         // 注释行
	equal   int                         // 等式行
	data    map[interface{}]interface{} // 解析以后的数据
	rawData map[string]string           // 原始数据
	section []string                    // 节
	err     error                       // 错误信息
}

// 文件读取
func (p *baseFileParse) read(filename string) *baseFileParse {
	if p.data == nil {
		p.data = map[interface{}]interface{}{}
	}

	if p.rawData == nil {
		p.rawData = map[string]string{}
	}

	if p.section == nil {
		p.section = []string{}
	}
	ln := NewLnRer(filename)
	// 行扫描
	secTmpDd := map[interface{}]interface{}{}
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
			secTmpDd = map[interface{}]interface{}{}
			isSecMk = true
			section = str[1 : len(str)-1]
			p.section = append(p.section, section)

			return
		}
		// 赋值
		idx := strings.Index(str, baseEqualToken)
		key := strings.TrimSpace(str[:idx])
		value := strings.TrimSpace(str[idx+1:])
		var dd interface{}
		switch value {
		case "true":
			dd = true
		case "false":
			dd = false
		default:
			dd = value
		}
		// 赋值
		if isSecMk {
			secTmpDd[key] = dd
		} else {
			p.data[key] = dd
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
func (p *baseFileParse) GetData() map[interface{}]interface{} {
	return p.data
}
