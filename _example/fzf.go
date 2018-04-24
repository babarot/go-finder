package main

import (
	"fmt"

	finder "github.com/b4b4r07/go-finder"
)

func main() {
	fzf, err := finder.New("fzf", "--reverse", "--height", "40")
	if err != nil {
		panic(err)
	}

	items, err := fzf.Run()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", items)
}
