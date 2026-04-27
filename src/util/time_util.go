package util

import (
	"time"
)

func GetCurrentTimeMills() int64 {
	curTime := time.Now()
	return curTime.UnixMilli()
}

func IsAM(timeStr string) (bool, error) {
	ret, err := IsPM(timeStr)
	return !ret, err
}

// true = 下午，false = 上午
func IsPM(timeStr string) (bool, error) {
	layout := "2006-01-02 15:04:05"

	t, err := time.Parse(layout, timeStr)
	if err != nil {
		return false, err
	}
	return t.Hour() >= 12, nil
}
