package util

import (
	"fmt"
	"strings"
)

const DIV string = "-"

const (
	MUSIC_SRC_FANGPI = "fangpi"
	MUSIC_SRC_MINE   = "mine"
)

func BuildMid(originId string, src string) string {
	originId = strings.TrimSpace(originId)
	return fmt.Sprintf("%s%s%s", src, DIV, originId)
}

func ParseMid(mid string) (pSrc *string, pId *string) {
	divIndex := strings.Index(mid, DIV)
	if divIndex < 0 {
		return nil, nil
	}

	mids := strings.Split(mid, DIV)
	var src string = mids[0]
	var id string = mids[1]

	return &src, &id
}
