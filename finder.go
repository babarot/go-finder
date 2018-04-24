package finder

import (
	"errors"
	"io"
	"os/exec"
)

var DefaultCommandsList = []string{
	"fzf",  // https://github.com/junegunn/fzf
	"peco", // https://github.com/peco/peco
}

type Finder struct {
	Command string
	Options []string
	Input   func(io.WriteCloser)

	oneliner string
	path     string
}

func New(command string, opts ...string) (*Finder, error) {
	if command == "" {
		return &Finder{}, errors.New("no command available as a CLI filter")
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
		Options:  opts,
		Command:  command,
		oneliner: oneliner,
		path:     path,
	}, nil
}

func (f *Finder) Run() []string {
	return withFilter(f.oneliner, f.Input)
}

func Command() string {
	for _, command := range DefaultCommandsList {
		_, err := exec.LookPath(command)
		if err == nil {
			return command
		}
	}
	return ""
}
