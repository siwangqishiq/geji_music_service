package controller

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"geji/config"
	"geji/util"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		util.Fail(c, http.StatusBadRequest, "未选择上传文件")
		return
	}

	ext := filepath.Ext(file.Filename)
	now := time.Now().UnixMilli()
	rawName := fmt.Sprintf("%d%s", now, file.Filename)
	hash := md5.Sum([]byte(rawName))
	fileName := hex.EncodeToString(hash[:]) + ext

	saveDir := config.MEDIA_PATH
	// 保存路径
	savePath := filepath.Join(saveDir, fileName)
	err = c.SaveUploadedFile(file, savePath)
	if err != nil {
		util.Fail(c, http.StatusBadRequest, "上传失败")
		return
	}

	url := util.GetUrlFromLocalPath(savePath)
	// 返回数据
	util.Success(c, gin.H{
		"url":      url,
		"mime":     util.FindFileMime(file),
		"filesize": file.Size,
	})
}
