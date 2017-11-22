// 对于框架路由器而言
// 需要根据请求 URL、Method、Regexp 来匹配对应的 Handler

package router

import (
	"Ksana/controller"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Handler 处理器
type Handler = controller.Handler

// Context 请求上下文
type Context = controller.Context

// Router 路由器
type Router struct {
	trie *trieNode
}

func (r *Router) registerRoute(method, fullPath string, handler Handler) {
	// 先判断请求路径是否合法
	if fullPath[0] != '/' {
		panic("path must begin with '/' in path " + fullPath)
	}
	if r.trie == nil {
		r.trie = new(trieNode)
		r.trie.path = "/"
		r.trie.nodeType = "string"
	}
	r.trie.addRoute(method, fullPath, handler)
}

// Get _
func (r *Router) Get(fullPath string, handler Handler) {
	r.registerRoute("GET", fullPath, handler)
}

// Post _
func (r *Router) Post(fullPath string, handler Handler) {
	r.registerRoute("POST", fullPath, handler)
}

// Put _
func (r *Router) Put(fullPath string, handler Handler) {
	r.registerRoute("PUT", fullPath, handler)
}

// Delete _
func (r *Router) Delete(fullPath string, handler Handler) {
	r.registerRoute("DELETE", fullPath, handler)
}

// TraversalNode 用于遍所有路由节点
func (r *Router) TraversalNode() {
	var recursionDFS func(*trieNode)
	recursionDFS = func(node *trieNode) {
		fmt.Println("path: " + node.path + " nodeType: " + node.nodeType)
		for _, n := range node.children {
			recursionDFS(n)
		}
	}
	recursionDFS(r.trie)
}

// 实现 Handler 接口
func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method
	fmt.Println(method, path)
	// todo 拿到请求路径和方法后塞给路由器处理
	if handler, params := r.trie.getHandler(method, path); handler != nil {
		ctx := Context{Req: req, Res: res, Params: params}
		if method == "POST" {
			body, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal("Parse Post Request Body: ", err)
			}
			ctx.Body = body
		}
		ctx.Res.Header().Set("Content-Type", "application/json; charset=utf-8")
		handler(ctx)
	} else {
		fmt.Fprintf(res, "404")
	}
}
