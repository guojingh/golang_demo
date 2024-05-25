package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Student(c *gin.Context) {
	fmt.Println("调用 /api/v1/students 路由")
	c.JSON(http.StatusOK, gin.H{
		"msg": "见到你很高兴",
	})
	return
}
