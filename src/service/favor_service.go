package service

import (
	"geji/dao"
	"geji/data"
	"geji/util"

	"github.com/gin-gonic/gin"
)

type FavorService struct {
}

var FavorSvr FavorService

func init() {
	util.Logi("init FavorService")
}

func (s *FavorService) FindFavorList(c *gin.Context, uid int64, ablumId int64) {
	list, err := dao.QueryFavorListByAid(dao.DB, uid, ablumId)
	if err != nil {
		util.Fail(c, data.ERR_CODE_DATABASE_ERROR, "数据库查询错误")
		return
	}

	for _, item := range list {
		util.Logi("favor item : %v", item.Mid)
	} //end for each
}
