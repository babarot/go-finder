package main

import (
	"fmt"

	finder "github.com/b4b4r07/go-finder"
)

func main() {
	var opts []string

	command := finder.Command()
	switch command {
	case "fzf":
		opts = []string{
			"--reverse",
			"--height", "40",
		}
	case "peco":
		opts = []string{
			"--layout=bottom-up",
		}
	}

	cli, err := finder.New(command, opts...)
	if err != nil {
		panic(err)
	}

	// You can select the data source to use as filter source
	cli.FromFile("some-file.txt")
	cli.FromText("sample\ntext\nfoo")
	cli.FromCommand("cat", "some-file.txt")
	cli.FromStdin() // default

	items, err := cli.Run()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", items)
}
