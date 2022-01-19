// Package logger basic and simple logger for application, it base the go embed `log` package.
// it's able to control output by log level, level order `all<debug<info<warn<error<none`.
package logger

import (
	"bytes"
	"fmt"
	"gitee.com/conero/uymas/bin/butil"
	"gitee.com/conero/uymas/fs"
	"log"
	"os"
	"strings"
	"time"
)

type Level int8

// logging Level constant
//
const (
	LogAll Level = iota
	LogDebug
	LogInfo
	LogWarn
	LogError
	LogNone
)

// DriverStdout logging driver support builtin
//
const (
	DriverStdout = "stdout"
	DriverFile   = "file"
	DriverBuffer = "buffer"
)

func Prefix(level Level) string {
	var prefix string
	switch level {
	case LogAll:
		prefix = "ALL"
	case LogDebug:
		prefix = "DEBUG"
	case LogInfo:
		prefix = "INFO"
	case LogWarn:
		prefix = "WARN"
	case LogError:
		prefix = "ERROR"
	}
	return prefix
}

type Logger struct {
	bufDriver *bytes.Buffer // only when Config.Driver is `buffer`
	logger    *log.Logger
	Level     Level
}

func (l Logger) Format(prefix, message string, args ...interface{}) {
	l.logger.Printf("[%v] %v", prefix, fmt.Sprintf(message, args...))
}

// format logging by level, logging creator
func (l Logger) formatLevel(level Level, message string, args ...interface{}) {
	if l.Level > level {
		return
	}
	l.Format(Prefix(level), message, args...)
}

// output logging with callback, logging creator
func (l Logger) outputFunc(level Level, callback func() string) {
	if l.Level > level {
		return
	}
	l.formatLevel(level, callback())
}

func (l Logger) Debugf(message string, args ...interface{}) {
	l.formatLevel(LogDebug, message, args...)
}

func (l Logger) DebugFunc(callback func() string) {
	l.outputFunc(LogDebug, callback)
}

func (l Logger) Infof(message string, args ...interface{}) {
	l.formatLevel(LogInfo, message, args...)
}

func (l Logger) InfoFunc(callback func() string) {
	l.outputFunc(LogInfo, callback)
}

func (l Logger) Warnf(message string, args ...interface{}) {
	l.formatLevel(LogWarn, message, args...)
}

func (l Logger) WarnFunc(callback func() string) {
	l.outputFunc(LogWarn, callback)
}

func (l Logger) Errorf(message string, args ...interface{}) {
	l.formatLevel(LogError, message, args...)
}

func (l Logger) ErrorFunc(callback func() string) {
	l.outputFunc(LogError, callback)
}

// Log get embed go lib log when you need the instance.
func (l Logger) Log() *log.Logger {
	return l.logger
}

func (l Logger) Buffer() *bytes.Buffer {
	return l.bufDriver
}

// NewLogger build a simple logger user it.
func NewLogger(cfgs ...Config) *Logger {
	var cfg Config
	if len(cfgs) > 0 {
		cfg = cfgs[0]
	} else {
		cfg = DefaultConfig
	}
	logging := &Logger{}
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
	default:
		panic(fmt.Sprintf("invalid level param, reference value: all, error, warn, info, debug, none"))
	}

	if cfg.Log == nil { // 默认日志
		if lv != LogNone {
			if cfg.Driver == DriverFile {
				output := cfg.OutputDir
				if output == "" {
					output = butil.GetPathDir("/.runtime")
				} else {
					output = butil.GetPathDir(output)
				}
				now := time.Now()
				output = fs.CheckDir(fmt.Sprintf("%v/%v", output, now.Format("2006/01")))
				fl, er := os.OpenFile(fmt.Sprintf("%v/%v.log", output, now.Format("02")),
					os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
				if er == nil {
					cfg.Log = log.New(fl, "", log.Ltime)
				}
			} else if cfg.Driver == DriverBuffer {
				var buf bytes.Buffer
				cfg.Log = log.New(&buf, "", log.Ltime)
				logging.bufDriver = &buf
			}
		}

		// 降级处理，所有驱动解析失败的使用控制台
		if cfg.Log == nil {
			cfg.Log = log.New(os.Stdout, "", log.Ltime)
		}
	}

	logging.Level = lv
	logging.logger = cfg.Log
	return logging
}
