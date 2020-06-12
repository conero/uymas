package butil

import (
	"os"
	"testing"
)

func TestRootDir(t *testing.T) {
	t.Log(GetBasedir())
	t.Log(os.Args)
}
