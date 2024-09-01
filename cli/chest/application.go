package chest

import (
	"errors"
	"fmt"
	"gitee.com/conero/uymas/v2/util/fs"
	"os/exec"
)

// CmdExist Used to check whether the command is available and to throw an exception.
func CmdExist(cmd string) (bool, error) {
	_, err := exec.LookPath(cmd)
	if err == nil {
		return true, nil
	}
	if _, ok := err.(*exec.Error); !ok {
		return false, err
	}
	return false, nil
}

// CmdAble determines whether a command is available, such as the command detection
func CmdAble(cmd string) bool {
	able, _ := CmdExist(cmd)
	return able
}

// CmdSearchRun try the search command and execute it
func CmdSearchRun(cmd string, args []string, children ...string) (output string, isSearch bool, runErr error) {
	baseDir := fs.RootPath()
	toRunFn := func(rlPath string) bool {
		execPath, err := exec.LookPath(rlPath)
		if err == nil {
			runnable := exec.Command(execPath, args...)
			byes, er := runnable.CombinedOutput()
			if er != nil {
				runErr = errors.Join(errors.New("command run error"), er)
			}
			output = string(byes)
			return true
		}
		return false
	}

	if toRunFn(baseDir + cmd) {
		isSearch = true
		return
	}

	for _, child := range children {
		rlPath := fs.StdPathName(baseDir + child + "/" + cmd)
		if toRunFn(rlPath) {
			isSearch = true
			return
		}
	}

	return "", false, fmt.Errorf("%s is not exist", cmd)
}
