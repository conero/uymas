// bin util package
// will not run the init(), but bin will
package butil

import (
	"fmt"
	"github.com/conero/uymas/fs"
	"os"
	"path"
	"strings"
)

var (
	cacheBaseDir string
)

//get the root base Dir
func GetBasedir() string {
	if cacheBaseDir != "" {
		return cacheBaseDir
	}
	rwd := os.Args[0]
	rwd = fs.StdPathName(rwd)
	rwd = path.Dir(rwd)
	cacheBaseDir = rwd
	return rwd
}

//the path dir by application same location.
func GetPathDir(vPath string) string {
	return fmt.Sprintf("%v%v", GetBasedir(), vPath)
}

//make the string to bin/Args, it's used in interactive cli
func StringToArgs(str string) []string {
	a := []string{}
	strArr := strings.Split(str, " ")
	for _, sv := range strArr {
		sv = strings.TrimSpace(sv)
		if sv == "" {
			continue
		}
		a = append(a, sv)
	}

	return a
}
