package xini

// @Date：   2019/1/22 0022 23:45
// @Author:  Joshua Conero
// @Name:    名称描述

// TomlParser Toml file syntax parse
// @todo need todo
type TomlParser struct {
	BaseParser
}

func (p *TomlParser) MoreToDo() string {
	// @TODO 待实现
	return `THERE'RE MANY THINGS NEED TO DO YET.`
}

func (p TomlParser) Driver() string {
	return SupportNameRong
}
