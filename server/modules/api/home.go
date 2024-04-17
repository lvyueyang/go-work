package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func NewHomeController(e *gin.Engine) {
	c := &HomeController{}
	router := e.Group("/")
	router.GET("/", c.HomePage)
}

// HomePage 主页
func (c *HomeController) HomePage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"title": "男生自用",
	})
}
