package routers

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RaInit(r *gin.Engine) {

	type Userinfo struct {
		Username string `json:"username" form:"username"`
		Password string `json:"password" form:"password"`
	}

	type Article struct {
		Title   string `xml:"title"`
		Content string `xml:"content"`
	}

	groupRouter := r.Group("/api")
	{
		groupRouter.GET("/users", func(c *gin.Context) {
			//var userinfo Userinfo
			userinfo := &Userinfo{}
			if err := c.ShouldBind(userinfo); err == nil {
				c.JSON(http.StatusOK, userinfo)
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			}
		})

		groupRouter.GET("/keywords", func(c *gin.Context) {
			article := &Article{}
			b, _ := c.GetRawData()
			if err := xml.Unmarshal(b, article); err == nil {
				c.JSON(http.StatusOK, article)
			} else {
				c.JSON(http.StatusBadRequest, err.Error())
			}
		})
	}

	v2GroupRouter := r.Group("/v2/api")
	{
		v2GroupRouter.GET("/users", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"username": "张三",
				"age":      18,
			})
		})

		v2GroupRouter.GET("/keywords", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"keywords": []string{"孕吐", "早孕", "发烧"},
			})
		})
	}
}
