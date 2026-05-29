package controller

import (
	"geji/model"
	"geji/service"
	"geji/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetFavorList(c *gin.Context) {
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

func AddFavor(c *gin.Context) {
	userClaims, exists := GetUserClaims(c)
	if !exists {
		util.Fail(c, http.StatusNetworkAuthenticationRequired, "未登录 无法收藏")
		return
	}

	var req model.AddFavorReq
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Fail(c, http.StatusBadRequest, "参数错误: 未传入mid")
		return
	}

	uid := userClaims.Uid
	util.Logi("AddFavor uid : %v mid:%v", uid, req.Mid)

	service.FavorSvr.AddFavor(c, uid, req.Mid)
}

func RemoveFavor(c *gin.Context) {
	userClaims, exists := GetUserClaims(c)
	if !exists {
		util.Fail(c, http.StatusNetworkAuthenticationRequired, "未登录 无法删除收藏")
		return
	}

	uid := userClaims.Uid
	aidStr := c.Query("aid")
	util.Logi("GetFavor uid : %v aid:%v", uid, aidStr)
}
