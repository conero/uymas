package bin

import (
	"fmt"
	"strings"
)

// @Date：   2019年11月11日 星期一
// @Author:  Joshua Conero
// @Name:    选项

type Option struct {
	Key         string      // 键值
	Description string      // 描述
	Logogram    string      // 简写，用于实现端选项
	Value       interface{} // 参数值
	RawValue    string      // 原始数据
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
			key = fmt.Sprintf("--%v", opt.Key)
		}
		logogram := opt.GetLogogram()
		if logogram != "" {
			logogram = fmt.Sprintf("-%v", logogram)
			if key != "" {
				logogram = ", " + logogram
			}
		}

		opt.Description = fmt.Sprintf("%v%v set the system optional", key, logogram)
	}
	return opt.Description
}

// option 选项字典
type OptionDick struct {
	Data map[string]Option
	App  *App
}

//数据新增
// key, descrip, logogram string
func (od *OptionDick) Add(datas ...string) *OptionDick {
	if vlen := len(datas); vlen > 0 {
		key := datas[0]
		var descrip, logogram string
		if vlen > 1 {
			descrip = datas[1]
		}
		if vlen > 2 {
			logogram = datas[2]
		}

		od.Data[key] = Option{
			Key:         key,
			Description: descrip,
			Logogram:    logogram,
			Value:       nil,
			RawValue:    "",
		}
	}
	return od
}

// 新的 option 字典初始化
func NewOptsDick(a *App) *OptionDick {
	od := &OptionDick{
		Data: map[string]Option{},
		App:  a,
	}
	return od
}
