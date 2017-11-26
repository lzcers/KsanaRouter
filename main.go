package main

import (
	"Ksana/controller"
	"Ksana/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	app := new(router.Router)
	app.Get("/", func(p controller.Context) {
		fmt.Fprintf(p.Res, "Hello World")
	})

	app.Post("/login", controller.Login)

	app.Get("/authorizationCheck", controller.AuthorizationCheck)

	app.Post("/post/add", controller.AddPost)

	app.Post("/post/update/:pID", controller.UpdatePost)

	app.Get("/post/get/:pID", controller.GetPost)

	app.Get("/tags/get", controller.GetTags)

	app.Get("/post/getByTag/:tag", controller.GetPostsByTag)

	if err := http.ListenAndServe(":9090", app); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
