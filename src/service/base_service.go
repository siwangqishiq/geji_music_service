package service

import (
	"io"
	"net/http"
)

type BaseService struct {
}

var BaseSvr *BaseService

func init() {
	BaseSvr = &BaseService{}
}

func (s *BaseService) FetchHtml(url string) (string, error) {
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
