// 对于框架路由器而言
// 需要根据请求 URL、Method、Regexp 来匹配对应的 Handler

package ksana

import (
	"fmt"
	"net/http"
)

// Handle 处理器
type Handle func(http.ResponseWriter, *http.Request)

// Router 路由器
type Router struct {
	Tries map[string]*trieNode
}

func (r *Router) registerRoute(method, fullPath string, handle Handle) {
	// 先判断请求路径是否合法
	if fullPath[0] != '/' {
		panic("path must begin with '/' in path" + fullPath)
	}
	if r.Tries == nil {
		r.Tries = make(map[string]*trieNode)
	}
	if root := r.Tries[method]; root == nil {
		root = new(trieNode)
		root.path = "/"
		r.Tries[method] = root
	}
	r.Tries[method].addRoute(fullPath, handle)
}

func (r *Router) GET(fullPath string, handle Handle) {
	r.registerRoute("GET", fullPath, handle)
}

func (r *Router) POST(fullPath string, handle Handle) {
	r.registerRoute("POST", fullPath, handle)
}

func (r *Router) PUT(fullPath string, handle Handle) {
	r.registerRoute("PUT", fullPath, handle)
}

func (r *Router) DELETE(fullPath string, handle Handle) {
	r.registerRoute("DELETE", fullPath, handle)
}

func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method
	fmt.Println(method, path)
	// todo 拿到请求路径和方法后塞给路由器处理
	if handle := r.Tries[method].getHandle(path); handle != nil {
		handle(res, req)
	} else {
		fmt.Fprintf(res, "404")
	}
}
