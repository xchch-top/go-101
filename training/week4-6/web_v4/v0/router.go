package v0

import "strings"

// router 代表路由
type router struct {
	// trees 代表的是森林, map的key是 http method 代表树的根节点
	trees map[string]*node
}

func newRouter() router {
	return router{
		trees: map[string]*node{},
	}
}

func (r *router) addRoute(method string, path string, handleFunc HandleFunc) {
	root, ok := r.trees[method]
	if !ok {
		// 构造根节点
		root = &node{path: "/"}
		r.trees[method] = root
	}

	if path == "/" {
		root.handler = handleFunc
		return
	}

	cur := root
	path = strings.Trim(path, "/")
	segs := strings.Split(path, "/")
	for _, seg := range segs {
		child := cur.childOrCreate(seg)
		cur = child
	}
	cur.handler = handleFunc
}

// childOrCreate 查找子节点, 如果子节点不存在就创建一个
// 并且将子节点放回去了 children中
func (n *node) childOrCreate(path string) *node {
	// 处理通配符
	if path == "*" {
		if n.starChild == nil {
			n.starChild = &node{
				path: path,
			}
		}
		return n.starChild
	}

	// 处理静态路由
	if n.children == nil {
		n.children = make(map[string]*node)
	}
	child, ok := n.children[path]
	if !ok {
		child = &node{path: path}
		n.children[path] = child
	}
	return child
}

func (r *router) findRoute(method string, path string) (*node, bool) {
	root, ok := r.trees[method]
	if !ok {
		return nil, false
	}
	if path == "/" {
		return root, true
	}

	path = strings.Trim(path, "/")
	cur := root
	segs := strings.Split(path, "/")
	for _, seg := range segs {
		if cur.children == nil {
			if cur.starChild != nil {
				return cur.starChild, cur.starChild.handler != nil
			}
			return nil, false
		}
		child, ok := cur.children[seg]
		if !ok {
			if cur.starChild != nil {
				return cur.starChild, cur.starChild.handler != nil
			}
			return nil, false
		}
		cur = child
	}

	if cur.handler == nil {
		return nil, false
	}
	return cur, true
}

type node struct {
	// /a/b/c 中b这一段
	path     string
	handler  HandleFunc
	children map[string]*node

	starChild *node
}
