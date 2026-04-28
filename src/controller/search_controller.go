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

	// var result QueryResultList = QueryResultList{
	// 	List: []QuerySong{
	// 		{Name: "东风破", Singer: "周杰伦"},
	// 		{Name: "遇见", Singer: "孙燕姿"},
	// 	},
	// }

	html, err := service.SongSvr.Query(query)
	if err != nil {
		util.Fail(c, data.HTTP_CODE_SERVER_ERR, err.Error())
		return
	}

	util.Success(c, html)
}
