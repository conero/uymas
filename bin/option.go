package bin

import (
	"fmt"
	"strings"
)

// @Date：   2019年11月11日 星期一
// @Author:  Joshua Conero
// @Name:    选项

type Option struct {
	Key         string // 键值
	Description string // 描述
	Logogram    string // 简写，用于实现端选项
}

// 获取简写
func (opt *Option) GetLogogram() string {
	if opt.Logogram == "" && opt.Key != "" {
		opt.Logogram = strings.ToUpper(opt.Key[:1])
	}
	return opt.Logogram
}

// 获取描述
func (opt *Option) GetDescrip() string {
	if opt.Description == "" {
		key := ""
		if opt.Key != "" {
			//key = fmt.Sprintf("--%v", opt.Key)
			key = "--" + opt.Key
		}
		logogram := opt.GetDescrip()
		if logogram != "" {
			logogram = fmt.Sprintf("-%v", logogram)
			if key != "" {
				logogram = ", " + logogram
			}
		}

		opt.Description = fmt.Sprintf("-%v, --%v set the system optional", key, logogram)
	}
	return opt.Description
}

// option 选项集合
type OptionSetting []Option
