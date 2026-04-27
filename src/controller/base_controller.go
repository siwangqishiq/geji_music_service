package controller

import (
	"geji/data"

	"github.com/gin-gonic/gin"
)

func parseUserId(c *gin.Context) (string, bool) {
	orgId, exists := c.Get(data.KEY_USER_ID)
	return orgId.(string), exists
}
