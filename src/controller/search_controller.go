package controller

import (
	"geji/data"
	"geji/service"
	"geji/util"

	"github.com/gin-gonic/gin"
)

type QuerySong struct {
	Name   string `json:"name"`
	Singer string `json:"singer"`
}

type QueryResultList struct {
	List []QuerySong `json:"list"`
}

func Search(c *gin.Context) {
	query := c.Query("query")
	util.Logi("query : %s", query)

	if util.IsEmpty(query) {
		util.Fail(c, data.HTTP_CODE_CLIENT_ERR, "未输入查询内容")
		return
	}

	resp, err := service.MusicSvr.Query(query)
	if err != nil {
		util.Fail(c, data.HTTP_CODE_SERVER_ERR, err.Error())
		return
	}

	util.Success(c, resp)
}

func MusicDetail(c *gin.Context) {
	mid := c.Query("mid")
	util.Logi("mid : %s", mid)

	if util.IsEmpty(mid) {
		util.Fail(c, data.ERR_CODE_NOMID, "没有mid参数传入")
		return
	}

	detail, err := service.MusicSvr.GetMusicDetailByMid(mid)
	if err != nil {
		util.Fail(c, data.ERR_CODE_NOTFOUND_MUSIC, err.Error())
		return
	}

	util.Success(c, detail)
}
