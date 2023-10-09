package xini

import "strconv"

// @Date：   2018/8/19 0019 10:54
// @Author:  Joshua Conero
// @Name:    解析器

// Parser the ini file base parse interface
type Parser interface {
	Get(key string) (bool, any)
	GetDef(key string, def any) any
	HasKey(key string) bool
	// SetFunc 函数式值获取
	// 实现如动态值获取，类似 js 对象中的 [get function()]
	SetFunc(key string, regFn func() any) Parser

	// Raw 支持多级数据访问，获取元素数据
	// 实际读取的原始数据为 map[string]string
	Raw(key string) string

	// Value get or set value: key, value(nil), default
	Value(params ...any) any

	GetAllSection() []string
	// Section the param format support
	//		1.     fun Section(section, key string) 	二级访问
	//		2.     fun Section(format string) 			点操作
	Section(params ...any) any

	GetData() map[string]any

	Set(key string, value any) Parser // 设置值
	Del(key string) bool              // 删除键值

	IsValid() bool
	OpenFile(filename string) Parser
	ReadStr(content string) Parser
	ErrorMsg() string // 错误信息

	Save() bool
	SaveAsFile(filename string) bool
	Driver() string
}

// 解析为数字类型，i64/f64
func parseNumber(vStr string) (value any, isOk bool) {
	i64Symbol := getRegByKey("reg_i64_symbol")
	if i64Symbol != nil && i64Symbol.MatchString(vStr) {
		i64, er := strconv.ParseInt(vStr, 10, 10)
		if er == nil {
			value = i64
			isOk = true
			return
		}
	}

	f64Symbol := getRegByKey("reg_f64_symbol")
	if f64Symbol != nil && f64Symbol.MatchString(vStr) {
		f64, er := strconv.ParseFloat(vStr, 10)
		if er == nil {
			value = f64
			isOk = true
			return
		}
	}
	return
}

// 字符串清理
func stringClear(vStr string) string {
	strSymbol := getRegByKey("reg_str_symbol")
	if strSymbol != nil && strSymbol.MatchString(vStr) {
		vStr = vStr[1 : len(vStr)-1]
	}
	return vStr
}

// 将字符串解析为参数
// 将原始的字符串解析为对应的参数
func parseValue(vStr string) any {
	var value any
	switch vStr {
	case "true", "TRUE":
		value = true
	case "false", "FALSE":
		value = false
	default:
		// 包裹找字符串如，`"string"` 或 `'string'`
		if v, isOk := parseNumber(vStr); isOk {
			value = v
		} else {
			value = stringClear(vStr)
		}
	}
	return value
}
