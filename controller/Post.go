package controller

import (
	"Ksana/models"
	"encoding/json"
	"fmt"
)

func AddPost(ctx Context) {
	var post models.Post
	json.Unmarshal(ctx.Body, &post)
	models.AddPost(post)
}

func GetPost(ctx Context) {
	pID := ctx.Params["pID"]
	posts := models.GetPost(pID)
	postsJSON, err := json.Marshal(posts)
	if err != nil {
		// 我他妈也不知道该做啥
	}
	fmt.Fprintf(ctx.Res, string(postsJSON))
}

func GetTags(ctx Context) {
	tagsJSON, err := json.Marshal(models.GetTags())
	if err != nil {
		// 我他妈也不知道该做啥
	}
	fmt.Fprint(ctx.Res, string(tagsJSON))
}
