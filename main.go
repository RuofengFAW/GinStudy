package main

import (
	"api/middlewares"
	"api/routers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {

	// 创建一个默认路由引擎
	r := gin.Default()


	// 加载模板路径
	r.LoadHTMLGlob("templates/**/*")


	// 加载静态资源
	r.Static("/static", "./static")


	// 注册中间件
	r.Use(middlewares.InitMiddleware)


	// 注册 session 中间件
	// 基于 cookie 存储session
	// store := cookie.NewStore([]byte("secret11111"))
	// r.Use(sessions.Sessions("mysession", store))

	// 基于 redis 存储session
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))


	// 注册路由
	routers.RaInit(r)
	routers.RbInit(r)
	routers.RfInit(r)

	r.Run(":8000")
}
