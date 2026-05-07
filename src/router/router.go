package router

import (
	"geji/config"
	"geji/controller"
	"geji/middleware"
	"geji/util"

	"github.com/gin-gonic/gin"
)

// 不需鉴权的普通url
func installCommonRouter(r *gin.Engine) {
	util.Logi("install common rounter ")

	r.GET("/version", func(c *gin.Context) {
		util.Success(c, gin.H{"version": config.VERSION})
	})
}

func installStaticRouter(r *gin.Engine) {
	util.Logi("install static router")
	r.Static("media", config.MEDIA_PATH)
}

// 需要鉴权的url
func installApiRouter(r *gin.Engine) {
	util.Logi("install api router")

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())

	api.GET("/search", controller.Search)      //巡查总数据统计
	api.GET("/detail", controller.MusicDetail) //查询详情
}

func InitRouter() *gin.Engine {
	util.Logi("初始化路由配置")
	r := gin.New()

	// 全局中间件
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())
	r.Use(middleware.Cors())

	// 公共接口
	installCommonRouter(r)
	installStaticRouter(r)

	//
	installApiRouter(r)

	return r
}
