package xini

// @Date：   2018/8/19 0019 14:49
// @Author:  Joshua Conero
// @Name:    rong 解析器

// RongParser the Rong Parser inherited from BaseParser
type RongParser struct {
	BaseParser
}

// MoreToDo @need todo
func (p *RongParser) MoreToDo() string {
	// @TODO 待实现
	return `THERE'RE MANY THINGS NEED TO DO YET.`
}

func (p RongParser) Driver() string {
	return SupportNameRong
}
