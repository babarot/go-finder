package main

import (
	"fmt"

	finder "github.com/b4b4r07/go-finder"
	"github.com/b4b4r07/go-finder/source"
)

func main() {
	fzf, err := finder.New("fzf")
	if err != nil {
		panic(err)
	}
	fmt.Printf("fzf obeject:   %#v\n", fzf)

	// Read files list within dir as data source of fzf
	fzf.Read(source.Dir(".", true))

	items, err := fzf.Run()
	if err != nil {
		panic(err)
	}
	fmt.Printf("selected items:%#v\n", items)
}
