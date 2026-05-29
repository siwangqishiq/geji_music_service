package service

import (
	"fmt"
	"geji/config"
	"geji/dao"
	"geji/model"
	"geji/util"
	"strings"
	"time"
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

	MusicSvr.loadMusicToCache()
}

func (msvr *MusicService) loadMusicToCache() error {
	util.Logi("loadMusicToCache ... ")
	musicList, err := dao.MucDao.QueryAllMusic()
	if err != nil {
		util.Loge("load music database error")
		return err
	}
	for _, m := range musicList {
		modelMu := model.DaoMusicToModel(&m)
		msvr.musicMap[m.Mid] = &modelMu
	}
	util.Logi("loadMusicToCache success size : %d ", len(msvr.musicMap))
	return nil
}

func (msvr *MusicService) FindCacheMusicByMid(mid string) *model.Music {
	music, exist := msvr.musicMap[mid]
	if exist {
		return music
	}

	return nil
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

	localPath, downloadErr := msvr.CatchMusicFromUrl(detail)
	if downloadErr != nil {
		return nil, fmt.Errorf("下载资源失败 %s", detail.PlayURL)
	}

	var daoMusic dao.Music = dao.Music{
		Mid:         mid,
		Author:      detail.Author,
		Name:        detail.Title,
		Href:        detail.Href,
		Cover:       detail.Cover,
		MusicUrl:    detail.PlayURL,
		DurationSec: int32(detail.DurSeconds),
		Desc:        "",
		LocalPath:   localPath,
	}

	dbErr := dao.MucDao.InsertMusic(&daoMusic)
	if dbErr != nil {
		return nil, fmt.Errorf("数据库操作失败")
	}

	var retMusic = model.DaoMusicToModel(&daoMusic)

	msvr.musicMap[mid] = &retMusic
	return &retMusic, nil
}

func (s *MusicService) CatchMusicFromUrl(detail *model.SongDetail) (string, error) {
	downloadUrl := detail.PlayURL
	if util.IsEmpty(downloadUrl) {
		util.Loge("download play url but url is null")
		return "", fmt.Errorf("url is empty")
	}

	util.Logi("download url : %s ", downloadUrl)
	ext := util.GetUrlExt(downloadUrl)
	util.Logi("download ext : %s ", ext)
	filename := util.MD5(downloadUrl) + ext
	localPath := fmt.Sprintf("%s/%s", config.MEDIA_PATH, filename)

	util.Logi("localPath : %s ", localPath)

	util.Logi("%s - %s start to download..", detail.Title, detail.Author)
	downloadStartTime := time.Now()
	err := util.DownloadFile(localPath, downloadUrl)
	if err != nil {
		util.Loge("%s - %s download failed %v", detail.Title, detail.Author, err)
		return "", err
	}
	downloadElapsed := time.Since(downloadStartTime)
	util.Logi("%s - %s download spent time %d ms", detail.Title, detail.Author, downloadElapsed.Milliseconds())
	util.Logi("download success localPath : %s ", localPath)
	return filename, nil
}

func (s *MusicService) Query(query string) ([]model.Song, error) {
	var songList []model.Song = []model.Song{}
	list, err := SpiderFangpiSvr.QueryFromFangpiWeb(query)
	songList = append(songList, list...)

	if len(songList) == 0 {
		revealList, reErr := s.RevealSearchContent(query)
		if reErr == nil {
			util.Logi("触发兜底搜素 reveal list size : %d", len(revealList))
			songList = append(songList, revealList...)
		}
	}

	return songList, err
}

func (s *MusicService) RevealSearchContent(query string) ([]model.Song, error) {
	var retList []model.Song = []model.Song{}
	if util.IsEmpty(query) {
		return retList, fmt.Errorf("Query is empty")
	}

	for _, music := range s.musicMap {
		if strings.ContainsAny(music.Name, query) || strings.ContainsAny(music.Author, query) {
			song := model.Song{
				Mid:    music.Mid,
				Author: music.Author,
				Name:   music.Name,
			}
			retList = append(retList, song)
		}
	} //end for music map
	return retList, nil
}
