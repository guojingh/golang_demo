package connect

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestGet(t *testing.T) {
	convey.Convey("基础用例", t, func() {
		url := "https://www.liwenzhou.com/posts/Go/golang-menu/"
		get := Get(url)

		// 断言
		convey.So(get, convey.ShouldEqual, true)
		//convey.ShouldBeTrue(get)

	})

	convey.Convey("url请求不通示例", t, func() {
		url := "posts/go/unit-test-s"
		get := Get(url)

		// 断言
		convey.So(get, convey.ShouldEqual, false)
		//convey.ShouldBeTrue(get)
	})
}
