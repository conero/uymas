package xini

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

// @Date：   2018/8/19 0019 14:25
// @Author:  Joshua Conero
// @Name:    基本 go 解析器
const (
	baseCommentReg = "^#|;"               // 注释符号
	baseSectionReg = "^\\[[^\\[\\]]+\\]$" // 节正则

	baseEqualToken = "="      // 等于符号
	baseSecRegPref = "__sec_" // 节前缀
)

// 基本/默认的解析器，支持标准的 ini 格式
type BaseParser struct {
	valid   bool
	section []string
	Container
	filename  string            // 文件名
	rawKvData map[string]string // 原始解析的参数
	errorMsg  string            // 错误信息
}

// 获取原始值，非解析后的
func (p *BaseParser) Raw(key string) string {
	var raw string
	if v, has := p.rawKvData[key]; has {
		raw = v
	}
	return raw
}

// 获取 ini 文件所有 “节”列表
func (p *BaseParser) GetAllSection() []string {
	return p.section
}

// 获取文件节
func (p *BaseParser) Section(params ...interface{}) interface{} {
	var value interface{}
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
			dd := data.(map[interface{}]interface{})
			if v, hasKey := dd[key]; hasKey {
				value = v
			}
		}
	}
	return value
}

// 设置解析器的值
func (p *BaseParser) Set(key string, value interface{}) Parser {
	p.GetData()
	p.Data[key] = value
	return p
}

// 删除键值
func (p *BaseParser) Del(key string) bool {
	return p.Container.Del(key)
}

// 函数式值获取
func (p *BaseParser) GetFunc(key string, regFn func() interface{}) Parser {
	p.Container.GetFunc(key, regFn)
	return p
}

// 判断解析器是否合法
func (p *BaseParser) IsValid() bool {
	return p.valid
}

// 打开文件并解析文件
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

// 解析字符串为参数
func (p *BaseParser) ReadStr(content string) Parser {
	return p
}

// 保存 ini 的值为文件
func (p *BaseParser) Save() bool {
	filename := p.filename
	return p.SaveAsFile(filename)
}

// 保存 ini 为文件
func (p *BaseParser) SaveAsFile(filename string) bool {
	successMk := true
	// 简单处理=字符串类型
	// @todo 需要做更多(20181105)
	iniTxt := "; power by (" + Name + "; V" + Version + "/" + Release + ")" +
		"\n;time: " + time.Now().String() +
		"\n; github.com/" + Name
	for k, v := range p.Data {
		switch k.(type) {
		case string:
			iniTxt += "\n" + k.(string) + "	= "
			if _, isStr := v.(string); isStr {
				iniTxt += v.(string)
			}
		}
	}
	// 0644 Append
	// 0755
	err := ioutil.WriteFile(filename, []byte(iniTxt), 0755)
	if err != nil {
		fmt.Println(err.Error())
		successMk = false
	}
	return successMk
}

//---------------------------- 来自 Container 对象的方法重写 -------------------------

// 只获取
func (p *BaseParser) Get(key string) (bool, interface{}) {
	return p.Container.Get(key)
}

// 带默认值得值获取
func (p *BaseParser) GetDef(key string, def interface{}) interface{} {
	return p.Container.GetDef(key, def)
}

// 带默认值得值获取
func (p *BaseParser) HasKey(key string) bool {
	return p.Container.HasKey(key)
}

// 错误错误信息
func (p *BaseParser) ErrorMsg() string {
	return p.errorMsg
}

// 当前项目获取驱动名称
func (p BaseParser) Driver() string {
	return SupportNameIni
}

// =>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>(BaseStrParse)>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

// baseStrParse
// 基本字符串解析器
type BaseStrParse struct {
	data map[interface{}]interface{}
	line int
}

// 字符串行数
func (p *BaseStrParse) Line() int {
	return p.line
}

// 获取所有数据
func (p *BaseStrParse) GetData() map[interface{}]interface{} {
	return p.data
}

// 加载字符串参数
func (p *BaseStrParse) LoadContent(content string) StrParser {
	p.data = map[interface{}]interface{}{}
	lineCtt := 0
	str2lines(content, func(line string) {
		lineCtt += 1
	})
	p.line = lineCtt
	return p
}
