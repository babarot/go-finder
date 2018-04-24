package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"

	finder "github.com/b4b4r07/go-finder"
	"github.com/k0kubun/pp"
)

func main() {
	var opts []string
	command := finder.Command()
	switch command {
	case "fzf":
		opts = []string{"--reverse"}
	case "peco":
	}
	f, err := finder.New(command, opts...)
	if err != nil {
		log.Fatal(err)
	}

	// sample
	f.Input = func(in io.WriteCloser) {
		for i := 0; i < 1000; i++ {
			fmt.Fprintln(in, i)
		}
	}

	// file
	file, err := os.Open("test")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	f.Input = func(in io.WriteCloser) {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Fprintln(in, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	// echo
	f.Input = func(in io.WriteCloser) {
		command := "echo hoge"
		var cmd *exec.Cmd
		if runtime.GOOS == "windows" {
			cmd = exec.Command("cmd", "/c", command)
		} else {
			cmd = exec.Command("sh", "-c", command)
		}
		cmd.Stderr = os.Stderr
		cmd.Stdout = in
		cmd.Stdin = os.Stdin
		cmd.Run()
	}

	// stdin
	f.Input = func(in io.WriteCloser) {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Fprintln(in, scanner.Text())
		}
		if scanner.Err() != nil {
			log.Fatal(scanner.Err())
		}
	}

	pp.Println(f.Run())
}
