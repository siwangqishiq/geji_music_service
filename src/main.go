package main

import (
	"fmt"
	"geji/config"
	"geji/router"
	"geji/util"
)

func main() {
	util.Logi("歌姬音乐服务端启动...")
	r := router.InitRouter()

	util.Logi("启动服务 %v", config.VERSION)
	r.Run(fmt.Sprintf(":%d", config.Port))
}
