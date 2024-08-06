package cli

// Config command line program configuration item
type Config struct {
	Title      string
	ArgsConfig *ArgsConfig
}

// DefaultConfig default command line configuration
var DefaultConfig = Config{
	Title: "uymas",
}
