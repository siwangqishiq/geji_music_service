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

	r.GET("/hello", controller.Hello)
}

// 需要鉴权的url
func installApiRouter(r *gin.Engine) {
	util.Logi("install api router")
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())

	{
		// api.GET("/statistic", controller.GetInspectStatistic) //巡查总数据统计
	}
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
	installApiRouter(r)

	return r
}
