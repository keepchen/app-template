package middleware

import (
	"net/http"

	"github.com/keepchen/app-template/pkg/constants"

	"github.com/gin-gonic/gin"
	"github.com/keepchen/app-template/pkg/app/httpserver/api/response"
)

//AuthCheck 授权检查
func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if len(authorization) == 0 {
			respData := (&response.StandardResponse{}).Assemble(constants.ErrAuthorizationTokenInvalid, nil)
			c.JSON(http.StatusUnauthorized, respData)
			c.Abort()
			return
		}
		//to do something
		c.Next()
	}
}
