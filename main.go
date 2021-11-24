package main

import (
	"api/routers"
	"github.com/gin-gonic/gin"
)

func main() {

	// 创建一个默认路由引擎
	r := gin.Default()

	// 加载模板路径
	r.LoadHTMLGlob("templates/**/*")

	// 加载静态资源
	r.Static("/static", "./static")

	// 注册路由
	routers.RaInit(r)
	routers.RbInit(r)
	routers.RfInit(r)

	r.Run(":8000")
}
