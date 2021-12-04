// Package logger basic and simple logger for application
package logger

import (
	"fmt"
	"gitee.com/conero/uymas/bin/butil"
	"gitee.com/conero/uymas/fs"
	"log"
	"os"
	"strings"
	"time"
)

type Level int8

const (
	LogAll Level = iota
	LogDebug
	LogInfo
	LogWarn
	LogError
	LogNone
)

type Logger struct {
	logger *log.Logger
	Level  Level
}

// Config 日志配置文件
type Config struct {
	Log       *log.Logger
	Level     string
	Driver    string // 输出启驱动<stdout, file>
	OutputDir string // 输出日志，设置是为文件驱动
}

func (l Logger) Format(prefix, message string, args ...interface{}) {
	l.logger.Printf("[%v] %v", prefix, fmt.Sprintf(message, args...))
}

func (l Logger) Debugf(message string, args ...interface{}) {
	if l.Level > LogDebug {
		return
	}
	l.Format("DEBUG", message, args...)
}

func (l Logger) Infof(message string, args ...interface{}) {
	if l.Level > LogInfo {
		return
	}
	l.Format("INFO", message, args...)
}

func (l Logger) Warnf(message string, args ...interface{}) {
	if l.Level > LogWarn {
		return
	}
	l.Format("WARN", message, args...)
}

func (l Logger) Errorf(message string, args ...interface{}) {
	if l.Level > LogError {
		return
	}
	l.Format("ERROR", message, args...)
}

// Log get embed go lib log
func (l Logger) Log() *log.Logger {
	return l.logger
}

func NewLogger(cfg Config) *Logger {
	// default base log level is `Warn`
	lv := LogWarn
	switch strings.ToLower(cfg.Level) {
	case "all":
		lv = LogAll
	case "error", "err":
		lv = LogError
	case "warning", "warn":
		lv = LogWarn
	case "info":
		lv = LogInfo
	case "debug":
		lv = LogDebug
	case "none", "no", "mute", "quiet":
		lv = LogNone
	}

	if cfg.Log == nil { // 默认日志
		if lv != LogNone && (cfg.Driver == "file" || cfg.OutputDir != "") {
			output := cfg.OutputDir
			if output == "" {
				output = butil.GetPathDir("/.runtime")
			}
			now := time.Now()
			output = fs.CheckDir(fmt.Sprintf("%v/%v", output, now.Format("2006/01")))
			fl, er := os.OpenFile(fmt.Sprintf("%v/%v.log", output, now.Format("02")),
				os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
			if er == nil {
				cfg.Log = log.New(fl, "", log.LstdFlags)
			}
		}

		// 降级处理，所有驱动解析失败的使用控制台
		if cfg.Log == nil {
			cfg.Log = log.New(os.Stdout, "", log.LstdFlags)
		}
	}

	return &Logger{
		logger: cfg.Log,
		Level:  lv,
	}
}
