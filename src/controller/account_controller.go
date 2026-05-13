package controller

import (
	"geji/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateAccountReq struct {
	Account  string `json:"account" binding:"required"`
	Nickname string `json:"nickname"`
	Password string `json:"password" binding:"required"`
}

// 注册
func AccountCreate(c *gin.Context) {
	var req CreateAccountReq
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Fail(c, http.StatusBadRequest, "参数错误: account nickname password 为必填项")
		return
	}

}

type LoginReq struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 登录
func AccountLogin(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Fail(c, http.StatusBadRequest, "参数错误: account password 为必填项")
		return
	}

}
