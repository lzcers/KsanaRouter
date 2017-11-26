package controller

import (
	"Ksana/models"
	"Ksana/session"
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
)

func passwordToMD5(paswd string) string {
	md5Password := md5.New()
	io.WriteString(md5Password, paswd)
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "%x", md5Password.Sum(nil))
	newPass := buffer.String()
	return newPass
}

// Login 登陆校验
func Login(ctx Context) {
	userinfo := new(struct {
		Username string `json:"username"`
		Password string `json:"password"`
	})
	err := json.Unmarshal(ctx.Body, userinfo)
	if err == nil {
		newPass := passwordToMD5(userinfo.Password)
		uInfo := models.GetUser(userinfo.Username)
		if len(uInfo) != 0 && uInfo[0].Password == newPass {
			// 登陆成功， 分配 session ，并设置该 session 的 username 值为当前登陆用户
			sess := session.GlobalSessions.SessionStart(ctx.Res, ctx.Req)
			sess.Set("username", userinfo.Username)
			fmt.Fprintf(ctx.Res, `{"result": true}`)
		} else {
			fmt.Fprintf(ctx.Res, `{"result": false}`)
		}
	}
}

// AuthorizationCheck 权限校验
func AuthorizationCheck(ctx Context) {
	sess := session.GlobalSessions.SessionStart(ctx.Res, ctx.Req)
	sessUserName := sess.Get("username")

	if sessUserName != nil && sessUserName.(string) == "admin" {
		fmt.Fprintf(ctx.Res, `{"result": true}`)
		return
	}
	// 其他一律提示未授权
	fmt.Fprintf(ctx.Res, `{"result": false}`)
}
