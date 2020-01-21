// bin util package
// will not run the init(), but bin will
package butil

import (
	"github.com/conero/uymas/fs"
	"os"
	"path"
	"strings"
)

//get the root base Dir
func GetBasedir() string {
	rwd := os.Args[0]
	rwd = fs.StdPathName(rwd)
	rwd = path.Dir(rwd)
	return rwd
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
