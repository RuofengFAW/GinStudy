package main

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 加载静态资源
	r.Static("/static", "./static")

	// 加载模板路径
	r.LoadHTMLGlob("templates/**/*")

	// 测试 api
	type Article struct {
		Title   string `xml:"title"`
		Content string `xml:"content"`
	}
	r.POST("/xml", func(c *gin.Context) {
		article := &Article{}
		b, _ := c.GetRawData()
		if err := xml.Unmarshal(b, article); err == nil {
			c.JSON(http.StatusOK, article)
		} else {
			c.JSON(http.StatusBadRequest, err.Error())
		}
	})

	// 前台
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "frontend/index", gin.H{
			"title": "前台 - 首页",
		})
	})

	r.GET("/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "frontend/news", gin.H{
			"title": "前台 - 新闻页",
		})
	})

	// 后台
	r.GET("/backend/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "backend/index", gin.H{
			"title": "后台 - 首页",
		})
	})

	r.GET("/backend/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "backend/index", gin.H{
			"title": "后台 - 新闻页",
		})
	})

	r.Run(":8000")
}
