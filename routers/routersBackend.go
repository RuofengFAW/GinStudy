package routers

import (
	"api/controllers/backend"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RbInit(r *gin.Engine) {
	groupRouter := r.Group("/backend")
	{
		groupRouter.GET("/index", func(c *gin.Context) {
			c.String(http.StatusOK, "后台 - 首页")
		})

		groupRouter.GET("/keywords", func(c *gin.Context) {
			c.String(http.StatusOK, "后台 - 关键词列表")
		})

		// 路由和控制器分离
		groupRouter.GET("/users", backend.UserController{}.List)
		groupRouter.GET("/user/add", backend.UserController{}.Add)
	}
}
