package finder

import (
	"os"
	"os/exec"
	"strings"

	"github.com/b4b4r07/go-finder/source"
	"github.com/pkg/errors"
)

// CLI is the command having a command-line interface
type CLI interface {
	Run() ([]string, error)
}

// Finder is the interface of a filter command
type Finder interface {
	CLI
	Install() error
	Read(source.Source)
}

// Command represents the command
type Command struct {
	Name  string
	Args  []string
	Path  string
	Input source.Source
}

// Run runs as a command
func (c *Command) Run() ([]string, error) {
	shell := os.Getenv("SHELL")
	if len(shell) == 0 {
		shell = "sh"
	}
	cmd := exec.Command(shell, "-c", c.Path+" "+strings.Join(c.Args, " "))
	cmd.Stderr = os.Stderr
	in, _ := cmd.StdinPipe()
	errCh := make(chan error, 1)
	go func() {
		if err := c.Input(in); err != nil {
			errCh <- err
			return
		}
		errCh <- nil
		in.Close()
	}()
	err := <-errCh
	if err != nil {
		return []string{}, err
	}
	result, _ := cmd.Output()
	return trimLastNewline(strings.Split(string(result), "\n")), nil
}

func trimLastNewline(s []string) []string {
	if len(s) == 0 {
		return s
	}
	last := len(s) - 1
	if s[last] == "" {
		return s[:last]
	}
	return s
}

// Install does nothing and is implemented to satisfy Finder interface
// This method should be overwritten by each finder command implementation
func (c *Command) Install() error {
	return nil
}

// Read sets the data sources
func (c *Command) Read(data source.Source) {
	c.Input = data
}

// New creates Finder instance
func New(name string, args ...string) (Finder, error) {
	path, err := exec.LookPath(name)
	if err != nil {
		return nil, errors.Wrapf(err, "%s: not found", name)
	}
	command := &Command{
		Name:  name,
		Args:  args,
		Path:  path,
		Input: source.Stdin(),
	}
	switch name {
	case "fzf":
		return Fzf{command}, nil
	case "fzy":
		return Fzy{command}, nil
	case "peco":
		return Peco{command}, nil
	default:
		return command, nil
	}
}
