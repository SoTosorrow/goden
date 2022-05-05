package goden

import (

)

type trieNode struct {
	value			string
	fullValue 		string
	isFuzzyMatch 	bool
	currentHeight	int
	children		[]*trieNode
	handlerFuncs	HandlerFuncs
}

func newTrieNode() *trieNode {
	return &trieNode{
		value: "",
		fullValue: "",
		currentHeight: 0,
	}
}

func (n *trieNode) matchStrictChild(value string) *trieNode {
	for _,child := range n.children {
		if child.value == value {
			return child
		}
	}
	return nil
}

func (n *trieNode) matchFuzzyChild(value string) *trieNode {
	for _,child := range n.children {
		if child.isFuzzyMatch {
			return child
		}
	}
	return nil
}

//!!! 严格插入一串完全前缀 (新的更长的前缀insert时，会导致一个分支出现多个fullValue，即多个终点)
func (n *trieNode) insertStrict(values []string) *trieNode{
	// 若前缀串长度 与 当前节点高度 相同， 则前缀串插入完毕
	if n.currentHeight == len(values) {
		// 拼接前缀之和
		n.fullValue = joinRouteStrings(values)
		return n
	}
	// 下一个前缀
	value := values[n.currentHeight]
	child := n.matchStrictChild(value)
	// 若没有可匹配的子节点，则新建
	if child == nil {
		// 新建的子节点高度+1，逐渐逼近前缀串末尾
		child = &trieNode{
			value: value,
			isFuzzyMatch: value[0]==':' || value[0]=='*',
			currentHeight: n.currentHeight+1,
		}
		n.children = append(n.children,child)
	}
	return child.insertStrict(values)
}

//
func (n *trieNode) insertAfter(value []string) {
	return
}

// 严格查询某完整前缀串是否存在 并储存途径所有节点
func (n *trieNode) searchStrict(values []string) []*trieNode {
	// 从root 0节点开始
	nodes := make([]*trieNode, 0)
	nodes = append(nodes, n)
	i := 0 
	resultNodeValue := joinRouteStrings(values)

	for node:=n; node.fullValue!=resultNodeValue; {
		if i>len(values)-1 {
			return nil
		}
		child := node.matchStrictChild(values[i])
		
		if child != nil {
			nodes = append(nodes, child)
			node = child
			i++
			
		} else{
			return nil
		}
	}
	return nodes
}

// 查询包含模糊匹配的第一个路径
func (n *trieNode) search(values []string) ([]*trieNode,map[string]string) {
	// 从root 0节点开始
	nodes := make([]*trieNode, 0)
	params := make(map[string]string)
	nodes = append(nodes, n)
	i := 0 
	resultNodeValue := joinRouteStrings(values)

	for node:=n; node.fullValue!=resultNodeValue; {
		if i>len(values)-1 {
			return nil,nil
		}
		// 严格查询是否有完全匹配的路由
		child := node.matchStrictChild(values[i])
		
		// 若没有完全符合的路由，查询是否有模糊匹配的路由
		if child == nil{
			child = node.matchFuzzyChild(values[i])
			if child == nil {
				return nil,nil
			}
			if child.value[0] == '*' {
				params[child.value[1:]] = joinRouteStrings(values[i:])[1:]
				// fmt.Println("* match replace",values[i],child.value)
				nodes = append(nodes, child)
				return nodes,params
			}
			if child.value[0] == ':' {
				// fmt.Println(": match replace",values[i],child.value)
				// TODO 正则替换
				params[child.value[1:]] = values[i]
				values[i] = child.value
				resultNodeValue = joinRouteStrings(values)
			}
		}
		
		// 若有匹配到的路由
		if child != nil {
			nodes = append(nodes, child)
			node = child
			i++
			
		} else{
			return nil,nil
		}
	}
	return nodes,params
}

// 查找前缀串
func (n *trieNode) searchPrefix(values []string) *trieNode {
	i := 0 
	node := n
	for ; i!=len(values); {
		child := node.matchStrictChild(values[i])
		
		if child != nil {
			node = child
			i++
		} else{
			return nil
		}
	}
	return node
}

