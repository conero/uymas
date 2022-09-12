package xini

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

	GetData() map[any]any

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
