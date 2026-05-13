package controller

import (
	"geji/data"
	"geji/service"
	"geji/util"

	"github.com/gin-gonic/gin"
)

func KvCachePut(c *gin.Context) {
	key := c.Query("key")
	value := c.Query("value")
	util.Logi("key : %s , value :%s", key, value)

	user, ok := GetUserClaims(c)
	if !ok {
		util.Logi("claims is not exist")
	}
	util.Logi("uid : %v", user.Uid)

	result := service.KVSvr.Put(key, value)
	if result {
		util.Success(c, "设置成功")
	} else {
		util.Fail(c, data.HTTP_CODE_SERVER_ERR, "设置失败")
	}
}

func KvCacheGet(c *gin.Context) {
	key := c.Query("key")
	util.Logi("key : %s", key)

	value := service.KVSvr.Get(key)
	util.Success(c, value)
}

func KvCacheDel(c *gin.Context) {
	key := c.Query("key")
	util.Logi("key : %s", key)

	result := service.KVSvr.Del(key)
	if result {
		util.Success(c, "删除成功")
	} else {
		util.Fail(c, data.HTTP_CODE_SERVER_ERR, "删除失败")
	}
}
