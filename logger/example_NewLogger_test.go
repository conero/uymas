package logger

func Example() {
	lg := NewLogger(Config{})
	lg.Debugf("it's a debug level message")
	lg.Errorf("it's a error level message")
	lg.Warnf("it's a warn level message")
	lg.Errorf("it's a error level message")
}
