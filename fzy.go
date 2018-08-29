package finder

// Fzy represents the filter instance
type Fzy struct {
	*Command
}

// Install installs the command
func (c Fzy) Install() error {
	return nil
}
