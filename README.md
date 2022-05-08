# GoDen

* A http web framework with golang

## Current
> http server  
> router/group mapping by trie-tree  
> Fuzzy Matching  
> basic static server

## Todo
> more file handle  
> cookie/session handle  
> distributed network system  
> http cache  
> mongo(or others) eazy-use orm  
> RPC

## Example
```go
// create new Web server
web := goden.New()

// router mapping -> xxx:xxx
web.AddGetRoute("/", func(c *goden.Context) {
    c.ReturnHtmlOk("<h1>Hello Goden</h1>")
})

// get/post params -> xxx:xxx/test?parm1=1&parm2=2
group := web.AddGetRoute("/test",func(c *goden.Context) {
    c.ReturnStringOk("get test %s,%s", c.GetParam("parm1"),c.GetParam("parm2"))
})

// router group -> xxx:xxx/test/secd/thrd
group.AddGetRoute("secd/thrd", func(c *goden.Context) {
    c.ReturnStringOk("group test")
})

// router fuzzy single mapping -> xxx:xxx/test/:match-anything-once
group.AddGetRoute(":more", func(c *goden.Context) {
    c.ReturnStringOk("once-match test")
})


// router fuzzy all-after mapping -> xxx:xxx/test/match/*match-anything-after
group.AddGetRoute("match/*all", func(c *goden.Context) {
    c.ReturnStringOk("all-match test")
})

// custom middleware (define out of main)
func MiddleWare() goden.HandlerFunc {
	return func(c *goden.Context) {
		fmt.Println("middleware test")
		c.Next()
	}
}

// middleware Use
web.Use("/",goden.TimerLogger())

// group middleware Use
group.Use(MiddleWare())

// static file server -> xxx:xxx/assets/xxx.xxx
web.Static("/assets","./")

// run web server
web.Run(":8995")
```