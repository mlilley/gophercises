package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/mlilley/gophercises/ex3/p1"
)

func loadStory() (p1.Story, error) {
	filename := flag.String("file", "gopher.json", "The Story JOSN file")
	flag.Parse()

	r, err := os.Open(*filename)
	if err != nil {
		return nil, err
	}

	story, err := p1.ParseStory(r)
	if err != nil {
		return nil, err
	}

	return story, nil
}

func main() {
	story, err := loadStory()
	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(":8080", p1.StoryHandler(&story))
}
