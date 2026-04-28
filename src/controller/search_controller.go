package controller

import (
	"geji/data"
	"geji/util"
	"time"

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

	time.Sleep(10 * time.Second)

	var result QueryResultList = QueryResultList{
		List: []QuerySong{
			{Name: "东风破", Singer: "周杰伦"},
			{Name: "遇见", Singer: "孙燕姿"},
		},
	}

	util.Success(c, result)
}
