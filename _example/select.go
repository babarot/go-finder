package main

import (
	"fmt"

	finder "github.com/b4b4r07/go-finder"
)

func main() {
	fzf, err := finder.New("fzf")
	if err != nil {
		panic(err)
	}

	type Book struct {
		Title string
		ISBN  string
	}

	books := []Book{
		Book{
			Title: "Book A",
			ISBN:  "aaa",
		},
		Book{
			Title: "Book B",
			ISBN:  "bbb",
		},
		Book{
			Title: "Book C",
			ISBN:  "ccc",
		},
	}
	for _, book := range books {
		fzf.Add(book.Title, book)
	}
	items, err := fzf.Select()
	if err != nil {
		panic(err)
	}
	for _, item := range items {
		fmt.Printf("ISBN of %s is %s\n", item.(Book).Title, item.(Book).ISBN)
	}
}
