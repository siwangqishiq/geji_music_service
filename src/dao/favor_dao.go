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

func QueryFavorCountByUidAndMid(db *gorm.DB, uid int64, mid string) (int64, error) {
	DbMutex.Lock()
	defer DbMutex.Unlock()

	var count int64 = 0
	err := db.Model(&FavorModel{}).Where("uid = ? and mid = ?", uid, mid).Count(&count).Error
	return count, err
}

func QueryFavorByFid(db *gorm.DB, fid int64) (*FavorModel, error) {
	DbMutex.Lock()
	defer DbMutex.Unlock()

	var result FavorModel = FavorModel{}
	err := db.Where("fid = ?", fid).First(&result).Error
	return &result, err
}

func RemoveFavorByFid(db *gorm.DB, fid int64) error {
	DbMutex.Lock()
	defer DbMutex.Unlock()
	return db.Delete(&FavorModel{}, "fid = ?", fid).Error
}

func AddFavor(db *gorm.DB, uid int64, mid string, ablumId int64) (*FavorModel, error) {
	var nowTime = time.Now()
	var favor FavorModel = FavorModel{
		Aid:        ablumId,
		Uid:        uid,
		Mid:        mid,
		Sort:       0,
		CreateTime: nowTime,
		UpdateTime: nowTime,
	}
	err := db.Create(&favor).Error
	return &favor, err
}
