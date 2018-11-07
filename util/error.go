package util

// @Date：   2018/11/7 0007 10:50
// @Author:  Joshua Conero
// @Name:    错误信息

// 基本处理类
type BaseError struct {
	Msg string
}

func (be *BaseError) Error() string {
	return be.Msg
}
