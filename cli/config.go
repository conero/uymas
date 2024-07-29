package cli

// Config command line program configuration item
type Config struct {
	Title string
}

// DefaultConfig default command line configuration
var DefaultConfig = Config{
	Title: "uymas",
}
