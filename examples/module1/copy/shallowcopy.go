package main

import (
	"fmt"
	"reflect"
)

func modify(s []int) {
	s[0] = 4
	fmt.Println("modify", s) // modify [4 2 3]
}

func main() {
	s := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	l1 := [3]int{1, 2, 3}
	l2 := [3]int{1, 2, 3}
	b1 := reflect.DeepEqual(s, s2)
	b2 := reflect.DeepEqual(l1, l2)
	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(l1 == l2)
	modify(s)
	fmt.Println("main", s) // main [4 2 3]
}
