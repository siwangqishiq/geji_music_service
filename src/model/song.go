package model

type Song struct {
	Mid           string `json:"mid"`
	Author        string `json:"author,omitempty"`
	Name          string `json:"name"`
	Href          string `json:"href,omitempty"`
	Cover         string `json:"cover,omitempty"`
	Lyc           string `json:"lyc,omitempty"`
	MusicUrl      string `json:"musicUrl,omitempty"`
	BackupUrl     string `json:"backupUrl,omitempty"`
	DurationMills int32  `json:"durationMills,omitempty"`
	Desc          string `json:"desc,omitempty"`
}

type SongDetail struct {
	ID         int64  `json:"mp3_id"`
	PlayID     string `json:"play_id"`
	Title      string `json:"mp3_title"`
	Author     string `json:"mp3_author"`
	Cover      string `json:"mp3_cover"`
	DurSeconds int
	PlayURL    string
}

type PlayUrlData struct {
	IsWhileURL bool   `json:"is_while_url"`
	URL        string `json:"url"`
	UT         bool   `json:"ut"`
}
