package chest

import "os/exec"

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
