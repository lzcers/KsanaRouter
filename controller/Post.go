package controller

import (
	"Ksana/models"
	"Ksana/session"
	"encoding/json"
	"fmt"
)

func AddPost(ctx Context) {
	var post models.Post
	json.Unmarshal(ctx.Body, &post)
	pID := models.AddPost(post)
	fmt.Fprintf(ctx.Res, pID)
}

func GetPost(ctx Context) {
	pID := ctx.Params["pID"]
	// 验证身份，非管理员只能读取公开的文章
	sess := session.GlobalSessions.SessionStart(ctx.Res, ctx.Req)
	sessUserName := sess.Get("username")
	var posts []models.Post
	if sessUserName != nil && sessUserName.(string) == "admin" {
		posts = models.GetPost(pID, "admin")
	} else {
		posts = models.GetPost(pID, "user")
	}
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
	var post models.Post
	pID := ctx.Params["pID"]
	json.Unmarshal(ctx.Body, &post)
	models.UpdatePost(pID, post)
}
