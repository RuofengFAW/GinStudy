package routers

import (
	"api/controllers/api"
	"github.com/gin-gonic/gin"
)

func RaInit(r *gin.Engine) {

	groupRouter := r.Group("/api")
	{
		// users
		groupRouter.GET("/ss", api.Userinfo{}.Ss)
		groupRouter.GET("/db", api.Userinfo{}.Db)
		groupRouter.GET("/users", api.Userinfo{}.Users)
		groupRouter.POST("/upload", api.Userinfo{}.Upload)

		// article
		groupRouter.GET("/keywords", api.Article{}.Keywords)
	}
}
