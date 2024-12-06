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
	"errors"
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
func ErrorIf(err error, prefixErr ...error) {
	if err == nil {
		return
	}
	vErr := errors.Join(prefixErr...)
	vErr = errors.Join(err)
	vLgr.Errorf(vErr.Error())
}

func Fatal(message string, args ...any) {
	vLgr.Fatalf(message, args...)
	os.Exit(0)
}

func FatalIf(err error, prefixErr ...error) {
	if err == nil {
		return
	}
	vErr := errors.Join(prefixErr...)
	vErr = errors.Join(err)
	vLgr.Fatalf(vErr.Error())
	os.Exit(1)
}

func Pref(logPref string) logger.Logger {
	vLgr.Pref(logPref)
	return *vLgr
}

func SetFlag(flag int) {
	vLgr.SetFlags(flag)
}
