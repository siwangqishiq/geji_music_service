package service

import (
	"encoding/json"
	"fmt"
	"geji/model"
	"geji/util"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const FANGPI_URL string = "https://www.fangpi.net"

type SpiderFangpiService struct {
}

var SpiderFangpiSvr *SpiderFangpiService

func init() {
	util.Logi("init web spider fangpi")
	SpiderFangpiSvr = &SpiderFangpiService{}
}

func (s *SpiderFangpiService) SpiderMusicDetail(id string) (*model.SongDetail, error) {
	if util.IsEmpty(id) {
		util.Loge("id is empty")
		return nil, nil
	}

	href := fmt.Sprintf("%s/music/%s", FANGPI_URL, id)
	util.Logi("request url: %s", href)
	htmlResp, err := BaseSvr.FetchHtml(href)
	if err != nil {
		return nil, err
	}
	songDetail, err := s.GetDetailFromHtml(htmlResp, href)
	songDetail.Href = href
	return songDetail, err
}

func (s *SpiderFangpiService) GetDetailFromHtml(html string, href string) (*model.SongDetail, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	var appData string

	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		scriptContent := s.Text()
		key := "window.appData = "
		idx := strings.Index(scriptContent, key)
		if idx != -1 {
			appData = strings.TrimSpace(scriptContent[idx+len(key):])
			appData = strings.TrimSuffix(appData, ";")
		}
	})

	if appData == "" {
		return nil, fmt.Errorf("appData not found")
	}

	if strings.HasPrefix(appData, "JSON.parse") {
		appData = extractJSON(appData)
		util.Logi("parsed appData:%s", appData)
	}

	songDetail, err := s.DecodeSongDetail(appData)
	if songDetail == nil {
		return nil, err
	}
	util.Logi("songDetail info:%v %v %v %v", songDetail.Title, songDetail.Author, songDetail.Cover, songDetail.DurSeconds)
	util.Logi("songDetail playId:%s", songDetail.PlayID)

	playUrl := s.FetchPlayUrl(songDetail.PlayID, href)
	util.Logi("fetch play url %s", playUrl)
	songDetail.PlayURL = playUrl
	return songDetail, nil
}

func (s *SpiderFangpiService) FetchPlayUrl(playId string, href string) string {
	rawResp, err := postForm(fmt.Sprintf("%s/member/common-play-url", FANGPI_URL), href, map[string]string{
		"id": playId,
	})

	if err != nil {
		util.Loge("postform error %v", err.Error())
		return ""
	}

	respCleanJson := CleanJSON(rawResp)
	util.Logi("fetch resp %s", respCleanJson)

	var resp model.HttpResp[model.PlayUrlData]
	err = json.Unmarshal([]byte(respCleanJson), &resp)
	if err != nil {
		util.Loge("postform parse %v", err.Error())
		return ""
	}
	return resp.Data.URL
}

func postForm(urlStr string, referer string, params map[string]string) (string, error) {
	data := url.Values{}
	for k, v := range params {
		data.Set(k, v)
	}

	req, err := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", strconv.Itoa(len(data.Encode())))
	req.Header.Set("Referer", "https://www.fangpi.net/music/67344")
	req.Header.Set("Referer", referer)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/151.0.0.0 Safari/537.36")

	client := &http.Client{
		Timeout: 45 * time.Second,
	}

	resp, err := client.Do(req)
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

func ConvertToJSON(input string) (string, error) {
	if !strings.HasPrefix(input, "{") {
		input = "{" + input + "}"
	}

	out := input

	for i := 0; i < 1; i++ {
		unquoted, err := strconv.Unquote(`"` + out + `"`)
		if err != nil {
			break
		}
		out = unquoted
	}

	return out, nil
}

func extractJSON(input string) string {
	// JSON.parse("xxxx")
	start := strings.Index(input, "(")
	end := strings.LastIndex(input, ")")

	if start == -1 || end == -1 || start >= end {
		return input
	}

	content := input[start+1 : end]
	content = strings.Trim(content, `"'`)
	// util.Logi("before parse:%v", content)
	json := CleanJSON(content)
	return json
}

func CleanJSON(input string) string {
	re := regexp.MustCompile(`\\u[0-9a-fA-F]{4}`)
	out := re.ReplaceAllStringFunc(input, func(match string) string {
		code, err := strconv.ParseInt(match[2:], 16, 32)
		if err != nil {
			return match
		}
		return string(rune(code))
	})
	out = strings.ReplaceAll(out, "\\", "")
	return out
}

func (s *SpiderFangpiService) DecodeSongDetail(jsonStr string) (*model.SongDetail, error) {
	if jsonStr == "" {
		return nil, nil
	}

	var raw map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &raw); err != nil {
		return nil, err
	}

	ret := &model.SongDetail{}

	if v, ok := raw["mp3_id"].(float64); ok {
		ret.ID = int64(v)
	}

	if v, ok := raw["play_id"].(string); ok {
		ret.PlayID = v
	}

	if v, ok := raw["mp3_title"].(string); ok {
		ret.Title = decodeUnicode(v)
	}

	if v, ok := raw["mp3_author"].(string); ok {
		ret.Author = decodeUnicode(v)
	}

	if v, ok := raw["mp3_cover"].(string); ok {
		ret.Cover = strings.ReplaceAll(v, "\\", "")
	}

	if v, ok := raw["mp3_duration"].(string); ok {
		ret.DurSeconds = parseTimeToSeconds(v)
	}

	return ret, nil
}

func decodeUnicode(s string) string {
	res, err := strconv.Unquote(`"` + s + `"`)
	if err != nil {
		return s
	}
	return res
}

func parseTimeToSeconds(t string) int {
	parts := strings.Split(t, ":")
	if len(parts) != 2 {
		return 0
	}

	min, _ := strconv.Atoi(parts[0])
	sec, _ := strconv.Atoi(parts[1])

	return min*60 + sec
}

// 从FP网站爬取查询数据
func (s *SpiderFangpiService) QueryFromFangpiWeb(query string) ([]model.Song, error) {
	url := fmt.Sprintf("%s/s/%s", FANGPI_URL, query)
	respHtml, err := BaseSvr.FetchHtml(url)
	if err != nil {
		util.Logi("fetch html form %s error", url)
		return nil, err
	}

	songList, err := s.ParseQueryHtml(respHtml)
	return songList, err
}

func (s *SpiderFangpiService) ParseQueryHtml(respHtml string) ([]model.Song, error) {
	// util.Logi("fetchhtml:%s", respHtml)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(respHtml))
	if err != nil {
		return nil, err
	}

	util.Logi("respHtml parse ...")

	var songList []model.Song
	doc.Find(".card-body .row.no-gutters.py-2d5.border-top.align-items-center").Each(func(i int, r *goquery.Selection) {
		song := s.fillSpiderSong(r)
		if song != nil {
			songList = append(songList, *song)
		}
	})
	util.Logi("respHtml parse end")
	return songList, nil
}

func (s *SpiderFangpiService) fillSpiderSong(row *goquery.Selection) *model.Song {
	if row == nil {
		return nil
	}

	author := strings.TrimSpace(row.Find("small").First().Text())
	music := strings.TrimSpace(row.Find("span").First().Text())
	href, _ := row.Find("a").First().Attr("href")
	util.Logi("read song: author %s music %s href %s", author, music, href)

	if author == "" || music == "" || href == "" {
		return nil
	}

	originId := s.FindOriginIdFromHref(href)
	var musicId string = util.BuildMid(originId, util.MUSIC_SRC_FANGPI)

	util.Logi("mid: %s\t%s - %s \t%s\n", musicId, music, author, href)
	return &model.Song{
		Mid:    musicId,
		Author: author,
		Name:   music,
	}
}

func (s *SpiderFangpiService) FindOriginIdFromHref(href string) string {
	if util.IsEmpty(href) {
		return ""
	}

	index := strings.LastIndex(href, "/")
	if index >= 0 {
		return href[index+1:]
	}
	return ""
}
