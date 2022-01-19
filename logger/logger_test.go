package logger

import (
	"fmt"
	"testing"
)

func TestNewLogger(t *testing.T) {
	lg := NewLogger(Config{
		Level: "info",
	})
	lg.Debugf("it's a debug level message") // it'll print nothing.
	lg.Infof("it's a debug level message")
	lg.Warnf("it's a warn level message")
	lg.Errorf("it's a error level message")
}

func TestNewLogger_buffer(t *testing.T) {
	lg := NewLogger(Config{
		Level:  "error",
		Driver: DriverBuffer,
	})

	buff := lg.Buffer()
	var outputString string
	if buff == nil {
		t.Errorf("buffer dirver init failure.")
	}
	lg.Debugf("it's a debug level message") // it'll print nothing.
	if fmt.Sprint(buff) != "" {
		t.Errorf("level (debug) control is failure, output: %v", fmt.Sprint(buff))
	}
	lg.Infof("it's a debug level message")

	if fmt.Sprint(buff) != "" {
		t.Errorf("level (info) control is failure")
	}
	lg.Warnf("it's a warn level message")
	if fmt.Sprint(buff) != "" {
		t.Errorf("level (warn) control is failure")
	}
	lg.Errorf("it's a error level message")
	outputString = fmt.Sprint(buff)
	if outputString == "" {
		t.Errorf("level (debug) control is failure")
	} else {
		t.Log(outputString)
	}
}
