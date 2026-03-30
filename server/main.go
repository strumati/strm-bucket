package main

import (
	"log"
	"net/http"

	"strumati.cloud/bucket/lib"
)

type Router struct{}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Printf("Method: %s | URL: %s | Header: %s\n", r.Method, r.URL.Path, r.Header)

	w.Header().Set("Conent-Type", "application/xml")

	lib.LS(w)
}

func main() {
	http.ListenAndServe(":3901", &Router{})
}
