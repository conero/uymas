package xini

// @Date：   2018/8/19 0019 10:54
// @Author:  Joshua Conero
// @Name:    解析器

// ini 文件基本的解析器接口
type Parser interface {
	// 读取参数
	Get(key string) (bool, interface{})
	GetDef(key string, def interface{}) interface{}
	HasKey(key string) bool
	// 函数式值获取
	// 实现如动态值获取，类似 js 对象中的 [get function()]
	GetFunc(key string, regFn func() interface{}) Parser

	// 支持多级数据访问，获取元素数据
	// 实际读取的原始数据为 map[string]string
	Raw(key string) string

	// 获取参数: key, value(nil), default
	Value(params ...interface{}) interface{}

	// 节处理
	GetAllSection() []string
	// 参数格式
	//		1.     fun Section(section, key string) 	二级访问
	//		2.     fun Section(format string) 			点操作
	Section(params ...interface{}) interface{}

	// 获取数据返回 nil
	GetData() map[interface{}]interface{}

	Set(key string, value interface{}) Parser // 设置值
	Del(key string) bool                      // 删除键值

	// 文件检测有效性
	IsValid() bool
	OpenFile(filename string) Parser
	ReadStr(content string) Parser
	ErrorMsg() string // 错误信息

	// 保存到文件
	Save() bool
	SaveAsFile(filename string) bool
	Driver() string
}
