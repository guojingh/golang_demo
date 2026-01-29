package main

import "fmt"

func main() {

	i := 3
	j := 6
	switch i {
	case 1, 2, 3, 4:
		fmt.Println("i命中了")
	default:
		fmt.Println("default")
	}

	switch j {
	case 5, 8, 9:
		fmt.Println("j没命中")
	default:
		fmt.Println("j命中了default")
	}
}
