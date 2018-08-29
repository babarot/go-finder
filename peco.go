package finder

// Peco represents the filter instance
type Peco struct {
	*Command
}

// Install installs the command
func (c Peco) Install() error {
	return nil
}
