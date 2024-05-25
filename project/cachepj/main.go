package main

import (
	"fmt"
	"time"

	"github.com/cncamp/golang/project/cachepj/cache"
)

func main() {

	c := cache.NewMemCache()
	c.SetMaxMemory("100MB")
	c.Set("cache", map[string]interface{}{"a": 1}, time.Second*5)
	c.Set("int", 1, time.Second*5)
	/*	c.Set("int", 1)
		cache.Set("bool", false)
		cache.Set("data", map[string]interface{}{"a": 1})*/
	get, b := c.Get("int")
	if b {
		fmt.Println(get)
	}
	c.Del("int")
	c.Flush()
	c.Keys()

}
