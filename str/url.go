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

// 仅仅作为命名空间

// URL 相关处理类
type Url struct {
}

// BUG(AbsHref): Url.AbsHref 中解析 "vpath" `test/p1/p2` 与 `./test/p1/p2` 的一致性问题

// 获取路径的绝对地址： path 地址路径， url 为顶级路径可为空
func (u Url) AbsHref(vpath, vurl string) string {
	var href string

	// 需要解析， vurl 含 http(s)://
	if strings.Index(vurl, "http://") > -1 || strings.Index(vurl, "https://") > -1 {
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
