package dao

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type AccountModel struct {
	Uid        int64     `gorm:"column:uid;primaryKey;autoIncrement" json:"uid"`
	Account    string    `gorm:"column:account;type:text;not null;uniqueIndex" json:"account"`
	Password   string    `gorm:"column:password;type:text" json:"password"`
	Nickname   string    `gorm:"column:nickname;type:text" json:"nickname"`
	Remark     string    `gorm:"column:remark;type:text" json:"remark"`
	Age        string    `gorm:"column:age;type:text" json:"age"`
	Avater     string    `gorm:"column:avater;type:text" json:"avater"`
	Desc       string    `gorm:"column:desc;type:text" json:"desc"`
	Status     int       `gorm:"column:status;default:0" json:"status"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time"`
}

// TableName 指定表名
func (AccountModel) TableName() string {
	return "account"
}

// InsertAccount 插入单条记录
func InsertAccount(db *gorm.DB, account *AccountModel) error {
	return db.Create(account).Error
}

func QueryAccountByAccount(db *gorm.DB, account string) (*AccountModel, error) {
	var acc AccountModel
	err := db.Where("account = ?", account).First(&acc).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &acc, nil
}

func UpdateAccount(db *gorm.DB, account *AccountModel) error {
	return db.Model(&AccountModel{}).Where("uid = ?", account.Uid).Updates(account).Error
}

// UpdateAccountByMap 使用 map 更新（零值也会更新）
func UpdateAccountByMap(db *gorm.DB, uid int64, updates map[string]any) error {
	return db.Model(&AccountModel{}).Where("uid = ?", uid).Updates(updates).Error
}
