package main

import (
	"fmt"
	"net/http"
	"io"
	"strconv"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		log.Debugf(ctx, "Content-Length error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := make([]byte, length)
	log.Debugf(ctx, strconv.Itoa(length))
	length, err2 := r.Body.Read(body)
	if err2 != nil && err2 != io.EOF {
		log.Debugf(ctx, "r.Body.Read error: " + err2.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Debugf(ctx, "request URL: " + r.URL.Path)
	json := string(body)
	log.Debugf(ctx, "r.Body.Read " + json)
	fmt.Fprint(w, "Hello, world!")
}
