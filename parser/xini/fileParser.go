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
// @todo 待删除结构体
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

func (p *baseFileParse) loadFile(flPath string, target *map[string]any) bool {
	bfp := &baseFileParse{
		data: *target,
	}
	bfp.read(flPath)
	target = &bfp.data

	return true
}

// 文件引入
// @todo 应该记载载入了那些文件，以及对应文件的大小等信息（可选）
func (p *baseFileParse) include(ln string, target *map[string]any) (loadOk bool, isIcl bool) {
	ln = strings.TrimSpace(ln)
	isIclReg := getRegByKey("reg_include_smbl")
	if isIclReg == nil || !isIclReg.MatchString(ln) {
		return
	}

	isIcl = true // 是否为include标签

	idx := strings.Index(ln, " ")
	if idx == -1 {
		return
	}

	// 文件读取
	filename := strings.TrimSpace(ln[idx:])
	filepath := path.Join(p.baseDir, filename)

	// 支持目录
	dirIdx := strings.Index(filepath, "*")
	if dirIdx > 0 {
		vDir := filepath[:dirIdx]
		entrys, err := os.ReadDir(vDir)
		if err != nil {
			return
		}

		likeName := filepath[dirIdx:]
		regStr := strings.ReplaceAll(likeName, ".", `\.`)
		regStr = strings.ReplaceAll(likeName, "*", `.*`)
		regStr = fmt.Sprintf(`^%s$`, regStr)
		reg, regEr := regexp.Compile(regStr)
		if regEr != nil {
			return
		}
		childLoadOk := false
		for _, entry := range entrys {
			if entry.IsDir() {
				continue
			}
			if !reg.MatchString(entry.Name()) {
				continue
			}

			childFp := path.Join(vDir, entry.Name())
			ldMk := p.loadFile(childFp, target)
			if !childLoadOk && ldMk {
				childLoadOk = true
			}
		}

		loadOk = childLoadOk
		return

	}

	fi, err := os.Stat(filepath)
	if err != nil {
		fi, err = os.Stat(filename)
		if err != nil {
			// 文件读取失败！
			return
		}
		filepath = filename
	}

	// 为目录，则中断
	if fi.IsDir() {
		return
	}

	loadOk = p.loadFile(filepath, target)
	return
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

	// 作用域解析， “{}”
	var isScope = false
	var scopeKey string
	var scopeValue map[string]any
	var scopeValStrSlice []string
	// 作用域符号
	var scopeSml = [2]string{IniParseSettings["scope1"], IniParseSettings["scope2"]}
	var handlerScope = func() {
		if scopeKey == "" {
			return
		}
		var vAny any
		if len(scopeValue) > 0 {
			vAny = scopeValue
		} else if len(scopeValStrSlice) > 0 {
			vAny = scopeValStrSlice
		}

		if isSecMk {
			secTmpDd[scopeKey] = vAny
		} else {
			p.data[scopeKey] = vAny
		}

		// 数据复原
		scopeKey = ""
		scopeValStrSlice = []string{}
		scopeValue = map[string]any{}
	}

	_, err := ln.ScanWithFlInfo(func(line string) {
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
			return
		}
		if isMutLineCmt {
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

		// 作用域结束
		if str == scopeSml[1] {
			handlerScope()
			isScope = false
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
		loadMk, isIcl := p.include(str, &target)
		if loadMk {
			if isSecMk {
				secTmpDd = target
			} else {
				p.data = target
			}
			return
		} else if isIcl {
			return
		}
		// kv-键值对解析
		dkvp := DecKvPairs(str)
		if dkvp.isString { // 此行为字符串
			if isScope {
				scopeValStrSlice = append(scopeValStrSlice, dkvp.value)
				return
			}
		}
		if !dkvp.isKv {
			// 非法行自动过滤，不是有效的等式时
			return
		}

		// 长字符串，`"""|'''`
		if strings.Index(dkvp.value, mt1) == 0 || strings.Index(dkvp.value, mt2) == 0 {
			isMutLineStr = true
			mutLineStr[0] = dkvp.key

			mlsIdx := strings.Index(line, mt1)
			if mlsIdx == -1 {
				mlsIdx = strings.Index(line, mt2)
			}

			mutLineStr[1] = line[mlsIdx+len(mt1):]
			return
		}

		// 作用域， `{}`
		if strings.Index(dkvp.value, scopeSml[0]) == 0 {
			isScope = true
			scopeValue = map[string]any{}
			scopeKey = dkvp.key

			// 变量处理
			value := strings.TrimSpace(dkvp.value[1:])
			if value != "" {
				kp := DecKvPairs(value)
				if kp.isString {
					scopeValStrSlice = append(scopeValStrSlice, kp.value)
				} else if kp.isKv {
					scopeValue[kp.key] = parseValue(kp.value)
				}
			}

			return
		}

		// 变量命令替换
		value := p.supportVariable(dkvp.value)
		vAny := parseValue(value)
		// 赋值
		if isScope {
			scopeValue[dkvp.key] = vAny
		} else if isSecMk {
			secTmpDd[dkvp.key] = vAny
		} else {
			p.data[dkvp.key] = vAny
		}
		p.rawData[dkvp.key] = value
	})

	// section 加到 data 中
	if isSecMk {
		p.data[baseSecRegPref+section] = secTmpDd
	}

	// 错误信息
	p.err = err

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
