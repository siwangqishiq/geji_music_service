package controller

import (
	"geji/data"
	"geji/service"
	"geji/util"

	"github.com/gin-gonic/gin"
)

func MusicDetail(c *gin.Context) {
	mid := c.Query("mid")
	util.Logi("mid : %s", mid)

	if util.IsEmpty(mid) {
		util.Fail(c, data.ERR_CODE_NOMID, "没有mid参数传入")
		return
	}

	music, err := service.MusicSvr.GetMusicDetailByMid(mid)
	if err != nil {
		util.Fail(c, data.ERR_CODE_NOTFOUND_MUSIC, err.Error())
		return
	}
	util.Success(c, *music)
}
