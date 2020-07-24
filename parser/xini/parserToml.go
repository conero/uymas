package xini

// @Date：   2019/1/22 0022 23:45
// @Author:  Joshua Conero
// @Name:    名称描述

// Toml 文件解析
type TomlParser struct {
	BaseParser
}

//
func (p *TomlParser) MoreToDo() string {
	// @TODO 待实现
	return `THERE'RE MANY THINGS NEED TO DO YET.`
}

// 获取驱动名称
func (p TomlParser) Driver() string {
	return SupportNameRong
}
