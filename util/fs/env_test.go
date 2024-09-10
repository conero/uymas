package fs

import (
	"testing"
)

func TestRootPath(t *testing.T) {
	baseDir := RootPath()
	rlPath := RootPath("///name.json")
	rfPath := baseDir + "name.json"

	if baseDir[len(baseDir)-1:] != "/" {
		t.Errorf("baseDir 不是标准的目录格式")
	}

	// case
	if rlPath != rfPath {
		t.Errorf("%s ≠ %s\n", rlPath, rfPath)
	}
}
