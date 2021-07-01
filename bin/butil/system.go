package butil

// the cli application call system

import (
	"os/exec"
	"runtime"
)

// Clear @experimental it's a try
//clear the cli app
func Clear() error {
	osName := runtime.GOOS
	var cmd *exec.Cmd
	switch osName {
	case "windows":
		//try the `cls` and `clear` all fail
		cmd = exec.Command("cls")
	default:
		cmd = exec.Command("clear")
	}
	return cmd.Run()
}
