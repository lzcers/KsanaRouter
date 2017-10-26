// 对于框架路由器而言
// 需要根据请求 URL、Method、Regexp 来匹配对应的 Handler

package router

import (
	"fmt"
	"net/http"
)

// Handle 处理器
type handle func(http.ResponseWriter, *http.Request)

// Router 路由器
type Router struct {
	trie *trieNode
}

func (r *Router) registerRoute(method, fullPath string, handleFunc handle) {
	// 先判断请求路径是否合法
	if fullPath[0] != '/' {
		panic("path must begin with '/' in path" + fullPath)
	}
	if r.trie == nil {
		r.trie = new(trieNode)
		r.trie.path = "/"
	}
	r.trie.addRoute(method, fullPath, handleFunc)
}

func (r *Router) Get(fullPath string, handleFunc handle) {
	r.registerRoute("GET", fullPath, handleFunc)
}

func (r *Router) Post(fullPath string, handleFunc handle) {
	r.registerRoute("POST", fullPath, handleFunc)
}

func (r *Router) Put(fullPath string, handleFunc handle) {
	r.registerRoute("PUT", fullPath, handleFunc)
}

func (r *Router) Delete(fullPath string, handleFunc handle) {
	r.registerRoute("DELETE", fullPath, handleFunc)
}

// TraversalNode 用于遍所有路由节点
func (r *Router) TraversalNode() {
	var recursionDFS func(*trieNode)
	recursionDFS = func(node *trieNode) {
		fmt.Println("path: ", node.path)
		fmt.Println("handle: ", node.handleMap)
		for _, n := range node.children {
			recursionDFS(n)
		}
	}
	recursionDFS(r.trie)
}

func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method
	fmt.Println(method, path)
	// todo 拿到请求路径和方法后塞给路由器处理
	if handle := r.trie.getHandle(method, path); handle != nil {
		handle(res, req)
	} else {
		fmt.Fprintf(res, "404")
	}
}
