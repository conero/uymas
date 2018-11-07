package fs

import (
	"io"
	"io/ioutil"
	"os"
)

// @Date：   2018/11/6 0006 14:33
// @Author:  Joshua Conero
// @Name:    读写

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

// 检测目录，并删除
func CheckDir(dir string) {
	_, err := os.Open(dir)
	if err != nil {
		os.MkdirAll(dir, 0666)
	}
}

// 检测目录是否存在
func IsDir(dir string) bool {
	_, err := os.Open(dir)
	if err != nil {
		return false
	}
	return true
}
