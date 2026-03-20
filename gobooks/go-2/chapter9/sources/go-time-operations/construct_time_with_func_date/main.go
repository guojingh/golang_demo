package main

import (
	"fmt"
	"time"
	"unsafe"
)

// 通过time.Parse、time.Date或time.Unix构建的time.Time结构体，其中的hasMonotonic均为0，
// 即这样构建的Time实例仅表示挂钟时间，而没有单调时间
func dumpWallAndExt(t time.Time) {
	var hasMonotonic int

	//输出wall字段的值
	pWall := (*uint64)(unsafe.Pointer(&t))
	fmt.Printf("0x%X\n", *pWall)
	if (1<<63)&(*pWall) != 0 {
		hasMonotonic = 1
	}
	fmt.Printf("hasMonotonic = %d\n", hasMonotonic)

	//输出ext字段的值
	pExt := (*int64)(unsafe.Pointer((uintptr(unsafe.Pointer(&t)) + unsafe.Sizeof(uint16(0)))))
	fmt.Printf("0x%X\n", *pExt)
	fmt.Printf("%d\n", *pExt/86400/365) //粗略计算距今的年数
}

func constructTimeByData() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("load time location failed:", err)
		return
	}
	t := time.Date(2020, 6, 18, 06, 0, 0, 10000, loc)
	fmt.Println(t)

	dumpWallAndExt(t)
}

func constructTimeByParse() {
	t, _ := time.Parse(time.RFC3339, "2020-06-18T06:00:00.00001+08:00")
	fmt.Println(t)
	dumpWallAndExt(t)
}

func main() {
	constructTimeByData()
	constructTimeByParse()

	fmt.Println("====================")
	t := time.Now()
	dumpWallAndExt(t)
}
