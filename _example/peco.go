package main

import (
	"log"

	finder "github.com/b4b4r07/go-finder"
	"github.com/k0kubun/pp"
)

func main() {
	peco, err := finder.New("peco", "--layout=bottom-up")
	if err != nil {
		log.Fatal(err)
	}
	pp.Println(peco.Run())
}
