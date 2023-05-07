// Package lgr An instance of library `logger`, used for direct output from the command line, etc
package lgr

import (
	"gitee.com/conero/uymas/logger"
)

var vLgr *logger.Logger

func init() {
	vLgr = logger.NewLogger(logger.Config{
		Level: logger.LevelAll,
	})
}

func Log() logger.Logger {
	return *vLgr
}

func Debug(message string, args ...any) {
	vLgr.Debugf(message, args...)
}

func Info(message string, args ...any) {
	vLgr.Infof(message, args...)
}

func Warn(message string, args ...any) {
	vLgr.Warnf(message, args...)
}

func Error(message string, args ...any) {
	vLgr.Errorf(message, args...)
}
