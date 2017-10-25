// 采用Trie算法来做路径匹配

package ksana

import (
	"fmt"
	"strings"
)

type trieNode struct {
	path     string
	handle   Handle
	children []*trieNode
}

func (root *trieNode) addRoute(fullPath string, handle Handle) {
	node, path, ok := root.searchNode(fullPath)
	if ok && len(path) == 0 && node.handle != nil {
		panic("该路由已注册")
	}
	for _, pathPart := range path {
		// 不注册空字符串节点
		if pathPart == "" {
			continue
		}
		newNode := new(trieNode)
		newNode.path = pathPart
		node.children = append(node.children, newNode)
		node = newNode
	}
	node.handle = handle
}

func (root *trieNode) getHandle(fullPath string) Handle {
	if n, p, ok := root.searchNode(fullPath); ok && len(p) == 0 {
		return n.handle
	}
	return nil
}

// searchNode 用于搜索trie树, 返回最长匹路径和节点
func (root *trieNode) searchNode(fullPath string) (*trieNode, []string, bool) {
	var DFS func(*trieNode, []string) (*trieNode, []string, bool)
	DFS = func(node *trieNode, path []string) (*trieNode, []string, bool) {
		seg := path[0]
		var rest []string
		if len(path) > 1 {
			rest = path[1:]
		}
		if seg == node.path {
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

// TraversalNode 用于遍所有路由节点
func (root *trieNode) TraversalNode() {
	var recursionDFS func(*trieNode)
	recursionDFS = func(node *trieNode) {
		fmt.Println("path: ", node.path)
		fmt.Println("handle: ", node.handle)
		for _, n := range node.children {
			recursionDFS(n)
		}
	}
	recursionDFS(root)
}
