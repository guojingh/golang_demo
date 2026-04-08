package main

import (
	"fmt"
	"reflect"
)

func main() {

	//简单原生类型
	var b = true
	val := reflect.ValueOf(b)
	typ := reflect.TypeOf(b)
	fmt.Println(typ.Name(), val.Bool())

	var i = 23
	val = reflect.ValueOf(i)
	typ = reflect.TypeOf(i)
	fmt.Println(typ.Name(), val.Int())

	var f = 3.14
	val = reflect.ValueOf(f)
	typ = reflect.TypeOf(f)
	fmt.Println(typ.Name(), val.Float())

	var s = "hello, reflection"
	val = reflect.ValueOf(s)
	typ = reflect.TypeOf(s)
	fmt.Println(typ.Name(), val.String())

	var fn = func(a, b int) int {
		return a + b
	}

	val = reflect.ValueOf(fn)
	typ = reflect.TypeOf(fn)
	fmt.Println(typ.Kind(), typ.String())

	// 两个变量pi和ps虽然是不同类型的指针，但是它们的Kind值都是ptr
	var pi = (*int)(nil)
	var ps = (*string)(nil)
	typ = reflect.TypeOf(pi)
	fmt.Println(typ.Kind(), typ.String())

	typ = reflect.TypeOf(ps)
	fmt.Println(typ.Kind(), typ.String())

	// 原生复合类型以及其他自定义类型的检视结果
	// 原生复合类型
	var sl = []int{5, 6}
	val = reflect.ValueOf(sl)
	typ = reflect.TypeOf(sl)
	fmt.Printf("[%d %d]\n", val.Index(0).Int(), val.Index(1).Int())
	fmt.Println(typ.Kind(), typ.String())

	//数组
	var arr = [3]int{5, 6}
	val = reflect.ValueOf(arr)
	typ = reflect.TypeOf(arr)
	fmt.Printf("[%d %d %d]\n", val.Index(0).Int(), val.Index(1).Int(), val.Index(2).Int())
	fmt.Println(typ.Kind(), typ.String())

	// map
	var m = map[string]int{
		"tony": 1,
		"jim":  2,
		"john": 3,
	}
	val = reflect.ValueOf(m)
	typ = reflect.TypeOf(m)
	iter := val.MapRange()
	fmt.Printf("{")
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()
		fmt.Printf("%s:%d,", k.String(), v.Int())
	}
	fmt.Printf("}\n")
	fmt.Println(typ.Kind(), typ.String())

	//结构体
	type Person struct {
		Name string
		Age  int
	}

	var p = Person{"tony", 23}
	val = reflect.ValueOf(p)
	typ = reflect.TypeOf(p)
	fmt.Printf("{%s,%d}\n", val.Field(0).String(), val.Field(1).Int())
	fmt.Println(typ.Kind(), typ.Name(), typ.String())

	// channel
	var ch = make(chan int, 1)
	val = reflect.ValueOf(ch)
	typ = reflect.TypeOf(ch)
	ch <- 17
	// 通过反射的非阻塞方式，尝试从通道ch中读取数据。
	v, ok := val.TryRecv()
	if ok {
		fmt.Println(v.Int())
	}
	fmt.Println(typ.Kind(), typ.String())

	// 其它自定义类型
	type MyInt int

	var mi MyInt = 19
	val = reflect.ValueOf(mi)
	typ = reflect.TypeOf(mi)
	fmt.Println(typ.Name(), typ.Kind(), typ.String(), val.Int())
}
