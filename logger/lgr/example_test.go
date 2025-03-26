package lgr

import (
	"testing"
	"time"
)

func ExampleDebug() {
	Debug("This is Debug logger.")
}

func ExampleError() {
	Error("This is Debug logger.")
}

func ExampleTmpMark() {
	TmpMark("temp mark when develop.")
	TmpMark("param: %s, Year %d", "Joshua Conero", time.Now().Year())
}

func TestDebug(t *testing.T) {
	ExampleDebug()
	ExampleError()
	ExampleTmpMark()
}
