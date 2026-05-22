package controller

import (
	"geji/component"
	"geji/data"

	"github.com/gin-gonic/gin"
)

func ParseUserId(c *gin.Context) (string, bool) {
	orgId, exists := c.Get(data.KEY_USER_ID)
	return orgId.(string), exists
}

func GetUserClaims(c *gin.Context) (component.UserClaims, bool) {
	if claims, exists := c.Get(data.KEY_USER_CLAIMS); exists {
		return claims.(component.UserClaims), exists
	}
	return component.UserClaims{}, false
}
