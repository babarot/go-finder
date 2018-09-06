package source

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// Source is the data source for filter command
type Source func(io.WriteCloser) error

// Text shows the string text splitted by newlines
func Text(text string) Source {
	return func(out io.WriteCloser) error {
		fmt.Fprintln(out, text)
		return nil
	}
}

// Dir scans the directories passed as arguments to enumerate items
func Dir(dir string, full bool) Source {
	return func(out io.WriteCloser) error {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			return err
		}
		for _, file := range files {
			fname := file.Name()
			if full {
				fname = filepath.Join(dir, fname)
			}
			fmt.Fprintln(out, fname)
		}
		return nil
	}
}

// Reader is the source of io.Reader
func Reader(r io.Reader) Source {
	return func(out io.WriteCloser) error {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			fmt.Fprintln(out, scanner.Text())
		}
		return scanner.Err()
	}
}

// Stdin reads the contents from os.Stdin
func Stdin() Source {
	return Reader(os.Stdin)
}

// File shows the file contents splitted by newlines
func File(file string) Source {
	return func(out io.WriteCloser) error {
		fp, err := os.Open(file)
		if err != nil {
			return err
		}
		defer fp.Close()
		scanner := bufio.NewScanner(fp)
		for scanner.Scan() {
			fmt.Fprintln(out, scanner.Text())
		}
		return scanner.Err()
	}
}

// Command reads the execution result of the external command as data source
func Command(command string, args ...string) Source {
	return func(out io.WriteCloser) error {
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
		cmd.Stdout = out
		cmd.Stdin = os.Stdin
		return cmd.Run()
	}
}

// Slice reads the string array as data source
func Slice(s []string) Source {
	return func(out io.WriteCloser) error {
		for _, item := range s {
			fmt.Fprintln(out, item)
		}
		return nil
	}
}
