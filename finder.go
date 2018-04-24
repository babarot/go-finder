package finder

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
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

// From sets Source
func (f *Finder) From(source func(io.WriteCloser) error) {
	f.Source = source
}

// FromFile sets the contents of the file as Source
func (f *Finder) FromFile(file string) {
	f.Source = func(in io.WriteCloser) error {
		fp, err := os.Open(file)
		if err != nil {
			return err
		}
		defer fp.Close()
		scanner := bufio.NewScanner(fp)
		for scanner.Scan() {
			fmt.Fprintln(in, scanner.Text())
		}
		return scanner.Err()
	}
}

// FromText sets the text as Source
func (f *Finder) FromText(text string) {
	f.Source = func(in io.WriteCloser) error {
		fmt.Fprintln(in, text)
		return nil
	}
}

// FromCommand sets the execution result of the command as Source
func (f *Finder) FromCommand(command string, args ...string) {
	f.Source = func(in io.WriteCloser) error {
		if _, err := exec.LookPath(command); err != nil {
			return err
		}
		for _, arg := range args {
			command += " " + arg
		}
		var cmd *exec.Cmd
		if runtime.GOOS == "windows" {
			cmd = exec.Command("cmd", "/c", command)
		} else {
			cmd = exec.Command("sh", "-c", command)
		}
		cmd.Stderr = os.Stderr
		cmd.Stdout = in
		cmd.Stdin = os.Stdin
		return cmd.Run()
	}
}

// FromReader sets io.Reader as Source
func (f *Finder) FromReader(r io.Reader) {
	f.Source = func(in io.WriteCloser) error {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			fmt.Fprintln(in, scanner.Text())
		}
		return scanner.Err()
	}
}

// FromStdin sets os.Stdin as Source
func (f *Finder) FromStdin() {
	f.FromReader(os.Stdin)
}
