package middleware

import (
	"geji/component"
	"geji/data"
	"geji/util"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		util.Logi("%s auth middleware", c.Request.URL.RawPath)

		token := c.GetHeader(data.KEY_TOKEN)
		util.Logi("request get token : %s", token)

		if !util.IsEmpty(token) {
			userClaims, err := component.ParseTokenToUserClaims(token)
			if err != nil {
				util.Logi("request from uid : %v", userClaims.Uid)
				c.Set(data.KEY_USER_CLAIMS, userClaims)
			}
		}

		c.Next()
	}
}
