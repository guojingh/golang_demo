package main

import (
	"fmt"

	_ "github.com/cncamp/golang/examples/module1/init/a"
	_ "github.com/cncamp/golang/examples/module1/init/b"
)

func init() {
	fmt.Println("test03 init")
}
func main() {

}
