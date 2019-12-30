// bin util package
// will not run the init(), but bin will
package butil

import (
	"github.com/conero/uymas/fs"
	"os"
	"path"
)

//get the root base Dir
func GetBasedir() string {
	rwd := os.Args[0]
	rwd = fs.StdPathName(rwd)
	rwd = path.Dir(rwd)
	return rwd
}
