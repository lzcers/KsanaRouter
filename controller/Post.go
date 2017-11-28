package controller

import (
	"Ksana/models"
	"encoding/json"
	"fmt"
)

func AddPost(ctx Context) {
	if authorCheck(ctx) {
		var post models.Post
		json.Unmarshal(ctx.Body, &post)
		pID := models.AddPost(post)
		fmt.Fprintf(ctx.Res, pID)
	} else {
		ctx.Res.WriteHeader(401)
		fmt.Fprintf(ctx.Res, `{"result": false}`)
	}
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
	fmt.Fprintf(ctx.Res, string(tagsJSON))
}

func GetPostsByTag(ctx Context) {
	tag := ctx.Params["tag"]
	postsJSON, err := json.Marshal(models.GetPostsByTag(tag))
	if err != nil {

	}
	fmt.Fprintf(ctx.Res, string(postsJSON))
}

func UpdatePost(ctx Context) {
	if authorCheck(ctx) {
		var post models.Post
		pID := ctx.Params["pID"]
		json.Unmarshal(ctx.Body, &post)
		models.UpdatePost(pID, post)
	} else {
		ctx.Res.WriteHeader(401)
		fmt.Fprintf(ctx.Res, `{"result": false}`)
	}
}
