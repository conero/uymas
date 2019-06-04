// 文件系统处理扩展，文件/目录操作
package fs

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// @Date：   2018/11/6 0006 14:33
// @Author:  Joshua Conero
// @Name:    读写

// 不稳定，不建议使用
// Fs 系统的文件读写接口尝试
type FsReaderWriter struct {
	content  []byte
	dstFile  string
	srcFile  string
	errorMsg string
}

func (f *FsReaderWriter) Read(p []byte) (n int, err error) {
	if f.content != nil {
		return len(f.content), nil
	} else if f.srcFile != "" {
		content, err := ioutil.ReadFile(f.srcFile)
		f.content = content
		return len(content), err
	}
	return 0, f
}

func (f *FsReaderWriter) Write(p []byte) (n int, err error) {
	if f.content != nil {
		if f.dstFile != "" {
			err := ioutil.WriteFile(f.dstFile, f.content, 0755)
			return len(f.content), err
		} else {
			f.errorMsg = "未设置目标文件，文件写入失败！"
			return len(f.content), f
		}
	}
	return 0, nil
}

func (f *FsReaderWriter) Error() string {
	return f.errorMsg
}

// 基于DIY 实现 Copy
func copyBaseDiy(dstFile, srcFile string) (bool, error) {
	frw := &FsReaderWriter{
		dstFile: dstFile,
		srcFile: srcFile,
	}
	if _, err := io.Copy(frw, frw); err != nil {
		return false, err
	}
	return true, nil
}

// 文件复制
// 基于读写
func Copy(dstFile, srcFile string) (bool, error) {
	// 获取源文件
	content, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return false, err
	}
	// 覆盖新的文件
	err = ioutil.WriteFile(dstFile, content, 0755)
	if err != nil {
		return false, err
	}
	return true, nil
}

// 全目录文件复制
func CopyDir(dst, src string) {
	dst = StdDir(dst)
	src = StdDir(src)
	if files, err := ioutil.ReadDir(src); err == nil {
		CheckDir(dst)
		for _, fl := range files {
			d1 := dst + fl.Name()
			s1 := src + fl.Name()
			if fl.IsDir() {
				d1 += "/"
				s1 += "/"
				CopyDir(d1, s1)
			} else {
				Copy(d1, s1)
			}
		}
	}
}

// 检测目录，不存在则并创建
// 获取并返回标准目录
func CheckDir(dir string) string {
	dir = StdDir(dir)
	_, err := os.Open(dir)
	if err != nil {
		os.MkdirAll(dir, 0666)
	}
	return dir
}

// 检测目录是否存在
func IsDir(dir string) bool {
	_, err := os.Open(dir)
	if err != nil {
		return false
	}
	return true
}

// 文件/文件等路径是否存在
func ExistPath(vpath string) bool {
	_, err := os.Stat(vpath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// 获取标准目录
func StdDir(d string) string {
	d = strings.Replace(d, "\\", "/", -1)
	if d != "" && "/" != d[len(d)-1:] {
		d += "/"
	}
	return d
}

// 文件尾部附加内容
func Append(filename, text string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	_, err = f.WriteString(text)
	f.Close()
	return err
}
