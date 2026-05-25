package controller

import (
	"geji/service"
	"geji/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetFavor(c *gin.Context) {
	userClaims, exists := GetUserClaims(c)
	if !exists {
		util.Fail(c, http.StatusNetworkAuthenticationRequired, "未查询到登录信息")
		return
	}

	uid := userClaims.Uid
	aidStr := c.Query("aid")
	util.Logi("GetFavor uid : %v aid:%v", uid, aidStr)

	var aid int64 = 0
	if util.IsEmpty(aidStr) {
		aid = 0
	} else {
		value, err := strconv.ParseInt(aidStr, 10, 0)
		if err != nil {
			aid = 0
		} else {
			aid = value
		}
	}

	service.FavorSvr.FindFavorList(c, uid, aid)
}
