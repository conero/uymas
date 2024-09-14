package cli

import (
	"fmt"
	"gitee.com/conero/uymas/v2"
	"gitee.com/conero/uymas/v2/util/fs"
)

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
	DisableVerify: false,
	DisableHelp:   false,
}

// IndexDoc Generate a default entry title
func (c Config) IndexDoc() {
	if c.Head != "" {
		title := c.Title
		if title == "" {
			title = fs.AppName()
		}
		fmt.Println()
		fmt.Println("-------------- " + title + " -----------------")
		fmt.Println(c.Head)
		fmt.Println()
		return
	}

	fmt.Println()
	fmt.Println("-------------- Uymas -----------------")
	fmt.Println()
	fmt.Println("Welcome to our world")
	fmt.Printf(":)- %s/%s\n", uymas.Version, uymas.Release)
	fmt.Println()
}
