package md5

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestSum(t *testing.T) {
	convey.Convey("基本用例", t, func() {
		shortUrl := "17kyD01"
		md5Value := Sum([]byte(shortUrl))

		// 断言
		convey.Println(md5Value)
	})
}
