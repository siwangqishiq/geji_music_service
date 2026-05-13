package dao

import (
	"geji/util"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

var MucDao *MusicDAO
var KvcDao *KvCacheDAO

func init() {
	util.Logi("init dao module")
	var err error

	DB, err = gorm.Open(sqlite.Open("../database/geji.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	MucDao = &MusicDAO{
		db: DB,
	}

	KvcDao = &KvCacheDAO{
		db: DB,
	}
}
