package gen

// Value Customized values for command-line options, supporting obtaining values and determining whether to set them. Used for special processing
type Value[V any] struct {
	Data V
	// Determine whether the option has been set
	IsSet bool
}
