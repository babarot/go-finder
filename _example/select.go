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
	items := finder.NewItems()
	for _, book := range books {
		items.Add(book.Title, book)
	}
	selectedItems, err := fzf.Select(items)
	if err != nil {
		panic(err)
	}
	for _, item := range selectedItems {
		fmt.Printf("ISBN of %s is %s\n", item.(Book).Title, item.(Book).ISBN)
	}
}
