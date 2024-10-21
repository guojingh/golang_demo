package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

// 基于游标分页

func bookListHandler(c *gin.Context) {
	// 参数校验
	pageToken := c.Query("page_token") // 从请求中获取分页令牌
	// 解token
	page := Token(pageToken).Decode()
	// 分页校验
	if page.NextID == "" {
		c.JSON(401, "bad page token")
		return
	}

	if page.NextTimeAtUTC > time.Now().Unix() || time.Now().Unix()-page.NextTimeAtUTC > int64(time.Hour)*24 {
		c.JSON(401, "bad page token")
		return
	}

	sql := `select id, title from books where id > ? order by id ASC limit ?` // page.NextID page.PageSize

	// 去数据库查数据
	data := db.Query(sql)

	// 拿到最后一条数据，拼接下一页的 page_toke
	nextPage := Page{
		NextID:        "20",
		NextTimeAtUTC: time.Now().Unix(),
		PageSize:      page.PageSize,
	}

	nextPageToken := nextPage.Encode()
	c.JSON(200, gin.H{
		"data":       data,
		"next_token": nextPageToken,
	})
}

func main() {
	r := gin.Default()
	r.GET("/api/v1/books", bookListHandler)
	r.Run()
}
