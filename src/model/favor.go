package model

import "time"

type Favor struct {
	Fid        int64     `json:"fid"`
	Aid        int64     `json:"aid,omitempty"`
	Mid        string    `json:"mid,omitempty"`
	Remark     string    `json:"remark,omitempty"`
	Status     int       `json:"status,omitempty"`
	Sort       int       `json:"sort,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
	UpdateTime time.Time `json:"update_time,omitempty"`
	Music      Music     `json:"music,omitempty"`
}

//
type AddFavorReq struct {
	Mid string `json:"mid" binding:"required"`
	Fid int64  `json:"fid"`
}

//
type RemoveFavorReq struct {
	Fid int64 `json:"fid" binding:"required"`
}
