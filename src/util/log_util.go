package util

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

var tagMap = make(map[string]int64)

func init() {
	initLogger()
}

func initLogger() {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.999999999"))
	}
	log_path := fmt.Sprintf("../logs/service_%s.log",
		time.Now().Format("2006-01-02_15_04_05"))
	cfg.DisableCaller = true
	cfg.Encoding = "console"
	cfg.OutputPaths = []string{"stdout", log_path}
	l, _ := cfg.Build()
	Logger = l
}

func Logi(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	Logger.Sugar().Info(msg)
}

func Loge(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	Logger.Sugar().Error(msg)
}

func LogLimitI(tag string, a ...any) {
	lstTime, ok := tagMap[tag]
	curTime := GetCurrentTimeMills()
	if ok {
		if curTime-lstTime >= 1000 {
			tagMap[tag] = curTime
			Logger.Sugar().Info(a)
		}
	} else {
		tagMap[tag] = curTime
		Logger.Sugar().Info(a)
	}
}
