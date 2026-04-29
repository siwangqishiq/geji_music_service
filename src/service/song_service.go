package service

import (
	"fmt"
	"geji/model"
	"geji/util"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const FANGPI_URL string = "https://www.fangpi.net"

type SongService struct {
	songsMap map[string]model.Song
}

var SongSvr *SongService

func init() {
	util.Logi("init song service")
	SongSvr = &SongService{
		songsMap: make(map[string]model.Song),
	}
}

func (s *SongService) Query(query string) ([]model.Song, error) {
	var songList []model.Song = []model.Song{}
	list, err := s.QueryFromFangpiWeb(query)
	songList = append(songList, list...)
	return songList, err
}

// 从FP网站爬取数据
func (s *SongService) QueryFromFangpiWeb(query string) ([]model.Song, error) {
	url := fmt.Sprintf("%s/s/%s", FANGPI_URL, query)
	respHtml, err := s.FetchHtml(url)
	if err != nil {
		util.Logi("fetch html form %s error", url)
		return nil, err
	}

	songList, err := s.ParseQueryHtml(respHtml)
	return songList, err
}

func (s *SongService) ParseQueryHtml(respHtml string) ([]model.Song, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(respHtml))
	if err != nil {
		return nil, err
	}

	util.Logi("respHtml parse ...")

	// 对应 getElementsByClass("col-8 col-content")
	doc.Find(".col-8.col-content").Each(func(i int, s *goquery.Selection) {
		html, _ := s.Html()
		util.Logi("%d html : %s\n", i, html)
	})

	var songList []model.Song
	// 对应:
	// doc.getElementsByClass("card-body").first()?.getElementsByClass("row")
	doc.Find(".card-body").First().Find(".row").Each(func(i int, row *goquery.Selection) {
		song := s.fillSpiderSong(row)
		if song != nil {
			songList = append(songList, *song)
		}
	})

	util.Logi("respHtml parse end")
	return songList, nil
}

func (s *SongService) fillSpiderSong(row *goquery.Selection) *model.Song {
	if row == nil {
		return nil
	}

	author := strings.TrimSpace(row.Find("small").First().Text())
	music := strings.TrimSpace(row.Find("span").First().Text())
	href, _ := row.Find("a").First().Attr("href")

	if author == "" || music == "" || href == "" {
		return nil
	}

	originId := FindOriginIdFromHref(href)
	var musicId string = util.BuildMid(originId, util.MUSIC_SRC_FANGPI)

	util.Logi("mid: %s\t%s - %s \t%s\n", musicId, music, author, href)
	return &model.Song{
		Mid:    musicId,
		Author: author,
		Name:   music,
		Href:   href,
	}
}

func FindOriginIdFromHref(href string) string {
	if util.IsEmpty(href) {
		return ""
	}

	index := strings.LastIndex(href, "/")
	if index >= 0 {
		return href[index+1:]
	}
	return ""
}

func (s *SongService) FetchHtml(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
