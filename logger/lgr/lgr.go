// Package lgr An instance of library logger, used for direct output from the command line, etc.
//
// To change logger level shou use the system environment "UYMAS_LGR_LEVEL", like:
//
//	//window powershell
//	$env:UYMAS_LGR_LEVEL='all'
//
//	// linux shell
//	export UYMAS_LGR_LEVEL=all
//
// if not info by default.
package lgr

import (
	"gitee.com/conero/uymas/v2/logger"
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

// ErrorIf print error message only when err is not nil
func ErrorIf(err error) {
	if err == nil {
		return
	}
	vLgr.Errorf(err.Error())
}

func FatalIf(err error) {
	if err == nil {
		return
	}
	vLgr.Fatalf(err.Error())
	os.Exit(1)
}

func Pref(logPref string) logger.Logger {
	vLgr.Pref(logPref)
	return *vLgr
}
