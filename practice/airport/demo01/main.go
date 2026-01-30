package main

import "time"

// 机场安检模拟--顺序
const (
	idCheckTmCost   = 60  //登记身份检查
	bodyCheckTmCost = 120 //人身检查
	xRayCheckTmCost = 180 //x光检查
)

func idCheck() int {
	time.Sleep(time.Millisecond * time.Duration(idCheckTmCost))
	println("\tidCheck ok")
	return idCheckTmCost
}

func bodyCheck() int {
	time.Sleep(time.Millisecond * time.Duration(bodyCheckTmCost))
	println("\tbodyCheck ok")
	return bodyCheckTmCost
}

func xRayCheck() int {
	time.Sleep(time.Millisecond * time.Duration(xRayCheckTmCost))
	println("\txRayCheck ok")
	return xRayCheckTmCost
}

// 安检程序
func airportSecurityCheck() int {
	println("airportSecurityCheck...")
	total := 0
	total += idCheck()
	total += bodyCheck()
	total += xRayCheck()

	println("airportSecurityCheck ok!")
	return total
}

func main() {
	total := 0
	passengers := 30 // 共30个旅客
	for i := 0; i < passengers; i++ {
		total += airportSecurityCheck()
	}
	println("total time cost:", total)
}
