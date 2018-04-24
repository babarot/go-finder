package main

import (
	"log"

	finder "github.com/b4b4r07/go-finder"
	"github.com/k0kubun/pp"
)

func main() {
	opts := []string{}
	f, err := finder.New(finder.Command(), opts...)
	if err != nil {
		log.Fatal(err)
	}
	pp.Println(f.Run())
}
