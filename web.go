package goden

import (
	"net/http"
	"path"
	"fmt"
)

type Web struct {
	Router *Router
}

func New() *Web {
	return &Web{
		Router: newRouter(),
	}
}

func (web *Web) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// t := time.Now()

	// 如果请求了图标ico，则会多一次请求
	c := newContext(response, request)

	// Router路由分发函数
	web.executeFunc(c)
}

func (web *Web) Run(listenAddress string) (err error) {
	return http.ListenAndServe(listenAddress, web)
}

func (web *Web) executeFunc(c *Context) {
	nodes,params := web.Router.getRoute(c.RequestMethod, c.RequestRoute)
	
	if nodes != nil {
		c.RouteParams = params

		collectHandlerFuncsByNodes(nodes,c)
		// 通过前缀树找到路由
		key := c.RequestMethod + "-" + nodes[len(nodes)-1].fullValue
		// 路由 + 方法
		if web.Router.routes[key] != nil{
			collectHandlerFunc(web.Router.routes[key],c)
			c.Next()
		} else {
			c.ReturnString(http.StatusNotFound, "404 NOT FOUND: %s\n", c.RequestRoute)
		}
		
	} else {
		c.ReturnString(http.StatusNotFound, "404 NOT FOUND: %s\n", c.RequestRoute)
	}
	
}

// !!! Add middleware by node as: node.Use(handlerFunc)
func (web *Web) Use(strs string, handlerFunc HandlerFunc) {
	values := string2strs(strs)
	node := web.Router.root.searchPrefix(values)
	// 存在对应前缀
	if node != nil {
		node.handlerFuncs = append(node.handlerFuncs, handlerFunc)
	}
}

// func (web *Web) AddHandlerFunc(strs string, handlerFunc HandlerFunc) {
// 	values := string2strs(strs)
// 	node := web.Router.root.searchPrefix(values)
// 	// 存在对应前缀
// 	if node != nil {
// 		node.handlerFuncs = append(node.handlerFuncs, handlerFunc)
// 	}
// }

func (web *Web) AddGetRoute(routeUrl string, routeFunc HandlerFunc) *RouterBranch{
	return web.Router.addRoute("GET", routeUrl, routeFunc)
}
func (web *Web) AddPostRoute(routeUrl string, routeFunc HandlerFunc) *RouterBranch{
	return web.Router.addRoute("POST", routeUrl, routeFunc)
}

func (web *Web) Static(requestPath string, originPath string) {
	handlerFunc := web.returnFileServerFunc(requestPath,http.Dir(originPath))
	url := path.Join(requestPath, "/*filepath")
	web.AddGetRoute(url, handlerFunc)
}

func (web *Web) returnFileServerFunc(requestPath string, fs http.FileSystem) HandlerFunc {
	fileServer := http.StripPrefix(requestPath, http.FileServer(fs))
	return func(c *Context) {
		file := c.RouteParams["filepath"]
		fmt.Println(file)
		// Check if file exists and/or if we have permission to access it
		f,err := fs.Open(file)
		if err != nil {
			c.SetStatus(http.StatusNotFound)
			return
		}

		f.Close()
		fileServer.ServeHTTP(c.Response, c.Request)
	}

}