package main

import (
	"fmt"
	"net/http"
	"github.com/golang/glog"
)

func init() {
	http.HandleFunc("/hello", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
	glog.Fatal("Hello, world! log")
}
