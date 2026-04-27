package middleware

import (
	"geji/config"
	"geji/data"
	"geji/util"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		util.Logi("%s auth middleware", c.Request.URL.RawPath)

		token := c.GetHeader(data.KEY_TOKEN)
		util.Logi("request get token : %s", token)

		// 调试模式下使用默认值
		if config.IS_DEBUG {
			c.Set(data.KEY_USER_ID, 1)
		}

		c.Next()
	}
}
