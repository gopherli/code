package test

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestNetHttp(t *testing.T) {
	http.HandleFunc("/", handler)
	log.Println("Listen And Serve localhost:8080 ")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
