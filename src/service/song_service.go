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

type SongService struct {
	BaseSvr *BaseService
}

var SongSvr SongService = SongService{
	BaseSvr: BaseSvr,
}

func (s *SongService) Query(query string) ([]model.Song, error) {
	url := fmt.Sprintf("https://www.fangpi.net/s/%s", query)
	respHtml, err := s.FetchHtml(url)
	if err != nil {
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

	util.Logi("doc parse ...")

	// 对应 getElementsByClass("col-8 col-content")
	doc.Find(".col-8.col-content").Each(func(i int, s *goquery.Selection) {
		html, _ := s.Html()
		util.Logi("%d html : %s\n", i, html)
	})

	var songList []model.Song
	// 对应:
	// doc.getElementsByClass("card-body").first()?.getElementsByClass("row")
	doc.Find(".card-body").First().Find(".row").Each(func(i int, row *goquery.Selection) {
		author := strings.TrimSpace(row.Find("small").First().Text())
		music := strings.TrimSpace(row.Find("span").First().Text())
		href, _ := row.Find("a").First().Attr("href")

		util.Logi("%s - %s \t%s\n", music, author, href)
		if author != "" && music != "" && href != "" {
			songList = append(songList, model.Song{
				Author: author,
				Name:   music,
				Href:   href,
			})
		}
	})

	util.Logi("doc parse end")
	return songList, nil
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
