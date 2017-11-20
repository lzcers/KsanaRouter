package controller

import (
	"Ksana/models"
	"encoding/json"
)

func AddPost(ctx Context) {
	var post models.Post
	json.Unmarshal(ctx.Body, &post)
	models.AddPost(post)
}
