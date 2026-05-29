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

func (s *FavorService) AddFavor(c *gin.Context, uid int64, mid string) {
	count, err := dao.QueryFavorCountByUidAndMid(dao.DB, uid, mid)
	if err != nil {
		util.Fail(c, data.ERR_CODE_DATABASE_ERROR, "数据库查询错误")
		return
	}

	if count > 0 {
		util.Fail(c, data.ERR_CODE_DATA_ERROR, "已经收藏了此歌曲")
		return
	}

	music := MusicSvr.FindCacheMusicByMid(mid)
	if music == nil {
		util.Fail(c, data.ERR_CODE_DATA_ERROR, "未找到此歌曲")
		return
	}

	favorModel, err := dao.AddFavor(dao.DB, uid, mid, 0)
	if err != nil || favorModel == nil {
		util.Fail(c, data.ERR_CODE_DATABASE_ERROR, "插入收藏错误")
		return
	}

	util.Success(c, *favorModel)
}
