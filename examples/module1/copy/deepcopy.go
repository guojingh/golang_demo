package main

import "fmt"

func modify1(a [3]int) {
	a[0] = 4
	fmt.Println("modify", a) //modify [4 2 3]
}

func main1() {
	a := [3]int{1, 2, 3}
	modify1(a)
	fmt.Println("main", a) //main [1 2 3]
}
