package main

import (
	"github.com/shopspring/decimal"
	"log"
)

/*
第三方库："github.com/shopspring/decimal"
decimal库用来解决float对象之间运算不准确的问题。
使用decimal过程：
1.必须先把float类型对象通过 decimal.NewFromFloat()函数转化成decimal.Decimal类型
2.计算
3.再转换成自己所需要的类型
*/

// golang本身存在精度缺失的问题
// 从输出结果来看，数值的精度出现了一定的问题，这不是我们想要看到的
func main1() {
	a := 1129.6
	log.Printf("%T=%v", a, a*100) //输出：float64=112959.99999999999

	var b float64 = 1129.6
	log.Printf("%T=%v", b, b*100) //输出：float64=112959.99999999999

	m1 := 8.2
	m2 := 3.8
	log.Printf("%v", m1-m2) //输出： 4.3999999999999995
}

/*
第三方库："github.com/shopspring/decimal"
decimal库用来解决float对象之间运算不准确的问题。
使用decimal过程：
1.必须先把float类型对象通过 decimal.NewFromFloat()函数转化成decimal.Decimal类型
2.计算
3.再转换成自己所需要的类型
*/
func main() {
	var v1 = decimal.NewFromFloat(0.1) //声明一个 decimal.Decimal类型的变量 v1
	var v2 = decimal.NewFromFloat(0.2) //声明一个 decimal.Decimal类型的变量 v2

	//decimal.Decimal类型变量之间的加减乘除
	//所得的结果也是 decimal.decimal 类型
	var v3 = v1.Add(v2) //0.3
	var v4 = v1.Sub(v2) //-0.1
	var v5 = v1.Mul(v2) //0.02
	var v6 = v1.Div(v2) //-1
	log.Println(v1, v2, v3, v4, v5, v6)

	//声明一个 decimal.Decimal 类型的对象
	var v7 = decimal.NewFromFloat(3.4625) //创建一个 decimal.Decimal 类型
	var data1 = v7.Round(1)               //3.5 保留一位小数，四舍五入的模式
	var data2 = v7.Truncate(1)            //3.4 保留一位小数，直接舍弃掉的形式
	log.Println(v7, data1, data2)

	//最后还需要把 decimal.Decimal 类型转换为自己需要的类型
	f, _ := decimal.NewFromFloat(7.989).Round(2).Float64()
	log.Println(f) //7.99
}
