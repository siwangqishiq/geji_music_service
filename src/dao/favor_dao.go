package dao

import (
	"time"

	"gorm.io/gorm"
)

type FavorModel struct {
	Fid        int64     `gorm:"column:fid;primaryKey;autoIncrement" json:"fid"`
	Aid        int64     `gorm:"column:aid;index:idx_aid" json:"aid"`
	Uid        int64     `gorm:"column:uid;index:idx_uid" json:"uid"`
	Mid        string    `gorm:"column:mid;type:text" json:"mid"`
	Remark     string    `gorm:"column:remark;type:text" json:"remark"`
	Status     int       `gorm:"column:status;default:0" json:"status"`
	Sort       int       `gorm:"column:sort;default:0" json:"sort"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time"`
}

func (FavorModel) TableName() string {
	return "favor"
}

func QueryFavorListByAid(db *gorm.DB, uid int64, aid int64) ([]FavorModel, error) {
	DbMutex.Lock()
	defer DbMutex.Unlock()

	var list []FavorModel

	err := db.
		Where("uid = ? and aid = ?", uid, aid).
		Order("update_time DESC").
		Find(&list).Error

	return list, err
}
