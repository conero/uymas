package fs

import (
	"io"
	"os"
	"path"
	"regexp"
	"strings"
)

// @Date：   2018/11/6 0006 14:33
// @Author:  Joshua Conero
// @Name:    读写

// FsReaderWriter [Experimental] file read interface
//
// Deprecated: next major version remove, because never use
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
		content, err := os.ReadFile(f.srcFile)
		f.content = content
		return len(content), err
	}
	return 0, f
}

func (f *FsReaderWriter) Write(p []byte) (n int, err error) {
	if f.content != nil {
		if f.dstFile != "" {
			err := os.WriteFile(f.dstFile, f.content, 0755)
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
// @todo test if real need
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

// Copy file by io
func Copy(dstFile, srcFile string) (bool, error) {
	// 获取源文件
	content, err := os.ReadFile(srcFile)
	if err != nil {
		return false, err
	}
	// 覆盖新的文件
	err = os.WriteFile(dstFile, content, 0755)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CopyDir copy all file in a dir
func CopyDir(dst, src string) {
	dst = StdDir(dst)
	src = StdDir(src)
	if files, err := os.ReadDir(src); err == nil {
		CheckDir(dst)
		for _, fl := range files {
			d1 := dst + fl.Name()
			s1 := src + fl.Name()
			if fl.IsDir() {
				d1 += "/"
				s1 += "/"
				CopyDir(d1, s1)
			} else {
				_, _ = Copy(d1, s1)
			}
		}
	}
}

// CheckDir checkout if dir exist, when not exist will try to build it and return the path.
func CheckDir(dir string) string {
	dir = StdDir(dir)
	_, err := os.Open(dir)
	if err != nil {
		_ = os.MkdirAll(dir, 0666)
	}
	return dir
}

// CheckFileDir detect whether the parent directory where the file is located exists,
// and use it to automatically generate the parent directory when generating files (adaptable)
func CheckFileDir(filename string) string {
	return CheckDir(path.Dir(filename))
}

// IsDir checkout string path is dir.
func IsDir(dir string) bool {
	fi, err := os.Stat(dir)
	return err == nil && fi.IsDir()
}

// ExistPath checkout the path of file/dir exist.
func ExistPath(vPath string) bool {
	_, err := os.Stat(vPath)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

// StdDir the standard dir format
func StdDir(d string) string {
	d = StdPathName(d)
	if d != "" && "/" != d[len(d)-1:] {
		d += "/"
	}
	return d
}

// StdPathName the standard path format
func StdPathName(vPath string) string {
	if vPath != "" {
		vPath = strings.ReplaceAll(vPath, "\\", "/")
		reg := regexp.MustCompile(`/{2,}`)
		vPath = reg.ReplaceAllString(vPath, "/")
	}
	return vPath
}

// Append append content to a file
func Append(filename, text string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	_, err = f.WriteString(text)
	_ = f.Close()
	return err
}

// Put rewrite content to file
func Put(filename, text string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	_, err = f.WriteString(text)
	_ = f.Close()
	return err
}
