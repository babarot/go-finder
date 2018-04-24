package main

import (
	"fmt"

	finder "github.com/b4b4r07/go-finder"
)

func main() {
	peco, err := finder.New("peco", "--layout=bottom-up")
	if err != nil {
		panic(err)
	}
	items, err := peco.Run()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", items)
}
