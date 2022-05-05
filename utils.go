package goden

import (
	"strings"
)

func string2strs(values string) []string {
	splits := strings.Split(values, "/")

	strs := make([]string,0)
	for _,split := range splits {
		if split != "" {
			strs = append(strs, split)
			if split[0] == '*' {
				break
			}
		}
	}
	return strs
}

// TODO 前缀和字符串拼接优化
func joinRouteStrings(values []string) string{
	result := ""
	if len(values) == 0 {
		result = "/"
	}
	for _,val := range values {
		result += "/" + val 
	}
	return result
}

func collectHandlerFuncsByNodes (nodes []*trieNode, c *Context) {
	for _,node := range nodes {
		c.HandlerFuncs = append(c.HandlerFuncs, node.handlerFuncs...)
	}
}

func collectHandlerFunc (handlerFunc HandlerFunc, c *Context) {
	c.HandlerFuncs = append(c.HandlerFuncs, handlerFunc)
}

func executeHandlerFuncs (nodes []*trieNode, c *Context) {
	for _,node := range nodes {
		for _,handlerFunc := range node.handlerFuncs {
			handlerFunc(c) 
		}
	}
}