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
