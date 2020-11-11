package middleware

import (
	"fmt"
	_ "fmt"
	"github.com/gin-gonic/gin"
	"scutrobot.buff/go_demo/response"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Fail(c, fmt.Sprint(err), nil)
				return
			}
		}()
		c.Next()
	}
}
