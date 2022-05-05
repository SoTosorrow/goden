package goden

import (
	"log"
)

type Router struct {
	routes map[string]HandlerFunc
	root *trieNode
}

func newRouter() *Router {
	return &Router{
		routes: make(map[string]HandlerFunc),
		root: newTrieNode(),
	}
}

func (router *Router) addRoute(method string, routeUrl string, routeFunc HandlerFunc) {
	log.Printf("[Router]: RouteAdd '%s - %s'", method, routeUrl)
	key := method + "-" + routeUrl
	
	values := string2strs(routeUrl)

	router.root.insertStrict(values)
	router.routes[key] = routeFunc
}

func (router *Router) getRoute(method string, routeUrl string) ([]*trieNode,map[string]string) {
	values := string2strs(routeUrl)
	// nodes := router.root.searchStrict(values)
	// 允许模糊匹配
	nodes,params := router.root.search(values)
	// fmt.Println(params)
	if nodes != nil {
		return nodes,params
	}
	return nil,nil
}

