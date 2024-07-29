// Package cli basic command line definition and processing tools
package cli

// Cli command line struct
type Cli struct {
	config Config
}

// Run execute command line as a program entry
func (c *Cli) Run(args ...string) error {
	return nil
}

// NewCli the command line program is instantiated and the driver is as light as possible
func NewCli(cfgs ...Config) *Cli {
	config := DefaultConfig
	if len(cfgs) > 0 {
		config = cfgs[0]
	}
	return &Cli{
		config: config,
	}
}
