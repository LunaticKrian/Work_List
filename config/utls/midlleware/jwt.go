package midlleware

import (
	"awesomeProject/config/utls/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = 200
		// Authorization：存放token
		token := c.GetHeader("Authorization")
		if token == "" {
			code = http.StatusNotFound
		} else {
			claims, err := jwt.ParseToken(token)
			if err != nil {
				code = http.StatusForbidden
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = http.StatusUnauthorized
			}
		}
		if code != http.StatusOK {
			c.JSON(http.StatusOK, gin.H{
				"status": code,
				"msg":    "token解析错误！",
				"data":   data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
