package cli

// Config command line program configuration item
type Config struct {
	Title      string
	Head       string
	ArgsConfig *ArgsConfig
	// Disable the default help information
	DisableHelp bool
	// In base authentication, the option does not update the option configuration for authentication
	DisableVerify bool
}

// DefaultConfig default command line configuration
var DefaultConfig = Config{
	Title: "uymas",
}
