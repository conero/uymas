package logger

import "testing"

func TestNewLogger(t *testing.T) {
	lg := NewLogger(Config{
		Level: "info",
	})
	lg.Debugf("it's a debug level message") // it'll print nothing.
	lg.Infof("it's a debug level message")
	lg.Warnf("it's a warn level message")
	lg.Errorf("it's a error level message")
}
