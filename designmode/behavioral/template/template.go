package template

import "fmt"

type Cooker interface {
	fire()
	cooke()
	outFire()
}

// CookMenu 类似于一个抽象类
type CookMenu struct{}

func (CookMenu) fire() {
	fmt.Println("开火")
}

// 做菜交给具体的子类实现
func (CookMenu) cooke() {

}

func (CookMenu) outFile() {
	fmt.Println("关火")
}

// 封装具体步骤
func doCook(cook Cooker) {
	cook.fire()
	cook.cooke()
	cook.outFire()
}

type XiHongShi struct {
	CookMenu
}

func (XiHongShi) cooke() {
	fmt.Println("做西红柿")
}

type ChaoJiDan struct {
	CookMenu
}

func (ChaoJiDan) cooke() {
	fmt.Println("做炒鸡蛋")
}
