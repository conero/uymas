package evolve

// CommandRefer the interface matches the registration command
type CommandRefer interface {
	// Init command initialization
	Init()
	// DefIndex entry definition
	DefIndex()
	// DefHelp help/reference command definition
	DefHelp()
	// DefLost No command definition exists
	DefLost()
}

// Command line basic (minimum) structure standard
type Command struct {
	X *Param
}
