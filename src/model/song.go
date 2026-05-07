package model

import (
	"geji/dao"
	"geji/util"
)

type Song struct {
	Mid    string `json:"mid"`
	Author string `json:"author,omitempty"`
	Name   string `json:"name"`
}

type Music struct {
	Mid         string `json:"mid"`
	Author      string `json:"author,omitempty"`
	Name        string `json:"name"`
	Href        string `json:"href,omitempty"`
	Cover       string `json:"cover,omitempty"`
	Lyc         string `json:"lyc,omitempty"`
	MusicUrl    string `json:"musicUrl,omitempty"`
	LocalPath   string `json:"localPath,omitempty"`
	DurationSec int32  `json:"durationSecs,omitempty"`
	Desc        string `json:"desc,omitempty"`
}

func DaoMusicToModel(music *dao.Music) Music {
	var url string = util.GetUrlFromLocalPath(music.LocalPath)
	return Music{
		Mid:         music.Mid,
		Name:        music.Name,
		Author:      music.Author,
		MusicUrl:    url,
		DurationSec: music.DurationSec,
		Cover:       music.Cover,
	}
}

type SongDetail struct {
	ID         int64  `json:"mp3_id"`
	PlayID     string `json:"play_id"`
	Title      string `json:"mp3_title"`
	Author     string `json:"mp3_author"`
	Cover      string `json:"mp3_cover"`
	DurSeconds int
	PlayURL    string
	Href       string
}

type PlayUrlData struct {
	IsWhileURL bool   `json:"is_while_url"`
	URL        string `json:"url"`
	UT         bool   `json:"ut"`
}
