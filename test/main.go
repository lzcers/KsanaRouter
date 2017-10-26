package main

import (
	"Ksana/router"
	"fmt"
	"net/http"
)

func main() {
	router := new(router.Router)
	router.Get("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello myroute! /")
	})
	router.Get("/aa/b/c/d", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello myroute! 1")
	})
	router.Get("/aa/b/c/e", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello myroute! 2")
	})
	router.Get("/aa/b/x/1", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello myroute! 3")
	})
	router.Get("/aa/f", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello myroute! 4")
	})
	router.TraversalNode()
	http.ListenAndServe(":9090", router)
	// ksana.Run("localhost:8080")
}
