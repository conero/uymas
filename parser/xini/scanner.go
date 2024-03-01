package xini

import (
	"crypto/sha512"
	"fmt"
	"gitee.com/conero/uymas/v2/fs"
	"gitee.com/conero/uymas/v2/number"
	"gitee.com/conero/uymas/v2/util/rock"
	"os"
	"path"
	"regexp"
	"strings"
)

type ScannerLog struct {
	Filename   string
	Size       int64
	IsOk       bool
	Line       int
	Hash       string
	ParentHash string
	Err        error
}

// Scanner File progressive scanner
type Scanner struct {
	data    map[string]any
	records []ScannerLog
	mySelf  *ScannerLog
	baseDir string

	section    []string
	filename   string
	rtIsCmt    bool           // 运行时-当前多行注释
	rtIsScp    bool           // 运行时-当前作用域
	rtIsSec    bool           // 运行时-当前节
	rtIsMtl    bool           // 运行时-当前跨行
	rtMtlKey   string         // 运行时-当前跨行键值
	rtMtlLine  []string       // 运行时-跨行字符串
	rtSecKey   string         // 运行时-当前节名
	rtSecDd    map[string]any // 节数据作用域
	rtScpKey   string         // 运行时-作用域数据当前键值
	rtScpMap   map[string]any // 作用域数据-map
	rtScpStrLs []string       // 作用域数据-[]string

}

func NewScanner(args ...string) *Scanner {
	return &Scanner{
		filename: rock.ExtractParam("", args...),
	}
}

func (c *Scanner) init() {
	c.mySelf = &ScannerLog{
		Filename: c.filename,
		Hash:     getHash(c.filename),
		Line:     0,
	}
	c.baseDir = fs.StdPathName(path.Dir(fs.StdPathName(c.filename)))
	if c.data == nil {
		c.data = map[string]any{}
	}
}

func (c *Scanner) resetData() {
	c.init()
	c.records = nil
	c.data = map[string]any{}
	c.section = nil
}

// Detect whether it is a single line comment, multi line comment, etc
func (c *Scanner) shouldSkip(line string) (isSkip bool, ln string) {
	var cmt = [2]string{IniParseSettings["mcomment1"], IniParseSettings["mcomment2"]}
	hdlLn := strings.TrimSpace(line)

	// 多行注释结束
	rtIsCmt := hdlLn == cmt[0] || hdlLn == cmt[1]
	if c.rtIsCmt {
		// 多行注释结束
		if rtIsCmt {
			c.rtIsCmt = false
		}

		return true, ""
	} else if rtIsCmt {
		c.rtIsCmt = true
		return true, ""
	}

	// 单行注释过滤
	if matched, _ := regexp.MatchString(baseCommentReg, hdlLn); matched {
		return true, ""
	}

	return hdlLn == "", hdlLn
}

func (c *Scanner) parseMtl(kv KvPairs) (isSkip bool) {
	var cmt = [2]string{IniParseSettings["mcomment1"], IniParseSettings["mcomment2"]}
	rtIsMtl := strings.Index(kv.value, cmt[0]) == 0 || strings.Index(kv.value, cmt[1]) == 0
	if !rtIsMtl {
		return
	}

	c.rtMtlKey = kv.key
	c.rtIsMtl = true
	c.rtMtlLine = append(c.rtMtlLine, kv.value[len(cmt[0]):]+"\n")
	return true
}

// save 保存
func (c *Scanner) parseMtlSave(raw string) (isSkip bool) {
	if c.rtIsMtl {
		var cmt = [2]string{IniParseSettings["mcomment1"], IniParseSettings["mcomment2"]}
		idx := strings.LastIndex(raw, cmt[0])
		sLen := len(raw)
		isEndMtl := false
		if idx > -1 && sLen-idx == 4 {
			isEndMtl = true
			raw = raw[:idx]
		}

		if !isEndMtl {
			idx = strings.LastIndex(raw, cmt[1])
			if idx > -1 && sLen-idx == 4 {
				isEndMtl = true
				raw = raw[:idx]
			}
		}

		// 多行结束处理
		if isEndMtl {
			c.rtMtlLine = append(c.rtMtlLine, raw)
			c.saveData(c.rtMtlKey, strings.Join(c.rtMtlLine, ""))

			// 重置
			c.rtIsMtl = false
			c.rtMtlLine = nil
			c.rtMtlKey = ""
			return true
		}
	}

	if c.rtIsMtl {
		c.rtMtlLine = append(c.rtMtlLine, raw)
		return true
	}

	return
}

// parse Kv
func (c *Scanner) parseKv(hdlLn string) {
	kv := DecKvPairs(hdlLn)
	if c.parseScope(kv) {
		return
	}
	if kv.key == "" {
		return
	}
	if c.parseMtl(*kv) {
		return
	}

	value := c.supportVariable(kv.value)
	vAny := parseValue(value)
	c.saveData(kv.key, vAny)
	return
}

func (c *Scanner) saveData(key string, value any) {
	// 将数据保存到节中
	if c.rtIsSec {
		if c.rtSecDd == nil {
			c.rtSecDd = map[string]any{}
		}
		c.rtSecDd[key] = value
		return
	}
	c.data[key] = value
}

func (c *Scanner) parseScope(kv *KvPairs) (isSkip bool) {
	var scopeSml = [2]string{IniParseSettings["scope1"], IniParseSettings["scope2"]}

	// 作用域开始
	isStartIdx := strings.Index(kv.value, scopeSml[0])
	if !c.rtIsScp && (kv.value == scopeSml[0] || isStartIdx == 0) {
		c.rtIsScp = true
		c.rtScpKey = kv.key

		// 解析 `{xxxxx`
		if isStartIdx == 0 {
			vs := strings.TrimSpace(kv.value[1:])
			if vs != "" {
				ckv := DecKvPairs(vs)
				if ckv.isKv {
					c.rtScpMap[ckv.key] = parseValue(vs)
				} else {
					c.rtScpStrLs = append(c.rtScpStrLs, vs)
				}
			}
		}
		return true
	} else if !c.rtIsScp {
		return false
	}

	// 作用域结束
	if kv.value == scopeSml[1] {
		var value any
		if c.rtScpMap != nil {
			value = c.rtScpMap
		} else {
			value = c.rtScpStrLs
		}
		c.saveData(c.rtScpKey, value)

		// 保存后作用域清理
		c.rtScpKey = ""
		c.rtScpStrLs = nil
		c.rtScpMap = nil
		c.rtIsScp = false
		return true
	}

	// 作用域保存
	if kv.isString {
		c.rtScpStrLs = append(c.rtScpStrLs, kv.value)
	} else {
		if c.rtScpMap == nil {
			c.rtScpMap = map[string]any{}
		}
		c.rtScpMap[kv.key] = parseValue(kv.value)
	}
	return true
}

func (c *Scanner) parseSection(hdlLn string) (isSkip bool) {
	// 节处理
	if matched, _ := regexp.MatchString(baseSectionReg, hdlLn); matched {
		// section 加到 data 中
		if c.rtIsSec {
			c.rtIsSec = false
			c.saveData(baseSecRegPref+c.rtSecKey, c.rtSecDd)
		}

		// 值重置
		c.rtSecDd = map[string]any{}
		c.rtIsSec = true
		c.rtSecKey = hdlLn[1 : len(hdlLn)-1]
		c.section = append(c.section, c.rtSecKey)
		return true
	}
	return
}

func (c *Scanner) handler(line string) {
	c.mySelf.Line += 1

	// 多行解析保存
	if c.parseMtlSave(line) {
		return
	}

	// 非法字符处理
	isSkip, hdlLn := c.shouldSkip(line)
	if isSkip {
		return
	}

	// 行字符串处理
	hdlLn = lnTrim(hdlLn)

	// 解析节
	if c.parseSection(hdlLn) {
		return
	}

	// include 命令解析
	if c.parseInclude(hdlLn) {
		return
	}

	c.parseKv(hdlLn)
}

// include 命令处理
func (c *Scanner) parseInclude(hdlLn string) (isSkip bool) {
	target := c.data
	if c.rtIsSec {
		target = c.rtSecDd
	}
	loadOk, isInc := c.cmdInclude(hdlLn, &target)
	if !isInc {
		return
	}

	if loadOk {
		if c.rtIsSec {
			c.rtSecDd = target
		} else {
			c.data = target
		}
	}

	return true
}

// Scan start to scan file
func (c *Scanner) Scan(args ...string) error {
	// 新输入文件时才进行重置操作，支持重复处理
	flName := rock.ExtractParam("", args...)
	if flName != "" && c.filename != "" && c.filename != flName {
		c.filename = flName
		c.resetData()
	} else {
		c.init()
	}

	lr := NewLnRer(c.filename)
	fi, err := lr.ScanWithFlInfo(c.handler)

	if fi != nil {
		c.mySelf.IsOk = true
		c.mySelf.Size = fi.Size()
	}

	c.mySelf.Err = err
	c.records = append(c.records, *c.mySelf)
	return err
}

func (c *Scanner) Record() []ScannerLog {
	return c.records
}

// 文件引入
// @todo 应该记载载入了那些文件，以及对应文件的大小等信息（可选）
func (c *Scanner) cmdInclude(ln string, target *map[string]any) (loadOk bool, isIcl bool) {
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
	filepath := path.Join(c.baseDir, filename)

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
			ldMk := c.loadFile(childFp, target)
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

	loadOk = c.loadFile(filepath, target)
	return
}

func (c *Scanner) loadFile(flPath string, target *map[string]any) bool {
	scanner := &Scanner{
		filename: flPath,
		data:     *target,
	}

	err := scanner.Scan()

	slg := scanner.mySelf
	slg.ParentHash = c.mySelf.Hash
	c.records = append(c.records, *slg)
	target = &scanner.data

	return err == nil
}

// 执行变量解析
func (c *Scanner) supportVariable(s string) string {
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
			vAny, exist := c.data[name]
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
			vAny, exist := c.data[name]
			rpl := ""
			if exist && vAny != nil {
				rpl = fmt.Sprintf("%v", vAny)
			}
			s = strings.ReplaceAll(s, vl, rpl)
		}
	}
	return s
}

func getHash(s string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(s)))
}

// FileLog pring file reload info
func FileLog(arr []ScannerLog) string {
	if len(arr) == 0 {
		return "无文件读取记录"
	}

	var dick = map[string]ScannerLog{}
	var parent []string
	for _, slg := range arr {
		dick[slg.Hash] = slg
		if slg.ParentHash == "" {
			parent = append(parent, slg.Hash)
		}
	}

	var info []string
	for _, pid := range parent {
		my := dick[pid]
		info = append(info, fmt.Sprintf("Line %v: %v，成功-%v. %s", my.Line, number.Bytes(my.Size), my.IsOk, my.Filename))
		for _, clg := range arr {
			if clg.ParentHash == pid {
				info = append(info, fmt.Sprintf("  --> Line %v: %v，成功-%v. \n    %s", clg.Line, number.Bytes(clg.Size), clg.IsOk, clg.Filename))
			}
		}
	}

	return strings.Join(info, "\n")
}
