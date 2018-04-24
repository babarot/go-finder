package finder

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
)

// DefaultCommandsList is the list of finder commands
var DefaultCommandsList = []string{
	"fzf",  // https://github.com/junegunn/fzf
	"peco", // https://github.com/peco/peco
}

// Finder represents the finder command attributes
type Finder struct {
	Command string
	Options []string
	Source  func(io.WriteCloser) error

	oneliner string
	path     string
}

// New returns new Finder object
func New(command string, opts ...string) (*Finder, error) {
	if command == "" {
		return &Finder{}, errors.New("no command available as a CLI finder")
	}
	path, err := exec.LookPath(command)
	if err != nil {
		return &Finder{}, err
	}
	oneliner := path
	for _, opt := range opts {
		oneliner += " " + opt
	}
	return &Finder{
		Options: opts,
		Command: command,
		Source: func(in io.WriteCloser) error {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				fmt.Fprintln(in, scanner.Text())
			}
			return scanner.Err()
		},
		oneliner: oneliner,
		path:     path,
	}, nil
}

// Run runs the finder command
func (f *Finder) Run() ([]string, error) {
	return filter(f.oneliner, f.Source)
}

// SetOptions sets options
func (f *Finder) SetOptions(opts ...string) {
	f.Options = opts
}

// Command returns the command name existing in your PATH
func Command(commands ...string) string {
	if len(commands) == 0 {
		commands = DefaultCommandsList
	}
	for _, command := range commands {
		_, err := exec.LookPath(command)
		if err == nil {
			return command
		}
	}
	return ""
}
