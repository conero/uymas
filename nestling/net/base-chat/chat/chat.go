package chat

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"strings"
	"time"
)

type Address struct {
	Protocol string
	Path     string
	Action   string
	URL      *url.URL
	Content  string
}

//消息发送
func (c *Address) Send(action string, value *url.Values) string {
	param := ""
	if value != nil {
		param = value.Encode()
		if param != "" {
			param = "?" + param
		}
	}
	var content = fmt.Sprintf("%v://%v%v", c.Protocol, action, param)
	return content
}

//`protocol://action`
func ParseAddress(content string) *Address {
	adr := &Address{
		Content: content,
	}
	seq := "://"
	index := strings.Index(content, seq)
	if index > 0 {
		tmpQue := strings.Split(content, seq)
		adr.Protocol = tmpQue[0]
		adr.Path = tmpQue[1]
		u, err := url.Parse(adr.Path)
		if err == nil {
			adr.URL = u
			adr.Action = u.Path
		}
	}
	return adr
}

//超时检测
func Timer(d time.Duration) func() bool {
	start := time.Now()
	return func() bool {
		var overtime bool
		now := time.Now()
		if now.Sub(start) > d {
			overtime = true
		}
		return overtime
	}
}

//获取内容
func RespondContent(conn net.Conn) *Address {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal("read err:", err)
		return nil
	}
	content := string(buf[0:n])
	content = strings.TrimSpace(content)

	return ParseAddress(content)
}
