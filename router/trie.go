// 采用Trie算法来做路径匹配

package router

import (
	"fmt"
	"strings"
)

type trieNode struct {
	path       string
	nodeType   string
	handlerMap map[string]Handler
	children   []*trieNode
}

func (root *trieNode) addRoute(method, fullPath string, handler Handler) {
	node, path, ok := root.searchNode(fullPath, "addRoute")
	fmt.Println(node)
	if ok && len(path) == 0 && node.handlerMap[method] != nil {
		panic("该路由已注册")
	}
	for _, pathPart := range path {
		// 不注册空字符串节点
		if pathPart == "" {
			continue
		}
		newNode := new(trieNode)
		newNode.path = pathPart
		newNode.nodeType = root.parseNodeType(pathPart)
		node.children = append(node.children, newNode)
		node = newNode
	}
	node.handlerMap = make(map[string]Handler)
	node.handlerMap[method] = handler
}

func (root *trieNode) parseNodeType(path string) string {
	switch {
	case strings.HasPrefix(path, ":"):
		return "params"
	default:
		return "string"
	}
}

func (root *trieNode) getHandler(method, fullPath string) (Handler, map[string]string) {
	params := make(map[string]string)
	if n, p, ok := root.searchNode(fullPath, "getHandler", params); ok && len(p) == 0 {
		return n.handlerMap[method], params
	}
	return nil, nil
}

func (root *trieNode) nodeMatch(seg string, node *trieNode, model string) (string, bool) {
	if model == "addRoute" {
		return "string", seg == node.path
	}
	switch node.nodeType {
	case "string":
		return "string", seg == node.path
	case "params":
		return "params", true
	default:
		return "string", false
	}
}

// searchNode 用于搜索trie树, 返回最长匹路径和节点
func (root *trieNode) searchNode(fullPath string, model string, other ...map[string]string) (*trieNode, []string, bool) {
	var DFS func(*trieNode, []string) (*trieNode, []string, bool)
	DFS = func(node *trieNode, path []string) (*trieNode, []string, bool) {
		seg := path[0]
		var rest []string
		if len(path) > 1 {
			rest = path[1:]
		}
		if nodeType, ok := root.nodeMatch(seg, node, model); ok {
			if len(other) > 0 && nodeType == "params" {
				params := other[0]
				params[node.path[1:]] = seg
			}
			if len(rest) == 0 {
				return node, rest, true
			}
			for _, n := range node.children {
				if cn, p, ok := DFS(n, rest); ok {
					return cn, p, ok
				}
			}
			return node, rest, true
		}
		return node, rest, false
	}
	arrPath := strings.Split(fullPath, "/")
	// 将路径首字符为空字符串的替换为斜杠
	if arrPath[0] == "" {
		arrPath[0] = "/"
	}
	// 根路径处理
	// todo 后面还要对路径里的空字符串做处理
	if arrPath[0] == "/" && arrPath[1] == "" {
		arrPath = []string{"/"}
	}
	return DFS(root, arrPath)
}
