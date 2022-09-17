package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
}

func (controller *IndexController) IndexController(context *gin.Engine) {
	context.GET("/", controller.indexHandler)
}
func (controller *IndexController) indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
