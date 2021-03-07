package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {	
	e := gee.New()
	e.GET("/ss", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "URL.PATH=%q\n", r.URL.Path)
	})

	e.POST("/hello", func(rw http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(rw, "HEADER[%q]=%q", k, v)
		}
	})

	e.Run(":9999")
}
