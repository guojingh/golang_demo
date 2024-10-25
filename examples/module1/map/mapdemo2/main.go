package main

import "fmt"

func main() {

	m := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	for k, v := range m {
		fmt.Printf("k: %v, v: %v\n", k, v)
	}

	if vA, ok := m["A"]; !ok {
		fmt.Println("vA is nil")
	} else {
		fmt.Printf("v: %v\n", vA)
	}

	if vD, ok := m["D"]; !ok {
		fmt.Println("vD is nil")
	} else {
		fmt.Printf("v: %v\n", vD)
	}

	//p := &m["A"]  //无法提取 'm["A"]' 的地址
	//fmt.Println(p)
}
