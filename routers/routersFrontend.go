package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RfInit(r *gin.Engine) {
	groupRouter := r.Group("/frontend")
	{
		groupRouter.GET("/index", func(c *gin.Context) {
			c.String(http.StatusOK, "前台 - 首页")
		})

		groupRouter.GET("/users", func(c *gin.Context) {
			c.String(http.StatusOK, "前台 - 用户列表")
		})

		groupRouter.GET("/keywords", func(c *gin.Context) {
			//c.JSON(http.StatusOK, gin.H{
			//	"a": 123,
			//	"b": 546,
			//})
			c.HTML(http.StatusOK, "frontend/index", gin.H{
				"title": "前台 - 关键词列表",
			})
		})
	}
}
