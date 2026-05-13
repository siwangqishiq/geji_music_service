package dao

import (
	"geji/util"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DbMutex *sync.Mutex

func init() {
	util.Logi("init dao module")
	var err error

	DB, err = gorm.Open(sqlite.Open("../database/geji.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DbMutex = &sync.Mutex{}
}
