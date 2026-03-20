package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)

	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("load location time error:", err)
		return
	}

	t1 := t.In(loc)
	fmt.Println(t1)
	fmt.Println(t.Equal(t1))
}
