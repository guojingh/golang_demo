package router

import (
	"net/http"

	"github.com/cncamp/golang/gin/Demo2/controller"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {

	r := gin.New()
	v1 := r.Group("/api/v1")
	//这里HandlerFunc传入的是函数的地址而不是实际调用，以后用的时候需要注意
	v1.GET("/students", controller.Student)

	//如果路由不存在返回 404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
