package main

import (
	"Ksana/router"
	"fmt"
	"log"
	"net/http"
)

type context = router.Context

func main() {
	app := new(router.Router)

	app.Get("/post/:fileName/xxxx", func(ctx context) {
		fmt.Fprintf(ctx.Res, ctx.Params["fileName"])
	})

	if err := http.ListenAndServe(":9090", app); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
