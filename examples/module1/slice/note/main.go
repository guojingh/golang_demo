package main

import "fmt"

func main() {

	fatherSlice := make([]int, 4, 5)

	fmt.Printf("father cap + %d\n", cap(fatherSlice))

	//母子切片共享一片內存空间
	sonSlice := fatherSlice[1:3]
	sonSlice = append(sonSlice, 1)
	sonSlice = append(sonSlice, 2)

	//子数组cap扩容导致母子切片内存空间分离
	//sonSlice = append(sonSlice, 3)
	/*	fatherSlice = append(fatherSlice, 1)
		fatherSlice = append(fatherSlice, 1)
		fatherSlice = append(fatherSlice, 1)*/

	fatherSlice[1] = 10

	fmt.Printf("son cap + %d\n", cap(sonSlice))
	fmt.Printf("father cap + %d\n", cap(fatherSlice))

	fmt.Printf("father num=%v\n", fatherSlice)
	fmt.Printf("son num=%v\n", sonSlice)

	fmt.Printf("father addr=%d\n", &fatherSlice)
	fmt.Printf("son addr=%d\n", &sonSlice)

}
