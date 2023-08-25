package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

// FileHash 文件hash，支持 md5/sha1
type FileHash struct {
	Vtype string // md5/sha1/sha256/sha512, 默认md5
}

type HashList struct {
	Filename string
	Hash     string
	Vpath    string
}

func (c *FileHash) DirList(vDir string) ([]HashList, error) {
	var hashLs []HashList
	entry, err := os.ReadDir(vDir)
	if err != nil {
		erLs := errors.New("目录读取错误")
		erLs = errors.Join(erLs, err)
		return hashLs, erLs
	}

	for _, ent := range entry {
		if ent.IsDir() {
			continue
		}

		vPath := path.Join(vDir, ent.Name())
		hash := c.GetHash(vPath)
		hashLs = append(hashLs, HashList{
			Filename: ent.Name(),
			Hash:     hash,
			Vpath:    vPath,
		})
	}

	return hashLs, nil
}

func (c *FileHash) PathList(vPath string) ([]HashList, error) {
	fi, err := os.Stat(vPath)
	if err != nil {
		err = errors.Join(errors.New("路径读取失败！"), err)
		return nil, err
	}

	if fi.IsDir() {
		return c.DirList(vPath)
	}

	var hashList []HashList

	hashList = append(hashList, HashList{
		Hash:     c.GetHash(vPath),
		Vpath:    vPath,
		Filename: path.Base(vPath),
	})

	return hashList, err
}

func (c *FileHash) GetHash(filepath string) string {
	fl, err := os.Open(filepath)
	if err != nil {
		return ""
	}
	vtype := strings.ToLower(c.Vtype)
	c.Vtype = vtype
	switch vtype {
	case "sha1":
		h := sha1.New()
		if _, err = io.Copy(h, fl); err != nil {
			return ""
		}
		return fmt.Sprintf("%x", h.Sum(nil))
	case "sha256":
		h := sha256.New()
		if _, err = io.Copy(h, fl); err != nil {
			return ""
		}
		return fmt.Sprintf("%x", h.Sum(nil))
	case "sha512":
		h := sha512.New()
		if _, err = io.Copy(h, fl); err != nil {
			return ""
		}
		return fmt.Sprintf("%x", h.Sum(nil))
	default:
		h := md5.New()
		if _, err = io.Copy(h, fl); err != nil {
			return ""
		}
		c.Vtype = "md5"
		return fmt.Sprintf("%x", h.Sum(nil))
	}

}
