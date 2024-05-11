// Package lgr An instance of library `logger`, used for direct output from the command line, etc
package lgr

import (
	"gitee.com/conero/uymas/logger"
	"os"
)

var vLgr *logger.Logger

const (
	// EnvLevelKey try set the lgr level by system environment, like `$ export EnvLevelKey=info`
	EnvLevelKey = "UYMAS_LGR_LEVEL"
)

func init() {
	lvl := os.Getenv(EnvLevelKey)
	if lvl == "" {
		lvl = logger.LevelInfo
	}
	vLgr = logger.NewLogger(logger.Config{
		Level: lvl,
	})
}

func Log() logger.Logger {
	return *vLgr
}

func Trace(message string, args ...any) {
	vLgr.Tracef(message, args...)
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
