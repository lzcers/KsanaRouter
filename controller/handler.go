package controller

import "net/http"

// Context 请求上下文
type Context struct {
	Req    *http.Request
	Res    http.ResponseWriter
	Params map[string]string
	Body   []byte
}

// HandlerFunc 处理器
type HandlerFunc func(Context)

// HandlerDecorator 装饰器
type HandlerDecorator func(HandlerFunc) HandlerFunc

// Handler 启用装饰器修饰 Handler
func Handler(h HandlerFunc, decors ...HandlerDecorator) HandlerFunc {
	dLen := len(decors)
	for i := range decors {
		d := decors[dLen-i-1]
		h = d(h)
	}
	return h
}
