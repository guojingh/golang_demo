package main

import (
	"fmt"
	"reflect"
)

func reflectFunc(i interface{}) {
	value := reflect.ValueOf(i)
	fmt.Println(value)
}

func main() {
	a := 10
	b := 3.14
	c := true
	d := "你好 reflect"

	reflectFunc(a)
	reflectFunc(b)
	reflectFunc(c)
	reflectFunc(d)
}
