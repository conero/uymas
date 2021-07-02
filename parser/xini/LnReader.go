/**
LnReader 文件行阅读器
2018年7月10日 星期二
*/

package xini

import (
	"bufio"
	"os"
)

// LnReader the lines of file reader
type LnReader struct {
	Filename string // 文件名
	error
}

func NewLnRer(filename string) *LnReader {
	return &LnReader{
		Filename: filename,
	}
}

// Scan scan file lines
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
