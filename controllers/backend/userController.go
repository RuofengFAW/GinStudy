package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	BaseController
}

func (u UserController) List(c *gin.Context) {
	u.success(c)
}

func (u UserController) Add(c *gin.Context) {
	c.String(http.StatusOK, "添加用户")
}

func (u UserController) Upload(c *gin.Context) {
	c.HTML(http.StatusOK, "backend/upload", gin.H{

	})
}
