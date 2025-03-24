package lgr

import (
	"testing"
	"time"
)

func ExampleDebug() {
	Debug("This is Debug logger.")
	TmpMark("temp mark when develop.")
	TmpMark("param: %s, Year %d", "Joshua Conero", time.Now().Year())
}

func ExampleError() {
	Error("This is Debug logger.")
}

func TestDebug(t *testing.T) {
	ExampleDebug()
	ExampleError()
}
