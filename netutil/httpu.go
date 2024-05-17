// Package netutil net util
package netutil

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// HttpUtil http util tool
type HttpUtil struct {
}

// PostForm post file and form data. support multi files or data
func (hu HttpUtil) PostForm(rUrl string, files map[string]string, data map[string]string) (*http.Response, error) {
	buf := &bytes.Buffer{}
	bufWrite := multipart.NewWriter(buf)

	// 文件
	for flKey, flName := range files {
		fw, er := bufWrite.CreateFormFile(flKey, filepath.Base(flName))
		if er != nil {
			panic(er)
		}
		fh, er := os.Open(flName)
		if er != nil {
			panic(er)
		}

		defer fh.Close()
		_, _ = io.Copy(fw, fh)
	}

	// 数据库
	for dk, dv := range data {
		_ = bufWrite.WriteField(dk, dv)
	}

	//获取请求Content-Type类型,后面有用
	contentType := bufWrite.FormDataContentType()
	bufWrite.Close()

	//创建 http 请求客服端
	client := &http.Client{}

	req, er := http.NewRequest("POST", rUrl, nil)
	if er != nil {
		panic(er)
	}

	//头部类型
	req.Header.Set("Content-Type", contentType)
	req.Body = io.NopCloser(buf)

	res, er := client.Do(req)
	return res, er
}

// PostFormString postForm post add the directly by string content
func (hu HttpUtil) PostFormString(rUrl string, files map[string]string, data map[string]string) string {
	res, er := hu.PostForm(rUrl, files, data)
	if er != nil {
		panic(er)
	}

	cttBys, er := io.ReadAll(res.Body)
	if er != nil {
		panic(er)
	}

	defer res.Body.Close()

	return string(cttBys)
}

var Httpu HttpUtil
