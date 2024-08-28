package gee

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// 打印执行时间
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since((t)))
	}
}
