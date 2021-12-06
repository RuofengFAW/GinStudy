package api

import (
	"api/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

type Userinfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (u Userinfo) Users(c *gin.Context) {
	//var userinfo Userinfo
	userinfo := &u
	if err := c.ShouldBind(userinfo); err == nil {
		c.JSON(http.StatusOK, userinfo)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

func (u Userinfo) Upload(c *gin.Context) {
	username := c.PostForm("username")
	file, _ := c.FormFile("file")

	// 上传至指定目录
	dst := path.Join("./static/upload", file.Filename)
	c.SaveUploadedFile(file, dst)

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"username": username,
		"dst":      dst,
	})
}

func (u Userinfo) Ss(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("count")
	var count int
	if v == nil {
		count = 0
	} else {
		count = v.(int)
		count++
	}

	session.Set("count", count)
	session.Save()
	c.JSON(200, gin.H{
		"count": count,
	})
}

// 数据库查询
func (u Userinfo) Db(c *gin.Context) {
	list := []models.DAdvertiserList{}
	models.DB.Find(&list)
	c.JSON(http.StatusOK, gin.H{
		"result": list,
	})
}
