package service

import (
	"geji/dao"
	"geji/data"
	"geji/model"
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

	var favorList []model.Favor = make([]model.Favor, 0)
	for _, favorModel := range list {
		util.Logi("favor item : %v", favorModel.Mid)

		music := MusicSvr.FindCacheMusicByMid(favorModel.Mid)
		if music != nil {
			var favor model.Favor = model.Favor{
				Fid:        favorModel.Fid,
				Aid:        favorModel.Aid,
				Mid:        favorModel.Mid,
				Remark:     favorModel.Remark,
				Status:     favorModel.Status,
				Sort:       favorModel.Sort,
				CreateTime: favorModel.CreateTime,
				UpdateTime: favorModel.UpdateTime,
				Music:      *music,
			}

			favorList = append(favorList, favor)
		} //end if
	} //end for each

	util.Success(c, favorList)
}

func (s *FavorService) RemoveFavor(c *gin.Context, uid int64, fid int64) {
	favor, err := dao.QueryFavorByFid(dao.DB, fid)
	if err != nil || favor == nil {
		util.Fail(c, data.ERR_CODE_DATABASE_ERROR, "数据库查询错误")
		return
	}

	if favor.Uid != uid {
		util.Fail(c, data.ERR_CODE_DATA_ERROR, "收藏人与删除账户不一致")
		return
	}

	err = dao.RemoveFavorByFid(dao.DB, fid)
	if err != nil {
		util.Fail(c, data.HTTP_CODE_SERVER_ERR, "数据库错误"+err.Error())
		return
	} else {
		util.Success(c, "移除收藏成功")
		return
	}
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

	var favor model.Favor = model.Favor{
		Fid:        favorModel.Fid,
		Aid:        favorModel.Aid,
		Mid:        favorModel.Mid,
		Remark:     favorModel.Remark,
		Status:     favorModel.Status,
		Sort:       favorModel.Sort,
		CreateTime: favorModel.CreateTime,
		UpdateTime: favorModel.UpdateTime,
		Music:      *music,
	}

	util.Success(c, favor)
}
