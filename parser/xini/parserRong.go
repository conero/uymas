package xini

// @Date：   2018/8/19 0019 14:49
// @Author:  Joshua Conero
// @Name:    rong 解析器

// Rong 解析器，继承与: BaseParser
type RongParser struct {
	BaseParser
}

// 待完成
func (p *RongParser) MoreToDo() string {
	// @TODO 待实现
	return `THERE'RE MANY THINGS NEED TO DO YET.`
}

// 获取驱动名称
func (p RongParser) Driver() string {
	return SupportNameRong
}
