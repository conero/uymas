// net util
package netutil

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// http 助手
type HttpUtil struct {
}

//post 文件以及键值
//支持多文件，以及多参数
func (hu HttpUtil) PostForm(rUrl string, files map[string]string, data map[string]string) (*http.Response, error) {
	buf := &bytes.Buffer{}
	bufWrite := multipart.NewWriter(buf)

	// 文件
	if files != nil {
		for flKey, flName := range files {
			fw, er := bufWrite.CreateFormFile(flKey, filepath.Base(flName))
			if er != nil {
				panic(er)
			}
			fh, er := os.Open(flName)
			if er != nil {
				panic(er)
			}

			io.Copy(fw, fh)

			fh.Close()
		}
	}

	// 数据库
	if data != nil {
		for dk, dv := range data {
			bufWrite.WriteField(dk, dv)
		}
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
	req.Body = ioutil.NopCloser(buf)

	res, er := client.Do(req)
	return res, er
}

// postForm 直接返回字符串内容
func (hu HttpUtil) PostFormString(rUrl string, files map[string]string, data map[string]string) string {
	res, er := hu.PostForm(rUrl, files, data)

	cttBys, er := ioutil.ReadAll(res.Body)
	if er != nil {
		panic(er)
	}

	defer res.Body.Close()

	return string(cttBys)
}

var Httpu HttpUtil
