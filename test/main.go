package main

import (
	"fmt"
	"ksana"
	"net/http"
)

func main() {
	router := new(ksana.Router)
	router.GET("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello myroute! /")
	})
	router.GET("/aa/b/c/d", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello myroute! 1")
	})
	router.GET("/aa/b/c/e", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello myroute! 2")
	})
	router.GET("/aa/b/x/1", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello myroute! 3")
	})
	router.GET("/aa/f", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello myroute! 4")
	})
	http.ListenAndServe(":9090", router)
	// ksana.Run("localhost:8080")
}
