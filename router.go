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

func (router *Router) addRoute(method string, routeUrl string, routeFunc HandlerFunc) *RouterBranch{
	if routeUrl == "" {
		routeUrl = "/"
	}
	if string(routeUrl[0]) != "/" {
		routeUrl = "/" + routeUrl
	}
	log.Printf("[Router]: RouteAdd '%s - %s'", method, routeUrl)
	key := method + "-" + routeUrl
	
	values := string2strs(routeUrl)

	node := router.root.insertStrict(values)
	router.routes[key] = routeFunc

	return &RouterBranch{
		Router: router, 		
		branchUrl: routeUrl,
		root: node,
	}
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

//TODO node2 = web.addRoute("url1",...) node2.addRoute("url2",...)
type RouterBranch struct {
	Router 		*Router
	branchUrl 	string
	root 		*trieNode
}

func (branch *RouterBranch) AddGetRoute(routeUrl string, routeFunc HandlerFunc)  *RouterBranch{
	return branch.addRoute("GET", routeUrl, routeFunc)
}
func (branch *RouterBranch) AddPostRoute(routeUrl string, routeFunc HandlerFunc)  *RouterBranch{
	return branch.addRoute("POST", routeUrl, routeFunc)
}

func (branch *RouterBranch) addRoute(method string, routeUrl string, routeFunc HandlerFunc) *RouterBranch{
	var totalRoute string
	// "/test1/test2" -> "test1/test2"
	if string(routeUrl[0]) == "/" {
		routeUrl = routeUrl[1:]
	}
	if branch.branchUrl == "/" {
		totalRoute = branch.branchUrl + routeUrl
	} else {
		totalRoute = branch.branchUrl + "/" + routeUrl
	}
	log.Printf("[Router]: RouteAdd '%s - %s'", method, totalRoute)
	key := method + "-" + totalRoute
	
	values := string2strs(totalRoute)

	node := branch.root.insertStrict(values)
	branch.Router.routes[key] = routeFunc

	return &RouterBranch{
		Router: branch.Router, 		
		branchUrl: totalRoute,
		root: node,
	}
}

func (branch *RouterBranch) Use(handlerFunc HandlerFunc) {
	branch.root.handlerFuncs = append(branch.root.handlerFuncs, handlerFunc)
}
