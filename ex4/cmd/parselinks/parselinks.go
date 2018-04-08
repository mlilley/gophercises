package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mlilley/gophercises/ex4"
)

func main() {
	filename := flag.String("f", "", "Name of html input file")
	flag.Parse()

	r, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}

	links, err := link.Parse(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", links)
}
