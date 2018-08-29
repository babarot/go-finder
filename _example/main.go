package main

import (
	"fmt"

	finder "github.com/b4b4r07/go-finder"
	"github.com/b4b4r07/go-finder/source"
)

func main() {
	fzf, err := finder.New("fzf", "--reverse", "--height=50%")
	if err != nil {
		panic(err)
	}
	fmt.Printf("fzf obeject:   %#v\n", fzf)

	// Set data source
	fzf.Read(source.Dir(".", true))
	fzf.Read(source.Slice([]string{"a", "b", "c"}))

	items, err := fzf.Run()
	if err != nil {
		panic(err)
	}
	fmt.Printf("selected items:%#v\n", items)
}
