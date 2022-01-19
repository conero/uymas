package logger

import "testing"

func ExampleNewLogger() {
	lg := NewLogger(Config{Level: "error"})
	lg.Debugf("it's a debug level message")
	lg.Warnf("it's a warn level message")
	lg.Errorf("it's a error level message")

	// Output:
	// 10:19:38 [ERROR] it's a error level message
}

func TestNewLogger_ExampleNewLogger(t *testing.T) {
	lg := NewLogger(Config{Level: "error"})
	lg.Debugf("it's a debug level message")
	lg.Warnf("it's a warn level message")
	lg.Errorf("it's a error level message")
}
