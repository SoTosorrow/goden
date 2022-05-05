package goden

import (
	"time"
	"log"
)

func TimerLogger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		// log.Printf("[%v]ns",time.Now().UnixNano())
		c.Next()
		// log.Printf("[%v]ns",time.Now().UnixNano())
		// .Nanoseconds()
		log.Printf("[%d] %s in %v", c.StatusCode, c.RequestRoute, time.Since(t))
	}
}