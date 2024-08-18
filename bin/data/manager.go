package data

import (
	"gitee.com/conero/uymas/v2/util/fs"
	"path"
)

// Manager the cli application data manger help
type Manager struct {
	basedir     string
	CommandName string
}

func NewManager(baseDir string) *Manager {
	return &Manager{
		basedir: baseDir,
	}
}

// AbsPath Obtain absolute parameter path
func (c *Manager) AbsPath(pathName string) string {
	return fs.StdPathName(path.Join(c.basedir, pathName))
}

func (c *Manager) Dir() string {
	return c.basedir
}
func CliManager() *Manager {
	name := "." + fs.AppName() + "-runtime"
	basedir := fs.RootPath(name)
	return NewManager(basedir)
}
