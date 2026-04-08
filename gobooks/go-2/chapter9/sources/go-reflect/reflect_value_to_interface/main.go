package main

import (
	"fmt"
	"reflect"
)

// 反射世界的出口
// reflect.Value.Interface()是reflect.ValueOf()的逆过程，
// 通过Interface方法我们可以将reflect.Value对象恢复成一个interface{}类型的变量值
func main() {
	var i = 5
	val := reflect.ValueOf(i)
	r := val.Interface().(int)
	fmt.Println(r)
	r = 6
	fmt.Println(i, r)

	val = reflect.ValueOf(&i)
	q := val.Interface().(*int)
	fmt.Printf("%p, %p, %d\n", &i, q, *q)
	*q = 7
	fmt.Println(i)
}
