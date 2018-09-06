package finder

// Fzy represents the filter instance
type Fzy struct {
	*Command
}

// Install installs the command
func (c Fzy) Install(path string) error {
	// not support yet
	return nil
}
