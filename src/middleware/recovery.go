package middleware

import (
	"fmt"
	"geji/data"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("panic: %v\n", err)
				fmt.Printf("request: %s %s\n", c.Request.Method, c.Request.URL.Path)
				fmt.Printf("IP: %s\n", c.ClientIP())
				debug.PrintStack()
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": data.HTTP_CODE_SERVER_ERR,
					"msg":  "服务器内部错误",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
