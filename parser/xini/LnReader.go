/**
LnReader 文件行阅读器
2018年7月10日 星期二
*/

package xini

import (
	"bufio"
	"errors"
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

// Scan file lines
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

// ScanWithFlInfo file lines
func (ln *LnReader) ScanWithFlInfo(callback func(line string)) (os.FileInfo, error) {
	fs, err := os.Open(ln.Filename)
	if err != nil {
		return nil, errors.Join(errors.New("文件读取失败"), err)
	}
	defer fs.Close()

	buf := bufio.NewReader(fs)
	for {
		line, err2 := buf.ReadString('\n')
		callback(line)
		// 错误
		if err2 != nil {
			break
		}
	}
	stat, err := fs.Stat()
	if err != nil {
		return nil, errors.Join(errors.New("获取文件信息错误"), err)
	}

	return stat, nil
}
