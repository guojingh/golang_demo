package main

import (
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {

	a := 10
	b := 20
	a, b = b, a
	fmt.Println(a, b)

	//Println 写到标准日志记录器
	log.Println("message")
	//Fatalln在调用Println之后，会调用 os.Exit(1)
	log.Fatalln("fatal message")
	//Panicln在调用Println之后，会调用 panic()
	log.Panicln("panic message")
}
