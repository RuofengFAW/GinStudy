package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

func (u UserController) List(c *gin.Context) {
	c.String(http.StatusOK, "用户列表")
}

func (u UserController) Add(c *gin.Context) {
	c.String(http.StatusOK, "添加用户")
}
