package xini

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// @Date：   2018/8/19 0019 14:25
// @Author:  Joshua Conero
// @Name:    基本 go 解析器
const (
	baseCommentReg = "^(#|;)"             // 注释符号
	baseSectionReg = "^\\[[^\\[\\]]+\\]$" // 节正则

	baseEqualToken   = "="      // 等于符号
	baseLimiterToken = ","      // 分隔符号
	baseSecRegPref   = "__sec_" // 节前缀
)

// BaseParser base and default file parse, support the standard ini configure file
type BaseParser struct {
	valid   bool
	section []string
	Container
	filename  string            // 文件名
	rawKvData map[string]string // 原始解析的参数
	errorMsg  string            // 错误信息
}

// Raw get the raw value that not parse to what the data by itself
func (p *BaseParser) Raw(key string) string {
	var raw string
	if v, has := p.rawKvData[key]; has {
		raw = v
	}
	return raw
}

func (p *BaseParser) GetAllSection() []string {
	return p.section
}

func (p *BaseParser) Section(params ...any) any {
	var value any
	var section, key string

	if nil == params {
		value = nil
	} else if len(params) == 1 {
		format := params[0].(string)
		if idx := strings.Index(format, "."); idx > -1 {
			section = format[0:idx]
			key = format[idx+1:]
		}
	} else if len(params) > -1 {
		section = params[0].(string)
		key = params[1].(string)
	}

	if section != "" && key != "" {
		if data, hasSection := p.Data[baseSecRegPref+section]; hasSection {
			dd := data.(map[any]any)
			if v, hasKey := dd[key]; hasKey {
				value = v
			}
		}
	}
	return value
}

func (p *BaseParser) Set(key string, value any) Parser {
	p.GetData()
	p.Data[key] = value
	return p
}

func (p *BaseParser) Del(key string) bool {
	return p.Container.Del(key)
}

func (p *BaseParser) SetFunc(key string, regFn func() any) Parser {
	p.Container.SetFunc(key, regFn)
	return p
}

func (p *BaseParser) IsValid() bool {
	return p.valid
}

func (p *BaseParser) OpenFile(filename string) Parser {
	reader := &baseFileParse{}
	reader.read(filename)
	p.Data = reader.GetData()
	p.filename = filename
	p.rawKvData = reader.rawData
	p.valid = true
	if reader.err != nil {
		p.errorMsg = reader.err.Error()
		p.valid = false
	}
	return p
}

func (p *BaseParser) ReadStr(content string) Parser {
	return p
}

func (p *BaseParser) Save() bool {
	filename := p.filename
	return p.SaveAsFile(filename)
}

func (p *BaseParser) SaveAsFile(filename string) bool {
	successMk := true
	// 简单处理=字符串类型
	// @todo 需要做更多(20181105)
	iniTxt := "; power by (" + Name + "; V" + Version + "/" + Release + ")" +
		"\n;time: " + time.Now().String() +
		"\n; github.com/" + Name
	for k, v := range p.Data {
		iniTxt += "\n" + k + "	= "
		if _, isStr := v.(string); isStr {
			iniTxt += v.(string)
		}
	}
	// 0644 Append
	// 0755
	err := os.WriteFile(filename, []byte(iniTxt), 0755)
	if err != nil {
		fmt.Println(err.Error())
		successMk = false
	}
	return successMk
}

//---------------------------- 来自 Container 对象的方法重写 -------------------------

func (p *BaseParser) Get(key string) (bool, any) {
	return p.Container.Get(key)
}

func (p *BaseParser) GetDef(key string, def any) any {
	return p.Container.GetDef(key, def)
}

func (p *BaseParser) HasKey(key string) bool {
	return p.Container.HasKey(key)
}

// ErrorMsg get the last error message
func (p *BaseParser) ErrorMsg() string {
	return p.errorMsg
}

// Driver the current reader driver type
func (p BaseParser) Driver() string {
	return SupportNameIni
}

// =>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>(BaseStrParse)>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

// BaseStrParse base string to parse using ini syntax
type BaseStrParse struct {
	data map[any]any
	line int
}

func (p *BaseStrParse) Line() int {
	return p.line
}

func (p *BaseStrParse) GetData() map[any]any {
	return p.data
}

func (p *BaseStrParse) LoadContent(content string) StrParser {
	p.data = map[any]any{}
	lineCtt := 0
	str2lines(content, func(line string) {
		lineCtt += 1
	})
	p.line = lineCtt
	return p
}
