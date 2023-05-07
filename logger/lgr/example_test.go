package lgr

import "testing"

func ExampleDebug() {
	Debug("This is Debug logger.")
}

func ExampleError() {
	Error("This is Debug logger.")
}

func TestDebug(t *testing.T) {
	ExampleDebug()
	ExampleError()
}
