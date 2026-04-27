package controller

import (
	"geji/util"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	var msg string = "hello geji music"
	util.Success(c, msg)
}
