package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//SayHello say hello
// @Tags httpServer相关
// @Summary httpServer
// @Description sayHello
// @Accept application/json
// @Produce json
// @Success 200 {string} string
// @Failure 400 {string} string
// @Router /say-hello [get]
func SayHello(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}
