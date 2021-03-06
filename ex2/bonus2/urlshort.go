package urlshort

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

type configItem struct {
	Path string
	Url  string
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if url, ok := pathsToUrls[r.RequestURI]; ok {
			fmt.Printf("%v -> %v\n", r.RequestURI, url)
			http.Redirect(w, r, url, 307) // temporary
			return
		}
		fmt.Printf("%v -> *unrecognised*\n", r.RequestURI)
		fallback.ServeHTTP(w, r)
	})
	return h
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var items []configItem

	err := yaml.Unmarshal(yml, &items)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	for i := 0; i < len(items); i++ {
		m[items[i].Path] = items[i].Url
	}

	return MapHandler(m, fallback), nil
}

// JSONHandler parses the provided json and returns an http.HandlerFunc.
// Expected format: [{"Path":"/foo,"Url":"http://bar.com"},{...}]
func JSONHandler(content []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var items []configItem

	err := json.Unmarshal(content, &items)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	for i := 0; i < len(items); i++ {
		m[items[i].Path] = items[i].Url
	}

	return MapHandler(m, fallback), nil
}
