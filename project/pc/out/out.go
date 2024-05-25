package out

import "fmt"

var out *Out

type Out struct {
	/*
		声明channel 注意：
		1.data chan interface{}：可写可读
		2.data <-chan interface{}：只读
		3.data chan<- interface{}：只写
	*/
	data chan interface{}
}

func NewOut() *Out {
	if out == nil {
		out = &Out{
			data: make(chan interface{}, 1024),
		}
	}
	return out
}

func Println(i interface{}) {
	out.data <- i
}

func (o *Out) OutPut() {
	for {
		select {
		case i := <-o.data:
			fmt.Println(i)
		}
	}
}
