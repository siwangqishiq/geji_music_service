package dao

import (
	"time"
)

type MusicModel struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement"`
	Mid         string    `gorm:"column:mid"`
	Author      string    `gorm:"column:author"`
	Name        string    `gorm:"column:name"`
	Href        string    `gorm:"column:href"`
	Cover       string    `gorm:"column:cover"`
	Lyc         string    `gorm:"column:lyc"`
	MusicUrl    string    `gorm:"column:music_url"`
	LocalPath   string    `gorm:"column:local_path"`
	DurationSec int32     `gorm:"column:duration_sec"`
	Desc        string    `gorm:"column:desc"`
	Status      int32     `gorm:"column:status"`
	CreateTime  time.Time `gorm:"column:create_time"`
	UpdateTime  time.Time `gorm:"column:update_time"`
}

type MusicDAO struct {
}

var MucDao *MusicDAO

func init() {
	MucDao = &MusicDAO{}
}

// TableName 指定表名
func (MusicModel) TableName() string {
	return "music"
}

// InsertMusic 插入音乐
func (m *MusicDAO) InsertMusic(music *MusicModel) error {
	return DB.Create(music).Error
}

// QueryAllMusic 查询全部音乐
func (m *MusicDAO) QueryAllMusic() ([]MusicModel, error) {
	var list []MusicModel
	err := DB.Find(&list).Error
	return list, err
}
