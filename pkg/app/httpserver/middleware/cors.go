package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Cors 允许浏览器跨域请求
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if len(origin) == 0 {
			origin = "*"
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, OPTIONS")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, "+
			"Access-Control-Allow-Headers, Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		//为解决浏览器跨域options探测
		if c.Request.Method == http.MethodOptions {
			c.Writer.WriteHeader(http.StatusOK)
			c.Abort()
			return
		}
	}
}
