package butil

import (
	"os"
	"testing"
)

func TestRootDir(t *testing.T) {
	t.Log(RootDir())
	t.Log(os.Args)
}
