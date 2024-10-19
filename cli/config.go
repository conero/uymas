package cli

import (
	"fmt"
	"gitee.com/conero/uymas/v2"
	"gitee.com/conero/uymas/v2/rock"
	"gitee.com/conero/uymas/v2/util/fs"
)

var (
	gConfig *Config
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
	// StructGenSep field Struct generation separator
	StructGenSep string
}

// DefaultConfig default command line configuration
var DefaultConfig = Config{
	DisableVerify: false,
	DisableHelp:   false,
	StructGenSep:  ":",
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

// ConfigWith read config with cache
func ConfigWith(isForces ...bool) Config {
	isForce := rock.Param(false, isForces...)
	if gConfig == nil || isForce {
		gConfig = &DefaultConfig
	}
	return *gConfig
}

// ConfigSet set config to cache
func ConfigSet(cfg Config) {
	gConfig = &cfg
}
