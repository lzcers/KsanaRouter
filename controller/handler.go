package controller

import "net/http"

// Handler 处理器
type Handler func(Context)

// Context 请求上下文
type Context struct {
	Req    *http.Request
	Res    http.ResponseWriter
	Params map[string]string
	Body   []byte
}
