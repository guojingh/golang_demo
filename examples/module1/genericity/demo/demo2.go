package main

import "fmt"

//其他的泛型类型：结构体 带函数的接口 channel
type MyStruct[T int | string] struct {
	Name string
	Data T
}

//一个泛型接口（关于泛型接口在后面会详解）
type IPrintData[T int | float32 | string] interface {
	Print(data T)
}

//一个泛型通道，可用类型实参 int 或者 string 实例化
type MyChan[T int | string] chan T

//类型形参的相互套用
type WowStruct[T int | float32, S []T] struct {
	Data     S
	MaxValue T
	MinValue T
}

//type CommonType[T int | string | float32] T  // 错误用法

//泛型函数
func Add[T int | float32 | float64](a T, b T) T {
	return a + b
}

func main() {
	sum1 := Add(1, 2)
	sum2 := Add(1.0, 2.0)
	fmt.Println(sum1)
	fmt.Printf("%T=%f", sum2, sum2)
}

type int interface {
	~int | ~int8
}
