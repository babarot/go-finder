package main

import (
	"log"

	finder "github.com/b4b4r07/go-finder"
	"github.com/k0kubun/pp"
)

func main() {
	fzf, err := finder.New("fzf", "--reverse", "--height", "40")
	if err != nil {
		log.Fatal(err)
	}
	pp.Println(fzf.Run())
}
