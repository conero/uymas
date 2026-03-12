// Package lgr An instance of library logger, used for direct output from the command line, etc.
//
// To change logger level should use the system environment "UYMAS_LGR_LEVEL",
// to close logger color style shou use the system environment "UYMAS_LGR_NOCOLOR",
// to change debug maker for test use the system environment "UYMAS_TMP_MARK",
// like:
//
//	# window powershell
//	$env:UYMAS_LGR_LEVEL='all'
//	$env:UYMAS_LGR_NOCOLOR='true'
//	$env:UYMAS_TMP_MARK='TMarkShouldDEL'
//	$env:UYMAS_LGR_FILE='log.txt' or true/1
//
//	# linux shell
//	export UYMAS_LGR_LEVEL=all
//	export UYMAS_LGR_NOCOLOR=true
//	export UYMAS_TMP_MARK=TMarkShouldDEL
//	export UYMAS_LGR_FILE=log.txt or true/1
//
// if not info by default.
package lgr

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"gitee.com/conero/uymas/v2/cli/ansi"
	"gitee.com/conero/uymas/v2/logger"
	"gitee.com/conero/uymas/v2/rock"
	"gitee.com/conero/uymas/v2/util/fs"
)

var vLgr *logger.Logger

const (
	// EnvLevelKey try set the lgr level by system environment, like `$ export UYMAS_LGR_LEVEL=info`
	EnvLevelKey = "UYMAS_LGR_LEVEL"
	// EnvNoColorKey try set the lgr no color by system environment, like `$ export UYMAS_LGR_NOCOLOR=true`
	EnvNoColorKey = "UYMAS_LGR_NOCOLOR"
	// EnvMarkKey try set the lgr mark by system environment, like `$ export UYMAS_TMP_MARK=mark`
	EnvMarkKey = "UYMAS_TMP_MARK"
	// EnvLogFile try set the lgr log file by system environment, like `$ export UYMAS_LGR_FILE=/tmp/log.log`
	EnvLogFile = "UYMAS_LGR_FILE"
)

func createLog() {
	lvl := fs.GetenvOr(EnvLevelKey, logger.LevelInfo)
	noColor := strings.ToLower(fs.GetenvOr(EnvNoColorKey, "false"))
	lgrConfig := logger.Config{
		Level: lvl,
	}
	logFile := fs.GetenvOr(EnvLogFile, "")
	if logFile != "" {
		lgrConfig.Driver = logger.DriverFile
		if !rock.InList([]string{"true", "1"}, logFile) {
			lgrConfig.OutputDir = logFile
		}
		noColor = "true"
	}
	vLgr = logger.NewLogger(lgrConfig)
	if noColor != "" && noColor != "false" && noColor != "0" {
		vLgr.NoColor()
	}
}

func Log() logger.Logger {
	return *getLgr()
}

func getLgr() *logger.Logger {
	if vLgr == nil {
		createLog()
	}
	return vLgr
}

func Trace(message string, args ...any) {
	getLgr().Tracef(message, args...)
}

func Debug(message string, args ...any) {
	getLgr().Debugf(message, args...)
}

func Info(message string, args ...any) {
	getLgr().Infof(message, args...)
}

func Warn(message string, args ...any) {
	getLgr().Warnf(message, args...)
}

func Error(message string, args ...any) {
	getLgr().Errorf(message, args...)
}

// ErrorIf print error message only when err is not nil
func ErrorIf(err error, prefixErr ...error) {
	if err == nil {
		return
	}
	vErr := errors.Join(prefixErr...)
	vErr = errors.Join(err, vErr)
	vLgr.Errorf(vErr.Error())
}

func Fatal(message string, args ...any) {
	getLgr().Fatalf(message, args...)
	os.Exit(0)
}

func FatalIf(err error, prefixErr ...error) {
	if err == nil {
		return
	}
	vErr := errors.Join(prefixErr...)
	vErr = errors.Join(err, vErr)
	getLgr().Fatalf(vErr.Error())
	os.Exit(1)
}

func Pref(logPref string) logger.Logger {
	getLgr().Pref(logPref)
	return *vLgr
}

func SetFlag(flag int) {
	getLgr().SetFlags(flag)
}

// TmpMark temporary tags are used for debugging, and debugging should be removed before release
//
// to global search keyword `lgr.TmpMark` then remove it.
func TmpMark(mark any, args ...any) {
	markString := fmt.Sprintf("%v", mark)
	markTitle := fs.GetenvOr(EnvMarkKey, "TMarkShouldDEL")
	_, flPath, flLine, _ := runtime.Caller(1)
	markString = ansi.Style("<"+markTitle+">", ansi.Red, ansi.BkgWhiteBr, ansi.Italic, ansi.Twinkle) +
		ansi.Style(fmt.Sprintf(" %s(%d) ", filepath.Base(flPath), flLine), ansi.Green) +
		markString
	getLgr().Errorf(markString, args...)
}

// TmpMarkExit Execute temporary test and exit the subsequent program
func TmpMarkExit(mark any, args ...any) {
	markString := fmt.Sprintf("%v", mark)
	markTitle := fs.GetenvOr(EnvMarkKey, "TMarkShouldDEL")
	_, flPath, flLine, _ := runtime.Caller(1)
	markString = ansi.Style("<"+markTitle+">", ansi.Red, ansi.BkgWhiteBr, ansi.Italic, ansi.Twinkle) +
		ansi.Style(fmt.Sprintf(" %s(%d) ", filepath.Base(flPath), flLine), ansi.Green) +
		markString
	getLgr().Errorf(markString, args...)
	os.Exit(1)
}
