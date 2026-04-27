package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 统一响应格式
type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

func SuccessResp(data any) Response {
	return Response{
		Code: 200,
		Msg:  "",
		Data: data,
	}
}

func SuccessRespWithMsg(data any, msg string) Response {
	return Response{
		Code: 200,
		Msg:  msg,
		Data: data,
	}
}

func FailedResp(code int, errMsg string) Response {
	return Response{
		Code: code,
		Msg:  errMsg,
	}
}

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, SuccessResp(data))
}

func SuccessWithMsg(c *gin.Context, data any, msg string) {
	c.JSON(http.StatusOK, SuccessRespWithMsg(data, msg))
}

func Fail(c *gin.Context, errCode int, msg string) {
	c.JSON(http.StatusOK, FailedResp(errCode, msg))
}
