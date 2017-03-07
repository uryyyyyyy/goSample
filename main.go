package main

import (
	"fmt"
	"net/http"
	"github.com/franela/goreq"
)

func init() {
	http.HandleFunc("/hello", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
	res, _ := goreq.Request{ Uri: "http://www.google.com" }.Do()
	fmt.Fprint(w, "Hello, world!" + res.Status)
}
