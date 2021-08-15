package main

import (
	"fmt"
	"net/http"
)

type Engine struct {
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.PAth = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND")

	}
}
func main() {
	r := gee.New()
	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.PAth = %q\n", req.URL.Path)

	})
	r.Get("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	r.Run(":9999")
}
