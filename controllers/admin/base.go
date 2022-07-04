package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct{}

func (con BaseController) success(c *gin.Context, message string, redirectUrl string) {

	c.HTML(http.StatusOK, "admin/public/success.html", gin.H{
		"message":     message,
		"redirectUrl": redirectUrl,
	})
}

func (con BaseController) error(c *gin.Context, message string, redirectUrl string) {

	c.HTML(http.StatusOK, "admin/public/error.html", gin.H{
		"message":     message,
		"redirectUrl": redirectUrl,
	})
}
