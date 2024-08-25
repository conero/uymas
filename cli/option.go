package cli

// Option Used for command option parsing document generation, or value validation and retrieval
type Option struct {
	Name     string
	Alias    []string
	Require  bool
	DefValue string
	Help     string
}

// CommandOptional Used for command registration as a parameter option
type CommandOptional struct {
	Help    string
	Alias   []string
	Options []Option
}
