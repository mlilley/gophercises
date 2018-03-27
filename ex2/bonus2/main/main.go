package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	urlshort "github.com/mlilley/gophercises/ex2/bonus2"
)

type args struct {
	filename   string
	configType string
}

func parseArgs() args {
	y := flag.String("yaml", "", "Name of yaml configuration input file")
	j := flag.String("json", "", "Name of json configuration input file")
	flag.Parse()
	if *y == *j {
		log.Fatal("Specify one, and only one, of json or yaml arguments")
	}
	if *y != "" {
		return args{filename: *y, configType: "yaml"}
	}
	return args{filename: *j, configType: "json"}
}

func loadConfig(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

func main() {
	args := parseArgs()
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	config := loadConfig(args.filename)
	var handler http.HandlerFunc
	var err error
	if args.configType == "yaml" {
		handler, err = urlshort.YAMLHandler([]byte(config), mapHandler)
	} else if args.configType == "json" {
		handler, err = urlshort.JSONHandler([]byte(config), mapHandler)
	} else {
		log.Fatal("Unrecognised configuration file type")
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
