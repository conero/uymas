package str

import (
	"net/url"
	"path"
	"strings"
)

/**
 * @DATE        2019/6/3
 * @NAME        Joshua Conero
 * @DESCRIPIT   Url 带命名空间的 URL 处理类
**/

// Url only any the url methods of namespace
type Url struct {
}

// AbsHref Get the absolute address of the path: path address path, URL is the top-level path, which can be empty
// BUG(AbsHref): Url.AbsHref 中解析 "vpath" `test/p1/p2` 与 `./test/p1/p2` 的一致性问题
func (u Url) AbsHref(vpath, vurl string) string {
	var href string

	// 需要解析， vurl 含 http(s)://
	if strings.Contains(vurl, "http://") || strings.Contains(vurl, "https://") {
		if u, err := url.Parse(vurl); err == nil {
			uHost := u.Scheme + "://" + u.Host

			// 字符连接处检测
			uFirstChar := ""
			if len(vpath) > 0 {
				uFirstChar = vpath[0:1]
			}
			if uFirstChar == "/" {
				href = uHost + path.Clean(vpath)
			} else {
				nS := u.Path + "/" + vpath
				nS = path.Clean(nS)
				href = uHost + nS
			}
		}
	} else {
		href = path.Clean(href)
	}

	return href
}
