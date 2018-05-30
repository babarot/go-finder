package finder

import (
	"io"
	"os"
	"os/exec"
	"strings"
)

func filter(command string, source func(out io.WriteCloser) error) ([]string, error) {
	shell := os.Getenv("SHELL")
	if len(shell) == 0 {
		shell = "sh"
	}
	cmd := exec.Command(shell, "-c", command)
	cmd.Stderr = os.Stderr
	in, _ := cmd.StdinPipe()
	errCh := make(chan error, 1)
	go func() {
		if err := source(in); err != nil {
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
