package logger

import "log"

// Config builder logging configure.
type Config struct {
	Log   *log.Logger
	Level string
	// driver list: stdout, file, buffer
	Driver string
	// when Driver is file it will be work
	OutputDir string
}

var DefaultConfig = Config{
	Level:  "warn",
	Driver: DriverStdout,
}
