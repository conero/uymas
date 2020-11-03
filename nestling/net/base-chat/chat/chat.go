package chat

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	DefChatPort    = "7400"
	DefChatNetwork = "tcp"
	DefChatHost    = "127.0.0.1"
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
	if "" == action {
		action = c.Action
	}
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

//直接发送信息
func (c *Address) SendValue(conn net.Conn, value *url.Values) error {
	content := c.Send("", value)
	_, er := conn.Write([]byte(content))
	return er
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
func RespondContent(conn net.Conn) (*Address, error) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}
	content := string(buf[0:n])
	content = strings.TrimSpace(content)

	return ParseAddress(content), nil
}

type LogLevel int

const (
	LogInfo LogLevel = iota
	LogDebug
	LogWarning
	LogError
	LogFatal
)

type Logger struct {
	level LogLevel
}

func (l *Logger) print(level LogLevel, format string, args ...interface{}) {
	var lv string
	switch level {
	case LogInfo:
		lv = "INFO"
	case LogDebug:
		lv = "DEBUG"
	case LogWarning:
		lv = "WARNING"
	case LogError:
		lv = "ERROR"
	case LogFatal:
		lv = "FATAL"
	default:
		lv = "INFO"
	}
	if l.level <= level {
		if len(args) > 0 {
			format = fmt.Sprintf(format, args...)
			log.Printf("[%v] %v", lv, format)
		} else {
			log.Printf("[%v] %v", lv, format)
		}
	}
}

func (l *Logger) Info(format string, args ...interface{}) {
	l.print(LogInfo, format, args...)
}

func (l *Logger) Debug(format string, args ...interface{}) {
	l.print(LogDebug, format, args...)
}

func (l *Logger) Warning(format string, args ...interface{}) {
	l.print(LogWarning, format, args...)
}

func (l *Logger) Error(format string, args ...interface{}) {
	l.print(LogError, format, args...)
}

func (l *Logger) Fatal(format string, args ...interface{}) {
	l.print(LogError, format, args...)
	os.Exit(1)
}
func NewLogger(level string) *Logger {
	var lv LogLevel
	switch level {
	case "info":
		lv = LogInfo
	case "debug":
		lv = LogDebug
	case "warning":
		lv = LogWarning
	case "error":
		lv = LogError
	case "fatal":
		lv = LogFatal
	default:
		lv = LogInfo
	}
	return &Logger{level: lv}
}

//默认日志
var Log = NewLogger("")
