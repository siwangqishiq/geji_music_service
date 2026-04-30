package service

import (
	"fmt"
	"geji/model"
	"geji/util"
)

type MusicService struct {
	musicMap map[string]*model.Music
}

var MusicSvr *MusicService

func init() {
	util.Logi("init song service")
	MusicSvr = &MusicService{
		musicMap: make(map[string]*model.Music),
	}
}

func (msvr *MusicService) GetMusicDetailByMid(mid string) (*model.Music, error) {
	if util.IsEmpty(mid) {
		return nil, fmt.Errorf("mid is empty")
	}

	music, exist := msvr.musicMap[mid]
	if exist {
		util.Logi("%s hit cache", mid)
		return music, nil
	}

	pSrc, pId := util.ParseMid(mid)
	if pSrc == nil || pId == nil {
		return nil, fmt.Errorf("解析mid失败")
	}

	src := *pSrc
	id := *pId
	util.Logi("query src: %s id: %s", src, id)

	var detail *model.SongDetail
	var err error = nil

	if src == util.MUSIC_SRC_FANGPI {
		detail, err = SpiderFangpiSvr.SpiderMusicDetail(id)
	}

	if detail == nil || err != nil {
		util.Loge("SpiderMusicDetail Get failed")
		return nil, err
	}

	var retMusic = model.Music{
		Mid:         mid,
		Name:        detail.Title,
		Author:      detail.Author,
		MusicUrl:    detail.PlayURL,
		DurationSec: int32(detail.DurSeconds),
		Cover:       detail.Cover,
	}

	msvr.musicMap[mid] = &retMusic
	return &retMusic, nil
}

func (s *MusicService) Query(query string) ([]model.Song, error) {
	var songList []model.Song = []model.Song{}
	list, err := SpiderFangpiSvr.QueryFromFangpiWeb(query)
	songList = append(songList, list...)
	return songList, err
}
