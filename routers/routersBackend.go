package routers

import (
	"api/controllers/backend"
	"github.com/gin-gonic/gin"
)

func RbInit(r *gin.Engine) {
	groupRouter := r.Group("/backend")
	{
		// 路由和控制器分离
		groupRouter.GET("/users", backend.UserController{}.List)
		groupRouter.GET("/user/add", backend.UserController{}.Add)
	}
}
