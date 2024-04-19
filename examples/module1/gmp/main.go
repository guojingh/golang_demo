package main

import (
	"fmt"
	"runtime"
)

// 获取 GOMAXPROCS GMP 模型中的 P(调度器) 的数量就是通过这个设置的
func getGOMAXPROCS() (int, int) {
	//_ := runtime.NumCPU()
	//runtime.ThreadCreateProfile()
	return runtime.GOMAXPROCS(0), runtime.NumCPU()
}

func main() {
	p, c := getGOMAXPROCS()
	fmt.Printf("GOMAXPROCS: %d; NumCPU: %d\n", p, c)
}
