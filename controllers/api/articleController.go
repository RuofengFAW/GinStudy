package api

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title   string `xml:"title"`
	Content string `xml:"content"`
}

func (a Article) Keywords(c *gin.Context) {
	article := &a
	b, _ := c.GetRawData()
	if err := xml.Unmarshal(b, article); err == nil {
		c.JSON(http.StatusOK, article)
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}
