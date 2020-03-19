package jwt

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/reedthink/pkg/e"
	"github.com/reedthink/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"mag":  e.GetMsg(code),
				"data": data,
			})
			c.Abort() //中止
			return

		}
		c.Next()
	}
}
