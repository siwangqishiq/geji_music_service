package service

import (
	"fmt"
	"geji/model"
	"geji/util"
)

type MusicService struct {
	musicMap map[string]model.Song
}

var MusicSvr *MusicService

func init() {
	util.Logi("init song service")
	MusicSvr = &MusicService{
		musicMap: make(map[string]model.Song),
	}
}

func (msvr *MusicService) GetMusicDetailByMid(mid string) (model.Song, error) {
	if util.IsEmpty(mid) {
		return model.Song{}, fmt.Errorf("mid is empty")
	}

	detail, exist := msvr.musicMap[mid]
	if exist {
		return detail, nil
	}

	pSrc, pId := util.ParseMid(mid)
	if pSrc == nil || pId == nil {
		return model.Song{}, fmt.Errorf("解析mid失败")
	}

	src := *pSrc
	id := *pId
	util.Logi("query src: %s id: %s", src, id)

	var result *model.Song
	var err error = nil
	if src == util.MUSIC_SRC_FANGPI {
		result, err = SpiderFangpiSvr.SpiderMusicDetail(id)
	}

	//add to cache
	if result != nil && err == nil {
		msvr.musicMap[mid] = *result
		//todo save to database
	}
	return model.Song{}, fmt.Errorf("未发现mid对应的详情数据")
}

func (s *MusicService) Query(query string) ([]model.Song, error) {
	var songList []model.Song = []model.Song{}
	list, err := SpiderFangpiSvr.QueryFromFangpiWeb(query)
	songList = append(songList, list...)
	return songList, err
}
