package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RbInit(r *gin.Engine) {
	groupRouter := r.Group("/backend")
	{
		groupRouter.GET("/index", func(c *gin.Context) {
			c.String(http.StatusOK, "后台 - 首页")
		})

		groupRouter.GET("/users", func(c *gin.Context) {
			c.String(http.StatusOK, "后台 - 用户列表")
		})

		groupRouter.GET("/keywords", func(c *gin.Context) {
			c.String(http.StatusOK, "后台 - 关键词列表")
		})
	}
}
