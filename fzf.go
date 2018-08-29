package finder

// Fzf represents the filter instance
type Fzf struct {
	*Command
}

// Install installs the command
func (c Fzf) Install() error {
	return nil
}
