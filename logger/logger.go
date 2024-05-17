// Package logger basic and simple logger for application, it base the go embed `log` package.
// it's able to control output by log level, level order `all<debug<info<warn<error<none`.
package logger

import (
	"bytes"
	"fmt"
	"gitee.com/conero/uymas/bin/butil"
	"gitee.com/conero/uymas/bin/color"
	"gitee.com/conero/uymas/fs"
	"log"
	"os"
	"strings"
	"time"
)

type Level int8

// logging Level constant
const (
	LogAll Level = iota
	LogTrace
	LogDebug
	LogInfo
	LogWarn
	LogError
	LogNone
)

// string level values.
const (
	LevelAll   = "all"
	LevelError = "error"
	LevelWarn  = "warn"
	LevelInfo  = "info"
	LevelDebug = "debug"
	LevelTrace = "trace"
	LevelNone  = "none"
)

// DriverStdout logging driver support builtin
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
	case LogTrace:
		prefix = "TRACE"
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
	bufDriver    *bytes.Buffer // only when Config.Driver is `buffer`
	logger       *log.Logger
	Level        Level
	cfg          Config
	DisableColor bool
}

func (l *Logger) autoColor(prefix string, level Level) string {
	if l.cfg.Driver != "" && l.cfg.Driver != DriverStdout {
		return prefix
	}

	if l.DisableColor {
		return prefix
	}

	var ansi int
	switch level {
	case LogError:
		ansi = color.AnsiTextRedBr
	case LogWarn:
		ansi = color.AnsiTextYellowBr
	case LogInfo:
		ansi = color.AnsiTextGreenBr
	case LogDebug:
		ansi = color.AnsiTextCyanBr
	case LogTrace:
		ansi = color.AnsiTextBlackBr
	}

	if ansi < 1 {
		return prefix
	}

	return color.StyleByAnsi(ansi, prefix)
}

func (l *Logger) Format(prefix, message string, args ...any) {
	l.logger.Printf("[%v] %v", prefix, fmt.Sprintf(message, args...))
}

// format logging by level, logging creator
func (l *Logger) formatLevel(level Level, message string, args ...any) {
	if l.Level > level {
		return
	}

	l.Format(l.autoColor(Prefix(level), level), message, args...)
}

// output logging with callback, logging creator
func (l *Logger) outputFunc(level Level, callback func() string) {
	if l.Level > level {
		return
	}
	l.formatLevel(level, callback())
}

func (l *Logger) Debugf(message string, args ...any) {
	l.formatLevel(LogDebug, message, args...)
}

func (l *Logger) Tracef(message string, args ...any) {
	l.formatLevel(LogTrace, message, args...)
}

func (l *Logger) DebugFunc(callback func() string) {
	l.outputFunc(LogDebug, callback)
}

func (l *Logger) Infof(message string, args ...any) {
	l.formatLevel(LogInfo, message, args...)
}

func (l *Logger) InfoFunc(callback func() string) {
	l.outputFunc(LogInfo, callback)
}

func (l *Logger) Warnf(message string, args ...any) {
	l.formatLevel(LogWarn, message, args...)
}

func (l *Logger) WarnFunc(callback func() string) {
	l.outputFunc(LogWarn, callback)
}

func (l *Logger) Errorf(message string, args ...any) {
	l.formatLevel(LogError, message, args...)
}

func (l *Logger) ErrorFunc(callback func() string) {
	l.outputFunc(LogError, callback)
}

// Log get embed go lib log when you need the instance.
func (l *Logger) Log() *log.Logger {
	return l.logger
}

func (l *Logger) Buffer() *bytes.Buffer {
	return l.bufDriver
}

// CoverLevel cover input string level into `Level`
// Deprecated: replace by the func `ToLevel`
func CoverLevel(lvl string, defLevel Level) Level {
	lv, er := ToLevel(lvl, defLevel)
	if er != nil {
		panic(er)
	}
	return lv
}

// ToLevel turn string to level
func ToLevel(lvl string, args ...Level) (Level, error) {
	var (
		lv Level
		er error
	)
	if len(args) > 0 {
		lv = args[0]
	}
	if lvl == "" { // empty input use default Level
		return lv, nil
	}
	lvl = ShortCover(lvl)
	switch strings.ToLower(lvl) {
	case LevelAll:
		lv = LogAll
	case LevelError, "err":
		lv = LogError
	case LevelWarn, "warning":
		lv = LogWarn
	case LevelInfo:
		lv = LogInfo
	case LevelDebug:
		lv = LogDebug
	case LevelTrace:
		lv = LogTrace
	case LevelNone, "no", "mute", "quiet":
		lv = LogNone
	default:
		er = fmt.Errorf("%v: invalid level param, reference value all, error, warn, info, debug, none", lvl)
	}
	return lv, er
}

// ShortCover short level string cover to matched level string.
// rule:
//
//	`a/A -> all`
//	`e/E -> error`
//	`w/W -> warning`
//	`i/I -> info`
//	`d/D -> debug`
//	`t/T -> trace`
//	`n/N -> none`
func ShortCover(short string) (lvlStr string) {
	lvlStr = short
	switch strings.ToLower(short) {
	case "a":
		lvlStr = LevelAll
	case "e":
		lvlStr = LevelError
	case "w":
		lvlStr = LevelWarn
	case "i":
		lvlStr = LevelInfo
	case "d":
		lvlStr = LevelDebug
	case "t":
		lvlStr = LevelTrace
	case "n":
		lvlStr = LevelNone
	}
	return
}

// NewLogger build a simple logger user it.
func NewLogger(cfgs ...Config) *Logger {
	var cfg Config
	if len(cfgs) > 0 {
		cfg = cfgs[0]
	} else {
		cfg = DefaultConfig
	}
	logging := &Logger{
		cfg: cfg,
	}
	// default base log level is `Warn`
	lv, er := ToLevel(cfg.Level, LogWarn)
	if er != nil {
		panic(er)
	}
	if cfg.Log == nil { // 默认日志
		if lv != LogNone {
			if cfg.Driver == DriverFile {
				output := cfg.OutputDir
				if output == "" {
					output = butil.RootPath("/.runtime")
				} else {
					if !fs.ExistPath(output) {
						output = butil.RootPath(output)
					}
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
