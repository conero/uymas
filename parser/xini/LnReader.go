/**
LnReader 文件行阅读器
2018年7月10日 星期二
*/
package xini

import (
	"bufio"
	"os"
)

// 行阅读器
type LnReader struct {
	Filename string // 文件名
	error
}

// 实例阅读器
func NewLnRer(filename string) *LnReader {
	return &LnReader{
		Filename: filename,
	}
}

// 行扫描
func (ln *LnReader) Scan(callback func(line string)) bool {
	fs, err := os.Open(ln.Filename)
	if err == nil {
		buf := bufio.NewReader(fs)
		for {
			line, err2 := buf.ReadString('\n')
			callback(line)
			// 错误
			if err2 != nil {
				break
			}
		}
	} else {
		ln.error = err
		return false
	}
	return true
}
