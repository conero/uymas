package xini

import (
	"fmt"
	"gitee.com/conero/uymas/fs"
	"os"
	"path"
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
	baseDir string            // 读取文件所在目录
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

// 文件引入
func (p *baseFileParse) include(ln string, target *map[string]any) bool {
	ln = strings.TrimSpace(ln)
	isIclReg := getRegByKey("reg_include_smbl")
	if isIclReg == nil || !isIclReg.MatchString(ln) {
		return false
	}

	idx := strings.Index(ln, " ")
	if idx == -1 {
		return false
	}

	// 文件读取
	filename := strings.TrimSpace(ln[idx:])
	filepath := path.Join(p.baseDir, filename)
	fi, err := os.Stat(filepath)
	if err != nil {
		fi, err = os.Stat(filename)
		if err != nil {
			// 文件读取失败！
			return false
		}
		filepath = filename
	}

	// 为目录，则中断
	if fi.IsDir() {
		return true
	}

	bfp := &baseFileParse{
		data: *target,
	}
	bfp.read(filepath)
	target = &bfp.data

	return true
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
	p.baseDir = fs.StdPathName(path.Dir(fs.StdPathName(filename)))

	// 行扫描
	secTmpDd := map[string]any{}
	isSecMk := false
	var section string

	// 多行注释
	var isMutLineCmt = false
	var mt1 = IniParseSettings["mcomment1"]
	var mt2 = IniParseSettings["mcomment2"]

	// 多行字符串
	var isMutLineStr = false
	var mutLineStr = [2]string{"", ""}
	// 多行字符串结束处理，返回是否中断
	var mtsEndHandler = func() bool {
		mlsKey := mutLineStr[0]
		if mlsKey == "" {
			mutLineStr = [2]string{"", ""}
			return true
		}

		// 赋值
		if isSecMk {
			secTmpDd[mlsKey] = mutLineStr[1]
		} else {
			p.data[mlsKey] = mutLineStr[1]
		}

		mutLineStr = [2]string{"", ""}
		return true
	}

	ln.Scan(func(line string) {
		p.line += 1
		str := strings.TrimSpace(line)

		// 多行注释或字符串多行
		if str == mt1 || str == mt2 {
			if isMutLineStr { // 多行字符串结束
				isMutLineStr = false
				if mtsEndHandler() {
					return
				}

				return
			}
			if isMutLineCmt { // 结束
				isMutLineCmt = false
				return
			}

			isMutLineCmt = true
			return
		}
		if isMutLineStr {
			mtsIdxEnd := strings.Index(line, mt1)
			if mtsIdxEnd == -1 {
				mtsIdxEnd = strings.Index(line, mt2)
			}
			if mtsIdxEnd > -1 {
				isMutLineStr = false
				mutLineStr[1] += line[:mtsIdxEnd]
				// 保存字符串
				if mtsEndHandler() {
					return
				}
				return
			}
			mutLineStr[1] += line
			//fmt.Printf("mutLineStr: %#v\n", mutLineStr)
			return
		}
		if isMutLineCmt {
			//fmt.Printf("mt: %s\n", line)
			return
		}

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

		target := p.data
		if isSecMk {
			target = secTmpDd
		}
		// include
		if p.include(str, &target) {
			if isSecMk {
				secTmpDd = target
			} else {
				p.data = target
			}
			return
		}
		// 赋值
		idx := strings.Index(str, baseEqualToken)
		if idx == -1 { // 非法行自动过滤
			return
		}

		// 作用域， `{}`

		key := strings.TrimSpace(str[:idx])
		value := lnTrim(str[idx+1:])

		// 长字符串，`"""|'''`
		if strings.Index(value, mt1) == 0 || strings.Index(value, mt2) == 0 {
			isMutLineStr = true
			mutLineStr[0] = key

			mlsIdx := strings.Index(line, mt1)
			if mlsIdx == -1 {
				mlsIdx = strings.Index(line, mt2)
			}

			mutLineStr[1] = line[mlsIdx+len(mt1):]
			//fmt.Printf("mutLineStr: %#v\n", mutLineStr)
			return
		}

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
