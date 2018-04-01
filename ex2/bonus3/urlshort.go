package urlshort

import (
	"fmt"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v, p := pathsToUrls[r.RequestURI]
		if p {
			fmt.Printf("%v -> %v\n", r.RequestURI, v)
			http.Redirect(w, r, v, 307) // temporary
		} else {
			fmt.Printf("%v -> *unrecognised*\n", r.RequestURI)
			fallback.ServeHTTP(w, r)
		}
	})
	return h
}
