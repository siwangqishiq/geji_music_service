package controller

import (
	"geji/model"
	"geji/service"
	"geji/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 注册
func AccountCreate(c *gin.Context) {
	var req model.CreateAccountReq
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Fail(c, http.StatusBadRequest, "参数错误: 账户 密码 为必填项")
		return
	}

	service.AccSvr.AccountCreate(c, &req)
}

// 登录
func AccountLogin(c *gin.Context) {
	var req model.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Fail(c, http.StatusBadRequest, "参数错误: 账户 密码 为必填项")
		return
	}

	service.AccSvr.Login(c, &req)
}
