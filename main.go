package main

import (
	"log"
	"net/http"
)

type Router struct{}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Method: %s | URL: %s | Header: %s\n", r.Method, r.URL.Path, r.Header)
	http.NotFound(w, r)
}

func main() {
	http.ListenAndServe(":3900", &Router{})
}
