package router

import (
	"geji/config"
	"geji/controller"
	"geji/middleware"
	"geji/util"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// 不需鉴权的普通url
func installCommonRouter(r *gin.Engine) {
	util.Logi("install common rounter ")

	r.GET("/version", func(c *gin.Context) {
		util.Success(c, gin.H{"version": config.VERSION})
	})

	//文件上传
	r.POST("/uploadfile", controller.UploadFile)

	//注册接口
	r.POST("/register", controller.AccountCreate)
	//登录接口
	r.POST("/login", controller.AccountLogin)
}

func installStaticRouter(r *gin.Engine) {
	util.Logi("install static router")
	r.Static("media", config.MEDIA_PATH)
	// r.Static("/web", "../web")

	webRoot := "../web"
	r.Use(func(c *gin.Context) {
		path := webRoot + c.Request.URL.Path
		// 文件存在
		if _, err := os.Stat(path); err == nil {
			http.FileServer(http.Dir(webRoot)).ServeHTTP(c.Writer, c.Request)
			c.Abort()
			return
		}
		// Flutter Web SPA
		// c.File(webRoot + "/index.html")
		// c.Abort()
	})
}

// 需要鉴权的url
func installApiRouter(r *gin.Engine) {
	util.Logi("install api router")

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())

	api.GET("/search", controller.Search)      //巡查总数据统计
	api.GET("/detail", controller.MusicDetail) //查询详情

	if config.IS_DEBUG {
		api.GET("kvput", controller.KvCachePut)
		api.GET("kvget", controller.KvCacheGet)
		api.GET("kvdel", controller.KvCacheDel)
	}
}

func InitRouter() *gin.Engine {
	util.Logi("初始化路由配置")
	r := gin.New()
	r.MaxMultipartMemory = config.MAX_UPLOAD_SIZE

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
