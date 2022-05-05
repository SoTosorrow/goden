package goden

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type Params map[string]any

type HandlerFunc func(c *Context)

//compare => []*HandlerFunc
type HandlerFuncs []HandlerFunc

type Context struct {
	Response    	http.ResponseWriter
	Request     	*http.Request
	StatusCode  	int

	RequestMethod   string
	RequestRoute    string		// Request URI
	RequestAddr     string		// Request Origin
	RequestTarget   string		// Request Target

	RouteParams 	map[string]string	// router fuzzy-uri match :test=>map[test:111]
	ContextParams 	map[string]any
	HandlerFuncs	HandlerFuncs
	executeIndex	int

	
	/*
		Body io.ReadCloser
		ContentLength int64     
		request.Body
		request.Host
		request.RemoteAddr 
		request.RequestURI 
		request.URL.Path
	*/

}

func newContext(response http.ResponseWriter, request *http.Request) *Context{
	return &Context{
		Response: 		response,
		Request:  		request,

		RequestRoute:   request.URL.Path,	//request.RequestURI,
		RequestMethod:  request.Method,
		RequestAddr:	request.RemoteAddr,
		RequestTarget:  request.Host,

		ContextParams:  make(map[string]any),
		executeIndex:   -1,
		
	}
}


func (c *Context) Next() {
	c.executeIndex++;
	total := len(c.HandlerFuncs)
	for ;c.executeIndex<total;c.executeIndex++{
		c.HandlerFuncs[c.executeIndex](c)
	}
}

// get route param from router match
func (c *Context) GetRouteParam(key string) string {
	value := c.RouteParams[key]
	return value
}

func (c *Context) PostParam(key string) string {
	return c.Request.FormValue(key)
}

func (c *Context) GetParam(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) SetHeader(key string, value string) {
	c.Response.Header().Set(key, value)
}

func (c *Context) SetStatus(status int) {
	c.StatusCode = status
	c.Response.WriteHeader(status)
}

func (c *Context) ReturnJson(status int, data any) {
	c.SetHeader("Content-Type", "application/json")
	c.SetStatus(status)
	encoder := json.NewEncoder(c.Response)
	if err:= encoder.Encode(data); err != nil {
		http.Error(c.Response, err.Error(), 500)
	}
}

func (c *Context) ReturnJsonOk(data any) {
	c.ReturnJson(http.StatusOK, data)
}

func (c *Context) ReturnHtml(status int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.SetStatus(status)
	// turn to strings.Builder
	c.Response.Write([]byte(html))
}

func (c *Context) ReturnHtmlOk(html string) {
	c.ReturnHtml(http.StatusOK, html)
}


func (c *Context) ReturnString(status int, format string, values ...any) {
	c.SetHeader("Content-Type", "text/plain")
	c.SetStatus(status)
	c.Response.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) ReturnStringOk(format string, values ...any) {
	c.ReturnString(http.StatusOK, format, values...)
}

func (c *Context) ReturnData(status int, data []byte) {
	c.SetStatus(status)
	c.Response.Write(data)
}

func (c *Context) ReturnDataOk(data []byte) {
	c.ReturnData(http.StatusOK, data)
}

func (c *Context) ReturnNotFound() {
	c.ReturnString(http.StatusNotFound, "404 NOT FOUND: %s\n", c.RequestRoute)
}