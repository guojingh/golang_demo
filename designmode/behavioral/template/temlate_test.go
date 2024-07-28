package template

import (
	"fmt"
	"testing"
)

func TestTemplate(t *testing.T) {
	//做西红柿
	xiHongShi := &XiHongShi{}
	doCook(xiHongShi)

	fmt.Println("\n=====> 做另一道菜")
	//做炒鸡蛋
	chaoJiDan := &ChaoJiDan{}
	doCook(chaoJiDan)
}
