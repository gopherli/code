package test

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

type Engine struct{}

func TestServeHttp(t *testing.T) {
	e := new(Engine)
	http.HandleFunc("/", handler)
	log.Println("Listen And Serve :9999")
	log.Fatal(http.ListenAndServe(":9999", e))
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.PATH=%q\n", r.URL.Path)
	case "/hello":
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 Not Found %s\n", r.URL.Path)
	}
}
