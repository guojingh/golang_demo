package main

import "fmt"

//泛型：在 go 1.8 之后新加了泛型的使用
//一些简单的泛型使用包括切片和map
type Slice[T int | float32 | float64] []T

// MyMap类型定义了两个类型形参 KEY 和 VALUE。分别为两个形参指定了不同的类型约束
// 这个泛型类型的名字叫： MyMap[KEY, VALUE]
type MyMap[KEY int | string, VALUE float32 | float64] map[KEY]VALUE

func main1() {

	// 这里传入了类型实参int，泛型类型Slice[T]被实例化为具体的类型 Slice[int]
	var a Slice[int] = []int{1, 2, 3}
	fmt.Printf("Type name: %T", a)

	// 传入类型实参float32, 将泛型类型Slice[T]实例化为具体的类型 Slice[float32]
	var b Slice[float32] = []float32{1.0, 2.0, 3.0}
	fmt.Printf("Type name: %T", b)

	// ✗ 错误。因为变量a的类型为Slice[int]，b的类型为Slice[float32]，两者类型不同
	a = b

	// ✗ 错误。string不在类型约束 int|float32|float64 中，不能用来实例化泛型类型
	var c Slice[string] = []string{"Hello", "World"}

	// ✗ 错误。Slice[T]是泛型类型，不可直接使用必须实例化为具体的类型
	var x Slice[T] = []int{1, 2, 3}

	// 用类型实参 string 和 flaot64 替换了类型形参 KEY 、 VALUE，泛型类型被实例化为具体的类型：MyMap[string, float64]
	var a1 MyMap[string, float64] = map[string]float64{
		"jack_score": 9.6,
		"bob_score":  8.4,
	}
}
