package main

import (
	"Ksana/controller"
	"Ksana/router"
	"fmt"
	"log"
	"net/http"
)

// Handler 带装饰的处理器
var Handler = controller.Handler

func main() {
	app := new(router.Router)

	app.Get("/", func(ctx controller.Context) {
		fmt.Fprintf(ctx.Res, "Hello World")
	})

	app.Post("/login", controller.Login)

	app.Get("/authorizationCheck", controller.AuthorizationCheck)

	app.Post("/post/add", Handler(controller.AddPost, controller.AuthorCheck))

	app.Post("/post/update/:pID", Handler(controller.UpdatePost, controller.AuthorCheck))

	app.Get("/post/get/:pID", controller.GetPost)

	app.Get("/tags/get", controller.GetTags)

	app.Get("/post/getByTag/:tag", controller.GetPostsByTag)

	if err := http.ListenAndServe(":9090", app); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
